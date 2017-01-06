// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Tue Dec 20 17:07:32 2016

package main

import (
    // "os"
    "fmt"
    // "time"
)

/*
func timeoutFlush(timeout time.Duration) {
    done := make(chan bool, 1)
    go func() {
        // Flush() // calls logging.lockAndFlushAll()
        time.Sleep(timeout / 2)
        done <- true
    }()
    select {
    case <-done:
        fmt.Fprintln(os.Stderr, "Flush took", timeout)
    case <-time.After(timeout):
        fmt.Fprintln(os.Stderr, "Flush took longer than", timeout)
    }
}
*/

func main() {
    fmt.Println("This is a test for goroutine!")
    // timeoutFlush(1000)
    c := make(chan int, 1)
    select {
    case c <- 10:
        fmt.Println("write 10 into c")
    default:
        fmt.Println("Nothing to write")
    }

    select {
    case c <- 11:
        fmt.Println("write 11 into c")
    default: // c中只有10，没有11
        fmt.Println("Nothing to write")
    }

    select {
    case v, ok := <-c:
        fmt.Println("frist read", v, ok)
        // 读出来一个，v=10, ok=true
    default:
        fmt.Println("Nothing to read")
    }

    select {
    case v, ok := <-c:
        fmt.Println("second read", v, ok)
    default: // 没有可读的，走这个分支
        fmt.Println("Nothing to read")
    }

    // v, ok := <-c // 从chan中读取
    // close(c)
    // v, ok = <-c // 从chan中读取
    // fmt.Println(v, ok)
}
