package main

func main() {
	test := NewTestCase("TestMethod")
	Assert(!test.wasRun) // Should print "false"
	test.Run()
	Assert(test.wasRun) // Should print "true"
}
