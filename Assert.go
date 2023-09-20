package main

import "fmt"

func Assert(expression bool) {
	if !expression {
		fmt.Println("Assertion failed.")
	}
}
