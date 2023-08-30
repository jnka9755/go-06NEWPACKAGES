package main

import (
	"cmp"
	"fmt"
	"math"
)

func main() {
	var myFloatValue1 float64 = 20
	var myFloatValue2 float64 = 10
	myNanValue := math.NaN()

	fmt.Printf("Numbers (%f, %f) - Compare: %d\n", myFloatValue1, myFloatValue2, cmp.Compare(myFloatValue2, myFloatValue1))
	fmt.Printf("Numbers (%f, %f) - Compare: %t\n", myFloatValue1, myFloatValue2, cmp.Less(myFloatValue2, myFloatValue1))

	fmt.Printf("Numbers (%f, %f) - Compare: %d\n", myFloatValue1, myNanValue, cmp.Compare(myFloatValue2, myFloatValue1))
	fmt.Printf("Numbers (%f, %f) - Compare: %t\n", myFloatValue1, myNanValue, cmp.Less(myFloatValue2, myNanValue))

}
