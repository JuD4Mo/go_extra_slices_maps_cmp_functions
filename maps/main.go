package main

import (
	"fmt"
	"maps"
)

func main() {
	mapita := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	fmt.Printf("Values: %v\n", maps.Values(mapita))

	fmt.Printf("Keys: %v\n", maps.Keys(mapita))

	otherMap := maps.Clone(mapita)
	fmt.Printf("First map: %v - Clone %v\n", mapita, otherMap)

	maps.DeleteFunc(mapita, func(k string, v int) bool {
		return k == "one"
	})
	fmt.Printf("Deleted map: %v\n", mapita)

	otherMap["four"] = 44
	otherMap["nine"] = 9
	mapita["ten"] = 10
	fmt.Printf("First map: %v - Clone %v\n", mapita, otherMap)

	maps.Copy(mapita, otherMap)
	fmt.Printf("First map: %v - Clone %v\n", mapita, otherMap)
}
