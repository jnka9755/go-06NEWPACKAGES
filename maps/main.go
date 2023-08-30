package main

import (
	"fmt"
	"maps"
)

func main() {

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	myOtherMap := maps.Clone(myMap)
	fmt.Printf("Clone: %v\n", myOtherMap)

	fmt.Printf("Equal: %t\n", maps.Equal(myMap, myOtherMap))

	maps.DeleteFunc(myMap, func(k string, v int) bool {
		return k == "one"
	})

	fmt.Printf("With deleted value %v\n", myMap)

	myOtherMap["fourteen"] = 14
	myOtherMap["nine"] = 9
	myMap["ten"] = 10

	fmt.Printf("First Map: %v - Clone %v\n", myMap, myOtherMap)

	maps.Copy(myMap, myOtherMap)

	fmt.Printf("First Map: %v - Clone %v\n", myMap, myOtherMap)
}
