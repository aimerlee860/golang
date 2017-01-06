// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Fri Jan  6 11:19:03 2017

package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { go func() { fmt.Println("Every second!") }() })
	c.Start()
	select {}
}
