// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Tue Jan 10 17:19:41 2017

package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "os"
	"os/exec"
	"regexp"
)

// const (
// RegexStr                 = "commit ([a-z0-9]{40})\n"
// ReplaceTrimToken         = "\nTHIS_IS_A_TRIM_TOKEN_FOR_SPLIT\nMd5: ${1}\n"
// RegexStrAfterReplacement = "\nTHIS_IS_A_TRIM_TOKEN_FOR_SPLIT\n"
// )

func main() {
	/*
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalf("Open file %s failed!\n", os.Args[1])
		}
		content := string(bytes)
		new_content := regexp.MustCompile(RegexStr).ReplaceAllString(content, ReplaceTrimToken)
		ss := regexp.MustCompile(RegexStrAfterReplacement).Split(new_content, -1)
		for _, s := range ss {
			fmt.Println(s)
			fmt.Println("===========================")
		}
	*/

	// str := "    Merge branch 'V_6_13' into 'V_6_13'\n\n     INK-000 move ut into docker\n\n     See merge request !4090"
	str := " INK-0000 move ut into docker"
	// reg := regexp.MustCompile(" ink-([0-9]{3}) ")
	// reg := regexp.MustCompile("[ |\t][INK|ink][ |-]([0-9]{3})[ |\t]")
	reg := regexp.MustCompile(" (?:INK|ink)-([0-9]{3,5}) ")
	fmt.Println(reg.FindStringSubmatch(str))
	// fmt.Println(reg.FindStringSubmatch(str)[1])

	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
