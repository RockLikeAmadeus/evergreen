package main

import (
	"reflect"
)

type TestCase struct {
	name     string
	wasRun   bool
	wasSetUp bool
}

func NewTestCase(name string) TestCase {
	test := TestCase{name: name}
	return test
}

func (t *TestCase) TestMethod() {
	t.wasRun = true
}

func (t *TestCase) SetUp() {
	t.wasRun = false
	t.wasSetUp = true
}

func (t *TestCase) Run() {
	t.SetUp()
	method := reflect.ValueOf(t).MethodByName(t.name)
	method.Call(nil)
}
