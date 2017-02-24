// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(JingJie Li)
// Created Time: Mon Feb 20 18:32:43 2017

package checker

import (
	"fmt"
	"reflect"
)

type registerCheckerMaps map[string]reflect.Type

var CheckerMaps registerCheckerMaps

func NewChecker(name string) (instance interface{}, err error) {
	if tpe, ok := CheckerMaps[name]; ok {
		instance = reflect.New(tpe).Interface()
	} else {
		err = fmt.Errorf("not found %s struct", name)
	}
	return
}

type CheckerInterface interface {
	SetName(name string)
	NameString() string
}
