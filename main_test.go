package main

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		in, want [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{4, 5}, {1, 4}}, [][]int{{1, 5}}},
		{[][]int{{1, 10}, {2, 3}}, [][]int{{1, 10}}},
		{[][]int{{1, 2}}, [][]int{{1, 2}}},
		{[][]int{}, [][]int{}},
	}
	for _, c := range cases {
		if got := MergeIntervals(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
