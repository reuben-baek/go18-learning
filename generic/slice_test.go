package generic

import (
	"testing"
)

func TestGenericSlice(t *testing.T) {
	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice([]string{"a", "b", "c", "d"})
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		print(v)
	}
	println()
}
