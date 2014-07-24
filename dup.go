package dup

import "syscall"
import "os"

// Dup the given `fd` and return a RW pipe.
func Dup(fd int, name string) (r *os.File, w *os.File, err error) {
	r, write, err := os.Pipe()
	if err != nil {
		return
	}

	dup, err := syscall.Dup(fd)
	if err != nil {
		return
	}

	w = os.NewFile(uintptr(dup), name)
	err = syscall.Dup2(int(write.Fd()), fd)

	return
}
