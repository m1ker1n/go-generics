package generics

import (
	"fmt"
	"strconv"
)

func ExampleMap_intToInt() {
	doubleInt := func(i int) (int, error) {
		return 2 * i, nil
	}
	x := []int{1, 2, 4}
	fmt.Printf("%#v", Map(x, doubleInt))
	// Output: []int{2, 4, 8}
}

func ExampleMap_intToString() {
	itoa := func(i int) (string, error) {
		return strconv.Itoa(i), nil
	}
	x := []int{1, 2, 3}
	fmt.Printf("%#v", Map(x, itoa))
	// Output: []string{"1", "2", "3"}
}

func ExampleMap_stringToInt() {
	x := []string{"1", "2", "3", "wtf"}
	// As strconv.Atoi("wtf") returns error, this value will be skipped in result.
	fmt.Printf("%#v", Map(x, strconv.Atoi))
	// Output: []int{1, 2, 3}
}

func ExampleMap_stringToIntToInt() {
	doubleInt := func(i int) (int, error) { return 2 * i, nil }
	x := []string{"1", "2", "3", "wtf"}
	// As strconv.Atoi("wtf") returns error, this value will be skipped in result.
	fmt.Printf("%#v", Map(Map(x, strconv.Atoi), doubleInt))
	// Output: []int{2, 4, 6}
}
