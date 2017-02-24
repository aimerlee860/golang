// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Thu Feb 23 18:35:33 2017

package main

import "fmt"
import "time"
import "context"
import "sync"

func Test() {
	exit := make(chan struct{})
	for {
		select {
		case <-exit:
			return
		default:
			go func() {
				for {
					fmt.Println("hello")
					time.Sleep(time.Second)
				}
			}()
			var tm int = 5
			for tm > 0 {
				time.Sleep(time.Second)
				tm--
			}
			close(exit)
		}
	}
}

func Test2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(context.Context) {
		go func() {
			for {
				fmt.Println("hello1")
				time.Sleep(time.Second)
			}
		}()
		select {
		case <-ctx.Done():
			return
		}
	}(ctx)
	go func(context.Context) {
		for {
			if ctx.Err() != nil {
				break
			}
			fmt.Println("hello2")
			time.Sleep(time.Second)
		}
	}(ctx)
	select {
	case <-time.After(time.Second * 5):
		cancel()
		return
	}
}

func Test3() {
	go func() {
		var name string
		fmt.Println("shuru")
		fmt.Scanln(&name)
		fmt.Println("shuru")
	}()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		Test2()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Test3()
		wg.Done()
	}()

	wg.Wait()
	// time.Sleep(time.Second * 10)
	select {}
}
