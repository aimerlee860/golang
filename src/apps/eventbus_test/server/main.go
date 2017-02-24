// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Fri Feb 24 14:51:04 2017

package main

import "fmt"
import "github.com/asaskevich/EventBus"

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	/*
		bus := EventBus.New()
		bus.Subscribe("main:calculator", calculator)
		bus.Publish("main:calculator", 20, 40)
		bus.Unsubscribe("main:calculator", calculator)
	*/
	server := EventBus.NewServer(":2010", "/_server_bus_", EventBus.New())
	server.Start()
	server.EventBus().Publish("main:calculator", 4, 6)

	select {}
	server.Stop()
}
