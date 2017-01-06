// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Wed Dec 28 19:31:52 2016

package main

import "time"

type Node struct {
	index   int
	numbers []int
}

func main() {
	exit := make(chan bool)
	go func() {
		for {
			node := new(Node)
			node.index = 1
			time.Sleep(10 * time.Millisecond)
		}
		exit <- true
	}()
	sig, ok := <-exit
	if sig && ok {

	}
}
