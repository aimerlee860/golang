package visibility

import "fmt"

var P int = 1
var p int = 2

func private_func() {
    fmt.Println("This is a private func!")
}

func PublicFunc() {
    fmt.Println("This is a public func!")
}
