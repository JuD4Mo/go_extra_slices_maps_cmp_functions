package main

import (
	"fmt"
	"slices"
)

func main() {

	deleteTest := []int64{14, 315, 6362, 724742, 426724}

	// No modifica el slice original
	fmt.Printf("delete test variable: %v\n", deleteTest)
	fmt.Printf("Delete: %v\n", slices.Delete(deleteTest, 1, 4))

	//Modifica el slice original
	reverseTest := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("reverse test variable: %v\n", reverseTest)
	slices.Reverse(reverseTest)
	fmt.Printf("Reverse: %v\n", reverseTest)

	//Modifica el slice original
	sortTest := []float64{12, 3, 4, 5}
	fmt.Printf("sort test variable: %v\n", sortTest)
	slices.Sort(sortTest)
	fmt.Printf("Sort: %v\n", sortTest)

	// No modifica el slice original
	replaceTest := []int64{31, 642, 7, 3, 724, 7, 6}
	fmt.Printf("replace test variable: %v\n", replaceTest)
	replaceTest = slices.Replace(replaceTest, 3, 6, 62, 27, 27)
	fmt.Printf("replace: %v\n", replaceTest)

	// No modifica el slice original
	cloneTest := []int{3, 46, 24, 762, 47, 4}
	fmt.Printf("clone test variable: %v\n", cloneTest)
	clonedTest := slices.Clone(cloneTest)
	cloneTest[2] = 5
	fmt.Printf("Clone: %v & %v\n", cloneTest, clonedTest)

	compareTest1 := []int{2, 399, 4, 56, 7}
	compareTest2 := []int{2, 200, 4, 56, 7}

	ans := slices.Compare(compareTest1, compareTest2)
	fmt.Printf("Compare: %d - %v & %v\n", ans, compareTest1, compareTest2)
	/*
		si a == b ; ans = 0
		si a < b ; ans = -1
		si a > b ; ans = 1

		si intentamos comparar 2 slices con diferente tamaño,
		la comparación se ira haciendo igualmente uno a uno,
		y si no encuentra diferencia tomara al slice con menor
		tamaño como menor.
	*/

	//Max
	fmt.Printf("Max: %d\n", slices.Max(compareTest1))

	//Min
	fmt.Printf("Min: %d\n", slices.Min(compareTest1))

	/*Binary Search
	Ordenamiento ascendente
	*/
	binaryTest := []int{1, 2456, 27, 47, 547, 84, 3, 7, 42, 6, 72, 52, 12, 4}
	slices.Sort(binaryTest)
	index, is := slices.BinarySearch(binaryTest, 7)
	fmt.Printf("Binary search - found: %t | position: %d - %v\n", is, index, binaryTest)

	/* Compact
	Ordenamiento ascendente
	*/

	compacTest := []int{1, 1, 1, 2, 2, 2, 2, 3, 5, 53, 6, 16, 1, 3, 5, 356, 6, 4}
	slices.Sort(compacTest)
	fmt.Printf("Compact: %v\n", compacTest)
	fmt.Printf("Compact: %v\n", slices.Compact(compacTest))
}
