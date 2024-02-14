package generics

import (
	"errors"
	"reflect"
	"testing"
)

type mapArgument[From, To any] struct {
	slice    []From
	callBack Transformation[From, To]
}

type mapTestCase[From, To any] struct {
	name string
	args mapArgument[From, To]
	want []To
}

type tryMapWant[To any] struct {
	slice       []To
	expectError bool
}

type tryMapTestCase[From, To any] struct {
	name string
	args mapArgument[From, To]
	want tryMapWant[To]
}

func TestMap(t *testing.T) {
	doubleInt := func(i int) (int, error) {
		return 2 * i, nil
	}

	set5 := func(i int) (int, error) {
		return 5, nil
	}

	tests := []mapTestCase[int, int]{
		{
			name: "double int slice",
			args: mapArgument[int, int]{
				slice:    []int{1, 2, 3},
				callBack: doubleInt,
			},
			want: []int{2, 4, 6},
		},
		{
			name: "set all elements to 5",
			args: mapArgument[int, int]{
				slice:    []int{1, 2, 3},
				callBack: set5,
			},
			want: []int{5, 5, 5},
		},
		{
			name: "nil x => nil result",
			args: mapArgument[int, int]{
				slice:    nil,
				callBack: set5,
			},
			want: nil,
		},
		{
			name: "empty x => empty result",
			args: mapArgument[int, int]{
				slice:    []int{},
				callBack: set5,
			},
			want: []int{},
		},
		{
			name: "cb can return error",
			args: mapArgument[int, int]{
				slice: []int{1, 2, 3, 4},
				callBack: func(i int) (int, error) {
					if i%2 == 1 {
						return 0, errors.New("skip odd")
					}
					return i, nil
				},
			},
			want: []int{2, 4},
		},
		{
			name: "cb returns only errors",
			args: mapArgument[int, int]{
				slice: []int{1, 2, 3, 4},
				callBack: func(i int) (int, error) {
					return 0, errors.New("skip everything")
				},
			},
			want: []int{},
		},
	}
	for i, tt := range tests {
		gotSlice := Map(tt.args.slice, tt.args.callBack)
		if !reflect.DeepEqual(gotSlice, tt.want) {
			t.Errorf("[%d] [%s] got != want: %#v != %#v", i, tt.name, gotSlice, tt.want)
		}
	}

	panicTest := mapTestCase[int, int]{
		name: "nil cb => panic",
		args: mapArgument[int, int]{
			slice:    []int{1, 2, 3, 4},
			callBack: nil,
		},
		want: nil,
	}
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("[%s] expect panic", panicTest.name)
			}
		}()
		Map(panicTest.args.slice, panicTest.args.callBack)
	}()
}

func TestTryMap(t *testing.T) {
	doubleInt := func(i int) (int, error) {
		return 2 * i, nil
	}

	set5 := func(i int) (int, error) {
		return 5, nil
	}

	tests := []tryMapTestCase[int, int]{
		{
			name: "double int slice",
			args: mapArgument[int, int]{
				slice:    []int{1, 2, 3},
				callBack: doubleInt,
			},
			want: tryMapWant[int]{
				slice:       []int{2, 4, 6},
				expectError: false,
			},
		},
		{
			name: "set all elements to 5",
			args: mapArgument[int, int]{
				slice:    []int{1, 2, 3},
				callBack: set5,
			},
			want: tryMapWant[int]{
				slice:       []int{5, 5, 5},
				expectError: false,
			},
		},
		{
			name: "nil x => nil result",
			args: mapArgument[int, int]{
				slice:    nil,
				callBack: set5,
			},
			want: tryMapWant[int]{
				slice:       nil,
				expectError: false,
			},
		},
		{
			name: "empty x => empty result",
			args: mapArgument[int, int]{
				slice:    []int{},
				callBack: set5,
			},
			want: tryMapWant[int]{
				slice:       []int{},
				expectError: false,
			},
		},
		{
			name: "cb can return error",
			args: mapArgument[int, int]{
				slice: []int{1, 2, 3, 4},
				callBack: func(i int) (int, error) {
					if i%2 == 0 {
						return 0, errors.New("skip even")
					}
					return i, nil
				},
			},
			want: tryMapWant[int]{
				slice:       []int{1},
				expectError: true,
			},
		},
		{
			name: "cb returns only errors",
			args: mapArgument[int, int]{
				slice: []int{1, 2, 3, 4},
				callBack: func(i int) (int, error) {
					return 0, errors.New("skip everything")
				},
			},
			want: tryMapWant[int]{
				slice:       []int{},
				expectError: true,
			},
		},
		{
			name: "nil cb",
			args: mapArgument[int, int]{
				slice:    []int{1, 2, 3, 4},
				callBack: nil,
			},
			want: tryMapWant[int]{
				slice:       nil,
				expectError: true,
			},
		},
	}
	for i, tt := range tests {
		gotSlice, err := TryMap(tt.args.slice, tt.args.callBack)
		if err == nil && tt.want.expectError {
			t.Errorf("[%d] [%s] expected error, got nil", i, tt.name)
		}

		if err != nil && !tt.want.expectError {
			t.Errorf("[%d] [%s] did not expect error, but got %s", i, tt.name, err)
		}

		if !reflect.DeepEqual(gotSlice, tt.want.slice) {
			t.Errorf("[%d] [%s] got != want: %#v != %#v", i, tt.name, gotSlice, tt.want)
		}
	}
}
