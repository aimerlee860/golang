// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Tue Feb 21 18:33:51 2017

package main

import "reflect"
import "fmt"
import "time"

var registerMaps map[string]TypeInterface

type TypeA struct {
}

type TypeB struct {
}

type TypeInterface interface {
	Print(message string)
}

func init() {
	if registerMaps == nil {
		registerMaps = make(map[string]TypeInterface)
	}
	registerMaps["TypeA"] = &TypeA{}
}

func init() {
	if registerMaps == nil {
		registerMaps = make(map[string]TypeInterface)
	}
	registerMaps["TypeB"] = &TypeA{}
}

func (TypeA) Print(message string) {
	fmt.Println("TypeA: " + message)
}

func (TypeB) Print(message string) {
	fmt.Println("TypeB: " + message)
}

func GetType(tpe string) TypeInterface {
	switch tpe {
	case "A":
		return &TypeA{}
	case "B":
		return &TypeB{}
	default:
		panic("Unknown")
	}
}

type NestedStruct struct {
	BaseStruct struct {
		Name string
		ID   string
	}
}

type NoInterface struct {
}

func (n NoInterface) Test() {
	fmt.Println("NoInterface")
}

func main() {
	tpe := GetType("B")
	tpe.Print("This is a test!")
	fmt.Println(reflect.TypeOf(&TypeA{}))
	registerMaps["TypeA"].Print("This is a test!")

	var ns NestedStruct
	ns.BaseStruct = struct{ Name, ID string }{"", ""}

	var maps map[string]int
	if maps == nil {
		fmt.Println("really nil")
	}

	n := NoInterface{}
	n.Test()

	exit := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		close(exit)
	}()
	<-exit
}
