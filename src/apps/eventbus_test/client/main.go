// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Fri Feb 24 15:01:14 2017

package main

import "fmt"
import "github.com/asaskevich/EventBus"

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	client := EventBus.NewClient(":2015", "/_client_bus_", EventBus.New())
	client.Start()
	client.Subscribe("main:calculator", calculator, ":2010", "/_server_bus_")
	// ...
	client.Stop()
}
