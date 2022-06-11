package main

import (
	"fmt"
	"try/try"
)

func main() {
	tryTo := try.Try{
		Tries:           3,
		Strategy:        try.ConstantWait{Interval: 5},
		Errors:          []error{},
		FailureCallback: func() { fmt.Println("recovered\n") },
	}

	tryTo.Execute(func() {
		fmt.Println("executing...")
		panic("SomeError")
	})
}
