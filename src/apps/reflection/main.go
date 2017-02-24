// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Mon Feb 20 17:58:53 2017

package main

import (
	"fmt"
	// "reflect"

	"apps/reflection/checker"
)

/*
type Blog struct {
	Name string
}

func (this Blog) Test() string {
	fmt.Println("this is Test method")
	return this.Name
}

func main() {
	var o interface{} = &Blog{"xiaorui.cc"}
	v := reflect.ValueOf(o)
	fmt.Println(v)
	m := v.MethodByName("Test")
	rets := m.Call([]reflect.Value{})
	fmt.Println(rets)
	fmt.Println(rets[0])
}
*/

/*
//定义注册结构map
type registerStructMaps map[string]reflect.Type

//根据name初始化结构
//在这里根据结构的成员注解进行DI注入，这里没有实现，只是简单都初始化
func (rsm registerStructMaps) New(name string) (c interface{}, err error) {
	if v, ok := rsm[name]; ok {
		c = reflect.New(v).Interface()
	} else {
		err = fmt.Errorf("not found %s struct", name)
	}
	return
}

//根据名字注册实例
func (rsm registerStructMaps) Register(name string, c interface{}) {
	rsm[name] = reflect.TypeOf(c).Elem()
}

type Test struct {
	value string
}

func (test *Test) SetValue(value string) {
	test.value = value
}
func (test *Test) Print() {
	log.Println(test.value)
}
func main() {
	rsm := registerStructMaps{}
	//注册test
	rsm.Register("test", &Test{})
	//获取新的test的interface
	test11, _ := rsm.New("test")
	test22, _ := rsm.New("test")
	//因为 test11 和 test22都是interface{},必须转换为*Test
	test1 := test11.(*Test)
	test2 := test22.(*Test)
	test1.SetValue("aaa")
	test2.SetValue("bbb")
	test1.Print()
	test2.Print()
}
*/

func main() {
	checker_, _ := checker.NewChecker("CheckerOne")
	checker := checker_.(*checker.CheckerOne)
	checker.SetName("name")
	fmt.Println(checker.NameString())
}
