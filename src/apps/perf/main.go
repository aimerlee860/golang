// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Mon Jan  9 16:19:13 2017

package main

import (
	"fmt"
	"runtime"
	// "sync"
	"math/rand"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	/*
		var wg sync.WaitGroup
		tStart := time.Now()
		for i := 0; i < 1000000; i++ {
			wg.Add(1)
			go func() {
				fmt.Println("Start a goroutine!")
				wg.Done()
			}()
		}
		tEnd := time.Now()
		fmt.Println("Time use ", tEnd.Sub(tStart))
		wg.Wait()
		tFinal := time.Now()
		fmt.Println("Final time use ", tFinal.Sub(tStart))
	*/
	x := int(rand.Int31())
	tStart := time.Now()
	for i := 0; i < 1000000000; i++ {
		if (x+i)%2 == 0 {
			x += i
		} else {
			x -= i
		}
	}
	tEnd := time.Now()
	fmt.Println("Time use ", tEnd.Sub(tStart))
}
