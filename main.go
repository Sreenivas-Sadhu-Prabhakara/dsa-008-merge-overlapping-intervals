package main

import (
	"fmt"
	"sort"
)

// MergeIntervals merges all overlapping [start, end] intervals and returns
// the result sorted by start. Touching intervals ([1,3] and [3,5]) merge too.
// The input slice is re-ordered (sorted) in the process.
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	out := [][]int{{intervals[0][0], intervals[0][1]}}
	for _, iv := range intervals[1:] {
		last := out[len(out)-1]
		if iv[0] <= last[1] {
			if iv[1] > last[1] {
				last[1] = iv[1]
			}
		} else {
			out = append(out, []int{iv[0], iv[1]})
		}
	}
	return out
}

func main() {
	fmt.Println(MergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
