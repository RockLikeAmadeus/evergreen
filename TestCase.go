package main

import (
	"reflect"
)

type TestCase struct {
	name   string
	wasRun bool
}

func NewTestCase(name string) *TestCase {
	test := TestCase{name: name}
	test.wasRun = false
	return &test
}

func (t *TestCase) TestMethod() {
	t.wasRun = true
}

func (t *TestCase) Run() {
	method := reflect.ValueOf(t).MethodByName(t.name)
	method.Call(nil)
}