// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Tue Dec 27 19:54:29 2016

package goroutine

import (
	"fmt"
)

func StartGoRoutine(index int) {
	go func(index int) {
		fmt.Printf("This is %d goroutine!\n", index)
	}(index)
}
