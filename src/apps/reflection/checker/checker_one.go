// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Mon Feb 20 18:30:51 2017

package checker

import "reflect"

type CheckerOne struct {
	Name string
}

func init() {
	CheckerMaps["CheckerOne"] = reflect.TypeOf(&CheckerOne{}).Elem()
}

func (c *CheckerOne) SetName(name string) {
	c.Name = name
}

func (c *CheckerOne) NameString() string {
	return c.Name
}
