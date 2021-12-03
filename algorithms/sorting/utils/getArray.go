package utils

import (
	"math/rand"
)

func GenerateArray(len int) []int {
	arrs := make([]int, len)

	for i := 0; i < len; i++ {
		rounds := rand.Intn(10000)
		arrs[i] = rounds
	}
	return arrs
}
