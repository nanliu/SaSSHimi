// +build !windows,!linux

package server

import (
	"golang.org/x/sys/unix"
)

const (
	unix_GETATTR = unix.TIOCGETA
)
