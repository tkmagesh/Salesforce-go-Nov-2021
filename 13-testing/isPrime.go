package main

import "math"

func IsPrime(no int) bool {
	if no < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(no))); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
