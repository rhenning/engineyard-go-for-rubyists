package main

import (
	"fmt"
)

func executeMyFunction(f func()) {
	f()
}

func main() {
	f := func() {
		fmt.Println("I am a function")
	}

	executeMyFunction(f)
}
