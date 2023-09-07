package main

import "syscall"

func init() {
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return
	}
	defer syscall.FreeLibrary(kernel32)

	freeConsole, err := syscall.GetProcAddress(kernel32, "FreeConsole")
	if err != nil {
		return
	}
	syscall.SyscallN(freeConsole)
}
