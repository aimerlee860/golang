package main

import (
	"fmt"
	"visibility"
)

func main() {
	// visibility.private_func()
	visibility.PublicFunc()
	fmt.Println(visibility.P)
	// fmt.Println(visibility.p)
}
