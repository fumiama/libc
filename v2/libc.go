// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package libc provides run time support for programs generated by the
// [ccgo] C to Go transpiler, version 4 or later.
//
// # Bugs
//
// A subset of musl v0.7.0, alpha version from 2011, is exported.
//
// Only a single instance of TLS exists. All calls to NewTLS return the same
// instance. As a consequnce, there's no concurrency support yet.  All client
// calls into libc must return before issuing any other call into libc.
//
// # libc API documentation copyright
//
// From [Linux man-pages Copyleft]
//
//	Permission is granted to make and distribute verbatim copies of this
//	manual provided the copyright notice and this permission notice are
//	preserved on all copies.
//
//	Permission is granted to copy and distribute modified versions of this
//	manual under the conditions for verbatim copying, provided that the
//	entire resulting derived work is distributed under the terms of a
//	permission notice identical to this one.
//
//	Since the Linux kernel and libraries are constantly changing, this
//	manual page may be incorrect or out-of-date. The author(s) assume no
//	responsibility for errors or omissions, or for damages resulting from
//	the use of the information contained herein. The author(s) may not have
//	taken the same level of care in the production of this manual, which is
//	licensed free of charge, as they might when working professionally.
//
//	Formatted or processed versions of this manual, if unaccompanied by the
//	source, must acknowledge the copyright and authors of this work.
//
// [Linux man-pages Copyleft]: https://spdx.org/licenses/Linux-man-pages-copyleft.html
// [ccgo]: http://modernc.org/ccgo/v4
package libc // import "modernc.org/libc/v2"
