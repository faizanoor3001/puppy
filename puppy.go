package puppy

import (
	"github.com/faizanoor3001/dog"
)

// the funcs are exported since they start with capital letter
func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func WhenGrownUp() string {
	return dog.WhenGrownUp(Barks())
}

//	func Fromv1() string {
//		return "From version 1"
//	}
//
//	func Fromv2() string {
//		return "From version 2"
//	}
func Fromv2Major() string {
	return "From major version 2"
}
