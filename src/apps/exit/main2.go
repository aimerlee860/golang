// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Fri Feb 24 11:36:34 2017

package main

import "fmt"

var exit chan struct{}

func Test() {
	var name string
	fmt.Scanln(&name)
	close(exit)
}

func main() {
	exit = make(chan struct{})
	go Test()
	<-exit

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}
