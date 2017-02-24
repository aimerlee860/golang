// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Tue Dec 20 15:13:18 2016

package main

import (
	"fmt"
	"time"
	// "log"
	"reflect"
)

const (
	infoLog int32 = iota
)

func init() {
	fmt.Println("Still working for initialization!")
}

type Funcer interface {
	Func(string, int) error
}

type TT struct{}

func (t TT) Func(name string, id int) (err error) { return nil }

// func print_args(args ...interface{}) {
//     for key, val := range args {
//         fmt.Println("key: %d, val: %s", key, val)
//     }
// }

type Action func() error

func Echo() error {
	return nil
}

func FuncTest(a Action) {
	if err := a(); err == nil {
		fmt.Println("Nil")
	}
}

type XXX struct {
	YYY struct {
		a string
		b int
	}
}

func (x *XXX) GetA() string {
	return x.YYY.a
}

type MyType struct {
	Name string
}

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

	FuncTest(Echo)

	var x XXX
	// fmt.Println(x.YYY.a)
	fmt.Println(x.GetA())

	strs := []string{"1", "2"}
	strs2 := []string{"3", "4"}
	fmt.Println(append(strs2, strs...))
	if _, err := time.Parse("2016-12-22 02:51:37", "2016-02-22 02:51:37"); err != nil {
		fmt.Println("error ", err)
	}

	fmt.Println(reflect.TypeOf(MyType{}).NumField())

	main := TT{}
	main.Func("x", 1)
}
