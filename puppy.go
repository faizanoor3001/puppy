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

func Fromv1() string {
	return "From version 1"
}
