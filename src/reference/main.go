// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Mon Dec 26 15:58:13 2016

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func f_1(a int) {
	a = 2
}

func f_1_1(a *int) {
	*a = 2
}

func f_2(s string) {
	s = "cba"
}

func f_2_1(s *string) {
	*s = "cba"
}

func f_3(v []string) {
	v[0] = "haha"
}

func f_3_1(v []string) {
	v = nil
}

func f_3_2(v *[]string) {
	*v = nil
}

func f_4(m map[int]int) {
	m[1] = 3
	m[3] = 1
}

func f_4_1(m map[int]int) {
	m = nil
}

func f_4_2(m *map[int]int) {
	*m = nil
}

func f_5(b []byte) {
	b[0] = 0x40
}

func f_5_1(b []byte) {
	b = bytes.Replace(b, []byte("1"), []byte("a"), -1)
}

func f_5_2(b *[]byte) {
	*b = bytes.Replace(*b, []byte("1"), []byte("a"), -1)
}

type why struct {
	s []string
}

func (ss why) SetV(s []string) {
	ss.s = s
}

func (ss *why) SetP(s []string) {
	ss.s = s
}

func (ss why) String() string {
	return strings.Join(ss.s, ",")
}

func main() {
	a := 1
	s := "abc"
	v := []string{"sd", "aa"}
	m := map[int]int{1: 1, 2: 2, 3: 3}
	f_1(a)
	f_2(s)
	f_3(v)
	f_4(m)
	fmt.Printf("%d,%s,%v,%v\n", a, s, v, m)
	f_3_1(v)
	f_4_1(m)
	fmt.Printf("%d,%s,%v,%v\n", a, s, v, m)
	f_1_1(&a)
	f_2_1(&s)
	f_3_2(&v)
	f_4_2(&m)
	fmt.Printf("%d,%s,%v,%v\n", a, s, v, m)
	b := []byte("12145178")
	f_5(b)
	fmt.Printf("%s\n", b)
	f_5_1(b)
	fmt.Printf("%s\n", b)
	f_5_2(&b)
	fmt.Printf("%s\n", b)
	ss := &why{}
	ss.SetV([]string{"abc", "efg"})
	fmt.Println(ss)
	ss.SetP([]string{"abc", "efg"})
	fmt.Println(ss)
}
