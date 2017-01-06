// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Wed Dec 28 18:36:56 2016

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	// "time"
)

func fab(num int) int {
	if num <= 0 || num == 1 {
		return 1
	} else {
		return fab(num-1) + fab(num-2)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println(os.Getpid())
	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		fmt.Println(os.Getpid())
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
			fab(int(char) % 20)
			// time.Sleep(time.Nanosecond * 300)
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()

		fmt.Println(os.Getpid())
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
			fab(int(number) % 20)
			// time.Sleep(time.Nanosecond * 200)
		}
		fmt.Println()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
