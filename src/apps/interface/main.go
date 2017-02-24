// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Tue Dec 27 15:26:30 2016

package main

import "fmt"

type TestInterface interface {
	Print(message string)
}

type TestInterfaceImpl struct{}

func (i TestInterfaceImpl) Print(message string) {
	fmt.Println(message)
}

func DoPrint(message string, i TestInterfaceImpl) {
	i.Print(message)
}

type TestType struct {
	String string
}

func FuncWithInterfaceArg(arg interface{}) string {
	return arg.(*TestType).String
}

func main() {
	var i TestInterfaceImpl
	DoPrint("This is a test!", i)

	fmt.Println(FuncWithInterfaceArg(&TestType{String: "hell"}))
}
