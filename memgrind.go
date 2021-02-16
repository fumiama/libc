// Copyright 2021 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !libc.membrk,libc.memgrind

// This is a debug-only version of the memory handling functions. When a
// program is built with -tags=libc.memgrind the functions MemAuditStart and
// MemAuditReport can be used to check for memory leaks.

package libc // import "modernc.org/libc"

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"unsafe"

	"modernc.org/libc/errno"
	"modernc.org/libc/sys/types"
	"modernc.org/memory"
)

const (
	memgrind     = true
	memGuardSize = 16
)

func init() {
	if memGuardSize%16 != 0 {
		panic(todo("internal error"))
	}
}

type memGuardT = [memGuardSize]byte

var memGuard memGuardT

type memReportItem struct {
	p, pc uintptr
	s     string
}

func (it *memReportItem) String() string {
	more := it.s
	if more != "" {
		a := strings.Split(more, "\n")
		more = "\n\t\t" + strings.Join(a, "\n\t\t")
	}
	return fmt.Sprintf("\t%s: %#x%s", pc2origin(it.pc), it.p, more)
}

type memReport []memReportItem

func (r memReport) Error() string {
	a := []string{"memory leaks"}
	for _, v := range r {
		a = append(a, v.String())
	}
	return strings.Join(a, "\n")
}

var (
	allocator       memory.Allocator
	allocs          map[uintptr]uintptr // addr: caller
	allocsMore      map[uintptr]string
	frees           map[uintptr]uintptr // addr: caller
	memAudit        memReport
	memAuditEnabled bool
)

func pc2origin(pc uintptr) string {
	f := runtime.FuncForPC(pc)
	var fn, fns string
	var fl int
	if f != nil {
		fn, fl = f.FileLine(pc)
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s", fn, fl, fns)
}

// void *malloc(size_t size);
func Xmalloc(t *TLS, size types.Size_t) uintptr {
	if size == 0 {
		return 0
	}

	size = types.Size_t(roundup(uintptr(size), 16))

	allocMu.Lock()

	defer allocMu.Unlock()

	p0, err := allocator.UintptrMalloc(int(size + 2*memGuardSize))
	if err != nil {
		t.setErrno(errno.ENOMEM)
		return 0
	}

	p := p0 + memGuardSize
	p2 := p + uintptr(size)
	*(*memGuardT)(unsafe.Pointer(p)) = memGuard
	*(*memGuardT)(unsafe.Pointer(p2)) = memGuard

	if memAuditEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			panic("cannot obtain caller's PC")
		}

		delete(frees, p)
		if pc0, ok := allocs[p]; ok {
			panic(fmt.Errorf("%v: malloc returns same address twice, previous call at %v:", pc2origin(pc), pc2origin(pc0)))
		}

		allocs[p] = pc
	}
	return p
}

// void *calloc(size_t nmemb, size_t size);
func Xcalloc(t *TLS, nmemb, size types.Size_t) uintptr {
	rq := int(nmemb * size)
	if rq == 0 {
		return 0
	}

	rq = int(roundup(uintptr(rq), 16))

	allocMu.Lock()

	defer allocMu.Unlock()

	p0, err := allocator.UintptrCalloc(rq + 2*memGuardSize)
	if err != nil {
		t.setErrno(errno.ENOMEM)
		return 0
	}

	p := p0 + memGuardSize
	p2 := p + uintptr(rq)
	*(*memGuardT)(unsafe.Pointer(p)) = memGuard
	*(*memGuardT)(unsafe.Pointer(p2)) = memGuard

	if memAuditEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			panic("cannot obtain caller's PC")
		}

		delete(frees, p)
		if pc0, ok := allocs[p]; ok {
			panic(fmt.Errorf("%v: calloc returns same address twice, previous call at %v:", pc2origin(pc), pc2origin(pc0)))
		}

		allocs[p] = pc
	}
	return p
}

// void *realloc(void *ptr, size_t size);
func Xrealloc(t *TLS, ptr uintptr, size types.Size_t) uintptr {
	switch {
	case ptr == 0 && size == 0:
		return 0
	case ptr == 0 && size != 0:
		return Xmalloc(t, size)
	case size == 0:
		Xfree(t, ptr)
		return 0
	case size != 0:
		size = types.Size_t(roundup(uintptr(size), 16))
		if *(*memGuardT)(unsafe.Pointer(ptr - memGuardSize)) != memGuard {
			panic(fmt.Errorf("corrupted memory guard: %#x", ptr))
		}
	}
	allocMu.Lock()

	defer allocMu.Unlock()

	var pc uintptr
	if memAuditEnabled {
		var ok bool
		if pc, _, _, ok = runtime.Caller(1); !ok {
			panic("cannot obtain caller's PC")
		}

		if ptr != 0 {
			if pc0, ok := frees[ptr]; ok {
				panic(fmt.Errorf("%v: realloc of freed memory, previous call at %v:", pc2origin(pc), pc2origin(pc0)))
			}

			if _, ok := allocs[ptr]; !ok {
				panic(fmt.Errorf("%v: realloc of unallocated memory: %#x", pc2origin(pc), p))
			}

			delete(allocs, ptr)
			delete(allocsMore, ptr)
		}
	}

	p0, err := allocator.UintptrRealloc(ptr, int(size+2*memGuardSize))
	if err != nil {
		t.setErrno(errno.ENOMEM)
		return 0
	}

	p := p0 + memGuardSize
	p2 := p + uintptr(size)
	*(*memGuardT)(unsafe.Pointer(p)) = memGuard
	*(*memGuardT)(unsafe.Pointer(p2)) = memGuard

	if memAuditEnabled {
		delete(frees, p)
		if pc0, ok := allocs[p]; ok {
			panic(fmt.Errorf("%v: realloc returns same address twice, previous call at %v:", pc2origin(pc), pc2origin(pc0)))
		}

		allocs[p] = pc
	}
	return p
}

// void free(void *ptr);
func Xfree(t *TLS, p uintptr) {
	if p == 0 {
		return
	}

	p0 := p - memGuardSize
	if *(*memGuardT)(unsafe.Pointer(p0)) != memGuard {
		panic(fmt.Errorf("corrupted memory guard: %#x", p))
	}

	allocMu.Lock()

	defer allocMu.Unlock()

	if memAuditEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			panic("cannot obtain caller's PC")
		}

		if pc0, ok := frees[p]; ok {
			panic(fmt.Errorf("%v: double free, previous call at %v:", pc2origin(pc), pc2origin(pc0)))
		}

		if _, ok := allocs[p]; !ok {
			panic(fmt.Errorf("%v: free of unallocated memory: %#x", pc2origin(pc), p))
		}

		delete(allocs, p)
		delete(allocsMore, p)
		frees[p] = pc
	}

	allocator.UintptrFree(p0)
}

func UsableSize(p uintptr) types.Size_t {
	p0 := p - memGuardSize
	if *(*memGuardT)(unsafe.Pointer(p0)) != memGuard {
		panic(fmt.Errorf("corrupted memory guard: %#x", p))
	}

	if memAuditEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			panic("cannot obtain caller's PC")
		}

		if _, ok := allocs[p]; !ok {
			panic(fmt.Errorf("%v: usable size of unallocated memory: %#x", pc2origin(pc), p))
		}
	}

	return types.Size_t(memory.UintptrUsableSize(p0)) - 2*memGuardSize
}

// MemAuditStart locks the memory allocator, initializes and enables memory
// auditing. Finaly it unlocks the memory allocator.
//
// Some memory handling errors, like double free or freeing of unallocated
// memory, will panic when memory auditing is enabled.
//
// This memory auditing functionality has to be enabled using the libc.memgrind
// build tag.
//
// It is intended only for debug/test builds. It slows down memory allocation
// routines and it has additional memory costs.
func MemAuditStart() {
	allocMu.Lock()

	defer allocMu.Unlock()

	allocs = map[uintptr]uintptr{} // addr: caller
	allocsMore = map[uintptr]string{}
	frees = map[uintptr]uintptr{} // addr: caller
	memAuditEnabled = true
}

// MemAuditReport locks the memory allocator, reports memory leaks, if any.
// Finally it disables memory auditing and unlocks the memory allocator.
//
// This memory auditing functionality has to be enabled using the libc.memgrind
// build tag.
//
// It is intended only for debug/test builds. It slows down memory allocation
// routines and it has additional memory costs.
func MemAuditReport() (r error) {
	allocMu.Lock()

	defer func() {
		allocs = nil
		allocsMore = nil
		frees = nil
		memAuditEnabled = false
		memAudit = nil
		allocMu.Unlock()
	}()

	if len(allocs) != 0 {
		for p, pc := range allocs {
			memAudit = append(memAudit, memReportItem{p, pc, allocsMore[p]})
		}
		sort.Slice(memAudit, func(i, j int) bool {
			return memAudit[i].String() < memAudit[j].String()
		})
		return memAudit
	}

	return nil
}

func MemAuditAnnotate(pc uintptr, s string) {
	allocMu.Lock()
	allocsMore[pc] = s
	allocMu.Unlock()
}
