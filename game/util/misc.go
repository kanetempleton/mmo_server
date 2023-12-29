// Package util provides utility functions.
package util

import (
	"math/rand"
	"time"
)

// Random returns a random integer between 0 and x inclusive.
func Random(x int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(x + 1)
}
