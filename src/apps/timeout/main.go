// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Thu Feb 23 14:25:18 2017

package main

import "fmt"
import "time"
import "context"

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			for {
				fmt.Println("Processing...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 3)
}
