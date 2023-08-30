package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {

	deleteTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Delete Test variable: %v\n", deleteTest)
	fmt.Printf("Delete: %v\n", slices.Delete(deleteTest, 1, 2))

	reverseTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Reverse Test variable: %v\n", reverseTest)
	slices.Reverse(reverseTest)
	fmt.Printf("Reverse: %v\n", reverseTest)

	sortTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Sort Test variable: %v\n", sortTest)
	slices.Sort(sortTest)
	fmt.Printf("Sort: %v\n", sortTest)

	replaceTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Replace Test variable: %v\n", replaceTest)
	replaceTest = slices.Replace(replaceTest, 3, 3, 10, 11, 12)
	fmt.Printf("Replace: %v\n", replaceTest)

	cloneTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Replace Test variable: %v\n", cloneTest)
	clonedTest := slices.Clone(cloneTest)
	fmt.Printf("Cloned: %v\n", clonedTest)

	compare1 := []int{3, 2, 4, 3, 1, 2, 4, 6}
	compare2 := []int{3, 2, 4, 3, 1, 2, 4, 6}

	result := slices.Compare(compare1, compare2)
	fmt.Printf("Compare : %d - %v & %v\n", result, compare1, compare2)

	maxTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Max: %d\n", slices.Max(maxTest))

	minTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Min: %d\n", slices.Min(minTest))

	binarySrcTest := []int{3, 2, 4, 7, 3, 1, 2, 4, 6}
	slices.Sort(binarySrcTest)

	i, found := slices.BinarySearch(binarySrcTest, 7)
	fmt.Printf("Binary search - found: %t | position: %d - %v\n", found, i, binarySrcTest)

	compactTest := []int{1, 1, 2, 2, 9, 9, 3, 3, 2, 1, 10, 10, 5, 1}
	slices.Sort(compactTest)
	fmt.Printf("Compact Test: %v\n", compactTest)
	fmt.Printf("Compact: %v\n", slices.Compact(compactTest))

	indexTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Index: %d\n", slices.Index(indexTest, 6))
	fmt.Printf("Index: %d\n", slices.Index(indexTest, 2))
	fmt.Printf("Index: %d\n", slices.Index(indexTest, 9))

	isSortedTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("isSorted: %t\n", slices.IsSorted(isSortedTest))
	slices.Sort(isSortedTest)
	fmt.Printf("isSorted: %t\n", slices.IsSorted(isSortedTest))

	insertTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Insert: %v\n", slices.Insert(insertTest, 4, 10, 20, 22))

	containsTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Contains: %v\n", slices.Contains(containsTest, 4))

	equalTest := []int{3, 2, 4, 3, 1, 2, 4, 6}
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{1, 2, 4, 6}))
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{5, 6}))
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{10, 200, 30}))
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{3, 2, 4, 3, 1, 2, 4, 6}))

	numbers := []int{0, 42, 8}
	strings := []string{"000", "42", "8"}

	equal := slices.EqualFunc(numbers, strings, func(i int, s string) bool {
		sn, err := strconv.ParseInt(s, 0, 64)

		if err != nil {
			return false
		}

		return i == int(sn)
	})

	fmt.Printf("Equal Func: %t\n", equal)
}
