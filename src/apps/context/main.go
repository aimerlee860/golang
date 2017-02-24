// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Wed Feb 22 16:34:55 2017

package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe(":8989", nil)
	ctx, cancel := context.WithCancel(context.Background())
	// startTime := time.Now()
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go func() {
		time.Sleep(3 * time.Second)
		defer cancel()
	}()
	log.Println(A(ctx))
	// defer cancel()
	// endTime := time.Now()
	// log.Println(endTime.Sub(startTime).String())
	select {}
}

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "C Done"
	}
	return ""
}

func B(ctx context.Context) string {
	ctx, _ = context.WithCancel(ctx)
	go log.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B Done"
	}
	return ""
}

func A(ctx context.Context) string {
	go log.Println(B(ctx))
	select {
	case <-ctx.Done():
		return "A Done"
	}
	return ""
}
