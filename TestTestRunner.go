package main

import "fmt"

func main() {
	test := newTest("testMethod")
	fmt.Println(test.wasRun) // Should print "false"
	test.testMethod()
	fmt.Println(test.wasRun) // Should print "true"
}

type test struct {
	name   string
	wasRun bool
}

func newTest(name string) *test {
	test := test{name: name}
	test.wasRun = false
	return &test
}

func (t *test) testMethod() {
	t.wasRun = true
}
