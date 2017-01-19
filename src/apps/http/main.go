// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Mon Dec 26 19:22:33 2016

package main

import (
	"fmt"
	// "goroutine"
	"log"
	"net/http"
	"runtime"
	// "runtime/debug"
)

func IsPrime(i int) bool {
	for j := 2; j < (i/2)+1; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	// debug.SetGCPercent(-1)
	request_num := 0

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		request_num++
		if IsPrime(request_num) {
			fmt.Printf("%d is prime\n", request_num)
		}
		// if request_num%100 == 0 {
		// 	runtime.GC()
		// }
		// goroutine.StartGoRoutine(request_num)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
