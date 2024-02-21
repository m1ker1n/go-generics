package generics

import (
	"slices"
	"testing"
)

func TestMapKeys(t *testing.T) {
	type args[M ~map[K]V, K comparable, V any] struct {
		m M
	}
	type testCase[M ~map[K]V, K comparable, V any] struct {
		name string
		args args[M, K, V]
		want []K
	}
	tests := []testCase[map[string]int, string, int]{
		{
			name: "ok",
			args: args[map[string]int, string, int]{
				m: map[string]int{
					"one":   1,
					"two":   2,
					"three": 3,
				},
			},
			want: []string{"one", "two", "three"},
		},
		{
			name: "empty map",
			args: args[map[string]int, string, int]{
				m: map[string]int{},
			},
			want: []string{},
		},
		{
			name: "nil map",
			args: args[map[string]int, string, int]{
				m: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeys(tt.args.m)
			if len(got) != len(tt.args.m) {
				t.Errorf("len(got) != len(m): %d != %d", len(got), len(tt.args.m))
			}
			for key := range tt.args.m {
				if !slices.Contains(got, key) {
					t.Errorf("MapKeys() does not contain %v", key)
				}
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	type args[M ~map[K]V, K comparable, V any] struct {
		m M
	}
	type testCase[M ~map[K]V, K comparable, V any] struct {
		name string
		args args[M, K, V]
		want []V
	}
	tests := []testCase[map[string]int, string, int]{
		{
			name: "ok",
			args: args[map[string]int, string, int]{
				m: map[string]int{
					"one":   1,
					"two":   2,
					"three": 3,
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "empty map",
			args: args[map[string]int, string, int]{
				m: map[string]int{},
			},
			want: []int{},
		},
		{
			name: "nil map",
			args: args[map[string]int, string, int]{
				m: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapValues(tt.args.m)
			if len(got) != len(tt.args.m) {
				t.Errorf("len(got) != len(m): %d != %d", len(got), len(tt.args.m))
			}
			for _, v := range tt.args.m {
				if !slices.Contains(got, v) {
					t.Errorf("MapKeys() does not contain %v", v)
				}
			}
		})
	}
}
