// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Thu Jan  5 17:58:16 2017

package main

import "time"

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
	select {}
}
