// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Wed Jan 11 14:04:32 2017

/*
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := "ls"
	args := []string{"-l", "-s", "-h"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully do ls -l -s -h")
}
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "ls"
	cmdArgs := []string{"-l", "-s", "-h"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running ls -l -s -h command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	fmt.Println(sha)
	// firstSix := sha[:6]
	// fmt.Println("The first six chars of the SHA at HEAD in this repo are", firstSix)
	arr := []int{1, 2, 3}
	fmt.Println(arr)
}
