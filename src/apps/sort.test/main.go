// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Wed Jan 11 15:07:05 2017

package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	id    string
	value int
}

type MyStructsString []*MyStruct

func (ss MyStructsString) Len() int {
	return len(ss)
}

func (ss MyStructsString) Less(i, j int) bool {
	return ss[i].id <= ss[j].id
}

func (ss MyStructsString) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

type MyStructsInt []*MyStruct

func (ss MyStructsInt) Len() int {
	return len(ss)
}

func (ss MyStructsInt) Less(i, j int) bool {
	return ss[i].value <= ss[j].value
}

func (ss MyStructsInt) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	ss := MyStructsString{
		{id: "abc", value: 3},
		{id: "12a", value: 5},
		{id: "fsf", value: 2},
	}
	sort.Sort(ss)
	for _, s := range ss {
		fmt.Println(*s)
	}

	fmt.Println("=======")

	sss := MyStructsInt(ss)
	// sss := MyStructsInt{
	// 	{id: "abc", value: 3},
	// 	{id: "12a", value: 5},
	// 	{id: "fsf", value: 2},
	// }
	sort.Sort(sss)
	for _, t := range sss {
		fmt.Println(*t)
	}
}
