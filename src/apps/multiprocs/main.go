// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Thu Dec 29 18:40:45 2016

package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"syscall"
)

func main() {
	debug.SetGCPercent(-1)

	var r1, r2 uintptr
	var status syscall.Errno

	fmt.Println("This is parent pid", os.Getpid())
	for i := 0; i < 2; i++ {
		r1, r2, status = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if status != 0 || r1 <= 0 || r2 < 0 {
			log.Fatal("Fork child failed!")
		} else if r2 == 1 { // Child process
			fmt.Println("This is child pid ", os.Getpid())
			break
		}
	}
	go func() {
		runtime.GC()
	}()
}
