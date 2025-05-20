package puppy

import (
	"github.com/HMSatishVishwakarma/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func Barking() string {
	return "Woof! Woof! Woof!"
}

func BarkingLoud(s string) string {
	return dog.WhenGrownUp(s)
}
