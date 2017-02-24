// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Thu Feb 23 10:41:30 2017

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		startTime := time.Now()
		if _, err := http.Get("http://localhost:8090/"); err != nil {
			fmt.Println("ERROR")
		}
		fmt.Println(time.Since(startTime))
		time.Sleep(100 * time.Microsecond)
	}
}
