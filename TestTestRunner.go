package main

func main() {
	test := NewTestCase("TestMethod")
	Assert(!test.wasRun)
	test.Run()
	Assert(test.wasRun)
}
