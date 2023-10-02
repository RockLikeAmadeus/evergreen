package main

import "fmt"

func main() {
	testWasRunTest := TestCaseTest{TestCase: TestCase{name: "testWasRun"}}
	fmt.Println(testWasRunTest)
	// testWasRunTest.Run()
	// testCaseTest.testSetUp()
	// testCaseTest.testWasRun()
}

type TestCaseTest struct {
	TestCase
	test TestCase
}

func (t *TestCaseTest) setUp() {}

func (t *TestCaseTest) testWasRun() {
	test := TestCase{name: "testMethod"}
	// t.setUp()
	// Assert(!t.testCase.wasRun)
	test.Run()
	Assert(test.wasRun)
}

func (t *TestCaseTest) testSetUp() {
	test := TestCase{name: "testMethod"}
	// t.setUp()
	test.Run()
	Assert(test.wasSetUp)
}
