// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Mon Dec 26 16:45:50 2016

package main

import (
	"fmt"
	"time"
)

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	ff := f(0)
	fmt.Println(ff())
	time.Sleep(time.Second * 100)
}
