// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Wed Dec 21 10:24:55 2016

package main

import "fmt"

func main() {
    fmt.Println("Start doing job!")
    // DoJob()
    DeferFunc(1)
    fmt.Println("Finish doing job!")
}

// func DoJob() {
//     defer func() {
//         if e := recover(); e != nil {
//             fmt.Println("Panic is caught:", e)
//         }
//     } ()
//     FuncWithException()
//     fmt.Println("Continue to execute main function!")
// }
// 
// func FuncWithException() {
//     fmt.Println("This is a test for defer, panic and recover!")
//     defer fmt.Println("FuncWithException exit...")
//     panic(12)
//     defer fmt.Println("FuncWithException exit after panic...")
// }

func DeferFunc(types int) {
    defer func() {
        fmt.Println("sfhf")
    } ()
}
