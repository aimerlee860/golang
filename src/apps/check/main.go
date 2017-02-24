// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Tue Feb 21 11:22:42 2017

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	if _, err := net.DialTimeout("tcp", "localhost:8090", 3*time.Second); err != nil {
		fmt.Println("Failed!")
	}
}
