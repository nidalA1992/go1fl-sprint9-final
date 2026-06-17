package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	m := map[int]int{
		0:    0,
		-1:   0,
		-100: 0,
		1:    1,
		100:  100,
	}

	for arg, exp := range m {
		slice := generateRandomElements(arg)
		assert.Len(t, slice, exp)
	}

	isSame := true
	slice := generateRandomElements(100)
	for i := 1; i < len(slice); i++ {
		if slice[0] != slice[i] {
			isSame = false
			break
		}
	}
	assert.False(t, isSame, "slice elements should be different")
}

func TestMaximum(t *testing.T) {
	m := map[int][]int{
		0:       []int{},
		2:       []int{2},
		3:       []int{1, 2, 3},
		5:       []int{3, 4, 5, 2, 1, 0},
		10:      []int{9, 1, 4, 10, 3, 5},
		-1:      []int{-10, -100, -34, -1, -34},
		9991243: []int{100, 341, 412345, 12342, 9991243, 0, 55544, 8345},
		//24
		23: []int{
			0, 1, 2, 3, 4, 5, 6, 7,
			8, 9, 10, 11, 12, 13, 14, 15,
			16, 17, 18, 19, 20, 21, 22, 23,
		},
		//16
		15: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		//17
		16: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		//18
		17: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
	}

	assert.Equal(t, 0, maximum(nil))
	for exp, sl := range m {
		assert.Equal(t, exp, maximum(sl))
	}
}
