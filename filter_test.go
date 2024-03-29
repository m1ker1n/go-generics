package generics

import (
	"reflect"
	"testing"
)

type filterArgument[T any] struct {
	slice     []T
	predicate Predicate[T]
}

type filterTestCase[T any] struct {
	name string
	arg  filterArgument[T]
	want []T
}

func TestFilter(t *testing.T) {
	passAll := func(i int) bool { return true }
	notPassAll := func(i int) bool { return false }
	isOdd := func(i int) bool { return i%2 == 1 }
	tests := []filterTestCase[int]{
		{
			name: "ok: is odd",
			arg: filterArgument[int]{
				slice:     []int{1, 2, 3},
				predicate: isOdd,
			},
			want: []int{1, 3},
		},
		{
			name: "ok: is odd for slices with even nums",
			arg: filterArgument[int]{
				slice:     []int{2, 4, 6},
				predicate: isOdd,
			},
			want: []int{},
		},
		{
			name: "nil slice => nil result",
			arg: filterArgument[int]{
				slice:     nil,
				predicate: isOdd,
			},
			want: nil,
		},
		{
			name: "empty slice => empty result",
			arg: filterArgument[int]{
				slice:     []int{},
				predicate: isOdd,
			},
			want: []int{},
		},
		{
			name: "ok: pass all",
			arg: filterArgument[int]{
				slice:     []int{2, 4, 5, 6},
				predicate: passAll,
			},
			want: []int{2, 4, 5, 6},
		},
		{
			name: "ok: not pass all",
			arg: filterArgument[int]{
				slice:     []int{2, 4, 5, 6},
				predicate: notPassAll,
			},
			want: []int{},
		},
	}

	for i, tt := range tests {
		got := Filter(tt.arg.slice, tt.arg.predicate)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("[%d] [%s] got != want: %#v != %#v", i, tt.name, got, tt.want)
		}
	}

	func() {
		nilPredicateTC := filterTestCase[int]{
			name: "nil predicate => panic",
			arg: filterArgument[int]{
				slice:     []int{1, 2, 3},
				predicate: nil,
			},
			want: nil,
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("[%s] expected panic", nilPredicateTC.name)
			}
		}()
		Filter(nilPredicateTC.arg.slice, nilPredicateTC.arg.predicate)
	}()
}

type findFirstArgument[T any] struct {
	slice     []T
	predicate Predicate[T]
}

type findFirstWant[T any] struct {
	val T
	ok  bool
}

type findFirstTestCase[T any] struct {
	name string
	arg  findFirstArgument[T]
	want findFirstWant[T]
}

func TestFindFirst(t *testing.T) {
	passAll := func(i int) bool { return true }
	notPassAll := func(i int) bool { return false }
	isOdd := func(i int) bool { return i%2 == 1 }
	tests := []findFirstTestCase[int]{
		{
			name: "ok: found odd number",
			arg: findFirstArgument[int]{
				slice:     []int{1, 2, 3},
				predicate: isOdd,
			},
			want: findFirstWant[int]{
				val: 1,
				ok:  true,
			},
		},
		{
			name: "ok: odd number is not found in even numbers",
			arg: findFirstArgument[int]{
				slice:     []int{2, 4, 6},
				predicate: isOdd,
			},
			want: findFirstWant[int]{
				ok: false,
			},
		},
		{
			name: "nil slice => not found",
			arg: findFirstArgument[int]{
				slice:     nil,
				predicate: isOdd,
			},
			want: findFirstWant[int]{
				ok: false,
			},
		},
		{
			name: "empty slice => not found",
			arg: findFirstArgument[int]{
				slice:     []int{},
				predicate: isOdd,
			},
			want: findFirstWant[int]{
				ok: false,
			},
		},
		{
			name: "predicate returns true => returns first element of slice",
			arg: findFirstArgument[int]{
				slice:     []int{2, 4, 5, 6},
				predicate: passAll,
			},
			want: findFirstWant[int]{
				val: 2,
				ok:  true,
			},
		},
		{
			name: "predicate returns false => not found",
			arg: findFirstArgument[int]{
				slice:     []int{2, 4, 5, 6},
				predicate: notPassAll,
			},
			want: findFirstWant[int]{
				ok: false,
			},
		},
	}

	for i, tt := range tests {
		el, ok := FindFirst(tt.arg.slice, tt.arg.predicate)
		got := findFirstWant[int]{
			val: el,
			ok:  ok,
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("[%d] [%s] got != want: %#v != %#v", i, tt.name, got, tt.want)
		}
	}

	func() {
		nilPredicateTC := findFirstTestCase[int]{
			name: "nil predicate => panic",
			arg: findFirstArgument[int]{
				slice:     []int{1, 2, 3},
				predicate: nil,
			},
			want: findFirstWant[int]{},
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("[%s] expected panic", nilPredicateTC.name)
			}
		}()
		FindFirst(nilPredicateTC.arg.slice, nilPredicateTC.arg.predicate)
	}()
}
