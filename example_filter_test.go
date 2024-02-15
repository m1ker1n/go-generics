package generics

import (
	"fmt"
	"unicode/utf8"
)

func ExampleFilter_getOddNums() {
	isOdd := func(i int) bool { return i%2 == 1 }
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%#v", Filter(x, isOdd))
	// Output: []int{1, 3, 5}
}

func ExampleFilter_getShortWords() {
	isShorterThanOrEqual := func(runesInString int) Predicate[string] {
		return func(x string) bool {
			return utf8.RuneCountInString(x) <= runesInString
		}
	}
	x := []string{"catJAM", "monkaS", "OMEGALUL", "Clap", "KEKW", "EZ",
		"大黑公鸡", "国王和傻瓜是最好的乐队"}
	fmt.Printf("%#v", Filter(x, isShorterThanOrEqual(4)))
	// Output: []string{"Clap", "KEKW", "EZ", "大黑公鸡"}
}
