package main

func main() {
	testWasRun()
	testSetUp()
}

func testWasRun() {
	test := NewTestCase("TestMethod")
	Assert(!test.wasRun)
	test.Run()
	Assert(test.wasRun)
}
func testSetUp() {
	test := NewTestCase("TestMethod")
	test.run()
	Assert(test.wasSetUp)
}
