package main

import (
	"reflect"
)

type Test struct {
	name   string
	wasRun bool
}

func NewTest(name string) *Test {
	test := Test{name: name}
	test.wasRun = false
	return &test
}

func (t *Test) TestMethod() {
	t.wasRun = true
}

func (t *Test) Run() {
	method := reflect.ValueOf(t).MethodByName(t.name)
	method.Call(nil)
}
