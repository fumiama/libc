// Code generated by 'ccgo -export-externs X -hide __syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6 -nostdinc -nostdlib -o ../musl_windows_386.go -pkgname libc -Iarch/i386 -Iarch/generic -Iobj/src/internal -Isrc/include -Isrc/internal -Iobj/include -Iinclude copyright.c src/ctype/isalnum.c src/ctype/isalpha.c src/ctype/isdigit.c src/ctype/islower.c src/ctype/isprint.c src/ctype/isspace.c src/ctype/isxdigit.c', DO NOT EDIT.

package libc

import (
	"math"
	"reflect"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ unsafe.Pointer

// musl as a whole is licensed under the following standard MIT license:
//
// ----------------------------------------------------------------------
// Copyright © 2005-2020 Rich Felker, et al.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// ----------------------------------------------------------------------
//
// Authors/contributors include:
//
// A. Wilcox
// Ada Worcester
// Alex Dowad
// Alex Suykov
// Alexander Monakov
// Andre McCurdy
// Andrew Kelley
// Anthony G. Basile
// Aric Belsito
// Arvid Picciani
// Bartosz Brachaczek
// Benjamin Peterson
// Bobby Bingham
// Boris Brezillon
// Brent Cook
// Chris Spiegel
// Clément Vasseur
// Daniel Micay
// Daniel Sabogal
// Daurnimator
// David Carlier
// David Edelsohn
// Denys Vlasenko
// Dmitry Ivanov
// Dmitry V. Levin
// Drew DeVault
// Emil Renner Berthing
// Fangrui Song
// Felix Fietkau
// Felix Janda
// Gianluca Anzolin
// Hauke Mehrtens
// He X
// Hiltjo Posthuma
// Isaac Dunham
// Jaydeep Patil
// Jens Gustedt
// Jeremy Huntwork
// Jo-Philipp Wich
// Joakim Sindholt
// John Spencer
// Julien Ramseier
// Justin Cormack
// Kaarle Ritvanen
// Khem Raj
// Kylie McClain
// Leah Neukirchen
// Luca Barbato
// Luka Perkov
// M Farkas-Dyck (Strake)
// Mahesh Bodapati
// Markus Wichmann
// Masanori Ogino
// Michael Clark
// Michael Forney
// Mikhail Kremnyov
// Natanael Copa
// Nicholas J. Kain
// orc
// Pascal Cuoq
// Patrick Oppenlander
// Petr Hosek
// Petr Skocik
// Pierre Carrier
// Reini Urban
// Rich Felker
// Richard Pennington
// Ryan Fairfax
// Samuel Holland
// Segev Finer
// Shiz
// sin
// Solar Designer
// Stefan Kristiansson
// Stefan O'Rear
// Szabolcs Nagy
// Timo Teräs
// Trutz Behn
// Valentin Ochs
// Will Dietz
// William Haddon
// William Pitcock
//
// Portions of this software are derived from third-party works licensed
// under terms compatible with the above MIT license:
//
// The TRE regular expression implementation (src/regex/reg* and
// src/regex/tre*) is Copyright © 2001-2008 Ville Laurikari and licensed
// under a 2-clause BSD license (license text in the source files). The
// included version has been heavily modified by Rich Felker in 2012, in
// the interests of size, simplicity, and namespace cleanliness.
//
// Much of the math library code (src/math/* and src/complex/*) is
// Copyright © 1993,2004 Sun Microsystems or
// Copyright © 2003-2011 David Schultz or
// Copyright © 2003-2009 Steven G. Kargl or
// Copyright © 2003-2009 Bruce D. Evans or
// Copyright © 2008 Stephen L. Moshier or
// Copyright © 2017-2018 Arm Limited
// and labelled as such in comments in the individual source files. All
// have been licensed under extremely permissive terms.
//
// The ARM memcpy code (src/string/arm/memcpy.S) is Copyright © 2008
// The Android Open Source Project and is licensed under a two-clause BSD
// license. It was taken from Bionic libc, used on Android.
//
// The AArch64 memcpy and memset code (src/string/aarch64/*) are
// Copyright © 1999-2019, Arm Limited.
//
// The implementation of DES for crypt (src/crypt/crypt_des.c) is
// Copyright © 1994 David Burren. It is licensed under a BSD license.
//
// The implementation of blowfish crypt (src/crypt/crypt_blowfish.c) was
// originally written by Solar Designer and placed into the public
// domain. The code also comes with a fallback permissive license for use
// in jurisdictions that may not recognize the public domain.
//
// The smoothsort implementation (src/stdlib/qsort.c) is Copyright © 2011
// Valentin Ochs and is licensed under an MIT-style license.
//
// The x86_64 port was written by Nicholas J. Kain and is licensed under
// the standard MIT terms.
//
// The mips and microblaze ports were originally written by Richard
// Pennington for use in the ellcc project. The original code was adapted
// by Rich Felker for build system and code conventions during upstream
// integration. It is licensed under the standard MIT terms.
//
// The mips64 port was contributed by Imagination Technologies and is
// licensed under the standard MIT terms.
//
// The powerpc port was also originally written by Richard Pennington,
// and later supplemented and integrated by John Spencer. It is licensed
// under the standard MIT terms.
//
// All other files which have no copyright comments are original works
// produced specifically for use as part of this library, written either
// by Rich Felker, the main author of the library, or by one or more
// contibutors listed above. Details on authorship of individual files
// can be found in the git version control history of the project. The
// omission of copyright and license comments in each file is in the
// interest of source tree size.
//
// In addition, permission is hereby granted for all public header files
// (include/* and arch/*/bits/*) and crt files intended to be linked into
// applications (crt/*, ldso/dlstart.c, and arch/*/crt_arch.h) to omit
// the copyright notice and permission notice otherwise required by the
// license, and to use these files without any requirement of
// attribution. These files include substantial contributions from:
//
// Bobby Bingham
// John Spencer
// Nicholas J. Kain
// Rich Felker
// Richard Pennington
// Stefan Kristiansson
// Szabolcs Nagy
//
// all of whom have explicitly granted such permission.
//
// This file previously contained text expressing a belief that most of
// the files covered by the above exception were sufficiently trivial not
// to be subject to copyright, resulting in confusion over whether it
// negated the permissions granted in the license. In the spirit of
// permissive licensing, and of not having licensing issues being an
// obstacle to adoption, that text has been removed.
const ( /* copyright.c:194:1: */
	__musl__copyright__ = 0
)

type ptrdiff_t = int32 /* <builtin>:3:26 */

type size_t = uint32 /* <builtin>:9:23 */

type wchar_t = uint16 /* <builtin>:15:24 */

type va_list = uintptr /* <builtin>:45:27 */

type locale_t = uintptr /* alltypes.h:366:32 */

func Xisalnum(tls *TLS, c int32) int32 { /* isalnum.c:3:5: */
	return (Bool32((func() int32 {
		if 0 != 0 {
			return Xisalpha(tls, c)
		}
		return (Bool32((((uint32(c)) | uint32(32)) - uint32('a')) < uint32(26)))
	}() != 0) || (func() int32 {
		if 0 != 0 {
			return Xisdigit(tls, c)
		}
		return (Bool32(((uint32(c)) - uint32('0')) < uint32(10)))
	}() != 0)))
}

func X__isalnum_l(tls *TLS, c int32, l locale_t) int32 { /* isalnum.c:8:5: */
	return Xisalnum(tls, c)
}

func Xisalpha(tls *TLS, c int32) int32 { /* isalpha.c:4:5: */
	return (Bool32(((uint32(c) | uint32(32)) - uint32('a')) < uint32(26)))
}

func X__isalpha_l(tls *TLS, c int32, l locale_t) int32 { /* isalpha.c:9:5: */
	return Xisalpha(tls, c)
}

func Xisdigit(tls *TLS, c int32) int32 { /* isdigit.c:4:5: */
	return (Bool32((uint32(c) - uint32('0')) < uint32(10)))
}

func X__isdigit_l(tls *TLS, c int32, l locale_t) int32 { /* isdigit.c:9:5: */
	return Xisdigit(tls, c)
}

func Xislower(tls *TLS, c int32) int32 { /* islower.c:4:5: */
	return (Bool32((uint32(c) - uint32('a')) < uint32(26)))
}

func X__islower_l(tls *TLS, c int32, l locale_t) int32 { /* islower.c:9:5: */
	return Xislower(tls, c)
}

func Xisprint(tls *TLS, c int32) int32 { /* isprint.c:4:5: */
	return (Bool32((uint32(c) - uint32(0x20)) < uint32(0x5f)))
}

func X__isprint_l(tls *TLS, c int32, l locale_t) int32 { /* isprint.c:9:5: */
	return Xisprint(tls, c)
}

func Xisspace(tls *TLS, c int32) int32 { /* isspace.c:4:5: */
	return (Bool32((c == ' ') || ((uint32(c) - uint32('\t')) < uint32(5))))
}

func X__isspace_l(tls *TLS, c int32, l locale_t) int32 { /* isspace.c:9:5: */
	return Xisspace(tls, c)
}

func Xisxdigit(tls *TLS, c int32) int32 { /* isxdigit.c:3:5: */
	return (Bool32((func() int32 {
		if 0 != 0 {
			return Xisdigit(tls, c)
		}
		return (Bool32(((uint32(c)) - uint32('0')) < uint32(10)))
	}() != 0) || (((uint32(c) | uint32(32)) - uint32('a')) < uint32(6))))
}

func X__isxdigit_l(tls *TLS, c int32, l locale_t) int32 { /* isxdigit.c:8:5: */
	return Xisxdigit(tls, c)
}
