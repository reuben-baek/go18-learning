package generic

import (
	"fmt"
	"testing"
)

func TestGenericMap(t *testing.T) {
	intMap := map[int]int{
		1: 1,
		2: 2,
	}
	PrintMap(intMap)

	fmt.Printf("value from map: %d: %v\n", 1, ValueFromMap(intMap, 1))

	stringMap := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	PrintMap(stringMap)
	fmt.Printf("value from map: %v: %v\n", "k1", ValueFromMap(stringMap, "k1"))
}

func PrintMap[K comparable, V any](m map[K]V) {
	for k, v := range m {
		fmt.Printf("key[%v] value[%v]\n", k, v)
	}
}

func ValueFromMap[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}
