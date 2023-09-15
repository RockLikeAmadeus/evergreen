package main

import "fmt"

func main() {
	test := NewTest("TestMethod")
	fmt.Println(test.wasRun) // Should print "false"
	test.Run()
	fmt.Println(test.wasRun) // Should print "true"
}
