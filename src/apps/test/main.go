// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Tue Dec 20 15:13:18 2016

package main

import (
	"fmt"
	"time"
	// "log"
)

const (
	infoLog int32 = iota
)

func init() {
	fmt.Println("Still working for initialization!")
}

// func print_args(args ...interface{}) {
//     for key, val := range args {
//         fmt.Println("key: %d, val: %s", key, val)
//     }
// }

func main() {
	done := make(chan bool)
	fmt.Println("This is a test!")
	fmt.Println(infoLog)
	// print_args("args", 1, true)

	fmt.Println("==========================================\n")

	fmt.Printf("number :%d\n", 1)

	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")

	go func(str string) {
		defer func() { done <- true }()
		fmt.Println(str)
	}("This is a test for goroutine!")

	sig, ok := <-done
	if sig && ok {
		fmt.Println("Everything is done, exit...")
	}
	// time.Sleep(time.Second)
}
