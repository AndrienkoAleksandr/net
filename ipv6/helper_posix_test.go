// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows zos

package ipv6_test

import (
	"errors"
	"os"
	"syscall"

	"github.com/AndrienkoAleksandr/net/ipv6"
)

func protocolNotSupported(err error) bool {
	switch err := err.(type) {
	case syscall.Errno:
		switch err {
		case syscall.EPROTONOSUPPORT, syscall.ENOPROTOOPT:
			return true
		}
	case *os.SyscallError:
		switch err := err.Err.(type) {
		case syscall.Errno:
			switch err {
			case syscall.EPROTONOSUPPORT, syscall.ENOPROTOOPT:
				return true
			}
		}
	}
	return errors.Is(err, ipv6.ErrNotImplemented)
}
