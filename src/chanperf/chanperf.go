// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Mon Dec 26 11:13:04 2016

package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	id   int64
	name string
}

const (
	loopTimes int = 1
)

func test() {
	t0 := time.Now()
	var out MyStruct
	test := make(chan MyStruct, loopTimes)
	for i := 0; i < 100000; i++ {
		for j := 0; j < loopTimes; j++ {
			test <- MyStruct{int64(j), "hello"}
		}
		for j := 0; j < loopTimes; j++ {
			out = <-test
		}
	}
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Println(out.id)
}

func main() {
	test()
}
