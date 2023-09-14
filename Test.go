package main

type Test struct {
	name   string
	wasRun bool
}

func NewTest(name string) *Test {
	test := Test{name: name}
	test.wasRun = false
	return &test
}

func (t *Test) testMethod() {
	t.wasRun = true
}

func (t *Test) Run() {
	t.testMethod()
}
