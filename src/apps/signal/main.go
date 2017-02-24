// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Thu Feb 16 18:18:14 2017

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RegistSignalHandler(stopFunction func()) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		fmt.Println("Received sigterm/sigint, stopping")
		stopFunction()
	}()
}

var exit bool = false

func Run() {
	for {
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

func Stop() {
	fmt.Println("Start stopping!")
	exit = true
}

func Main_f() {
	fmt.Println("Start running.")
	defer fmt.Println("stopped.")
	Run()
}

func main() {
	RegistSignalHandler(Stop)
	Main_f()
}
