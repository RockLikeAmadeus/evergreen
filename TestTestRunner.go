package main

import "fmt"

func main() {
	test := NewTest("testMethod")
	fmt.Println(test.wasRun) // Should print "false"
	test.run()
	fmt.Println(test.wasRun) // Should print "true"
}
