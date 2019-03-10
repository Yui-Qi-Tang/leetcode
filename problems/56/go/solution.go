package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type interval struct {
	start int
	end   int
}

type Interval struct {
	Start int
	End   int
}

// use sort
func merge2(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	// Let intervals sort
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
	ret := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		tail := ret[len(ret)-1]
		if tail.End < intervals[i].Start {
			ret = append(ret, intervals[i])
		} else if intervals[i].End > tail.End {
			ret[len(ret)-1].End = intervals[i].End
		}
	}
	return ret
}

// without sort
func merge(intervals []interval) []interval {
	for i := 0; i < len(intervals); i++ {
		next := i + 1
		if next >= len(intervals) {
			break
		}
		fmt.Println(i, next)
		if isOverlap(intervals[i], intervals[next]) {
			if intervals[i].start < intervals[next].start {
				intervals[next].start = intervals[i].start
			}
			intervals = intervals[1:]
			i = -1 // reset start index
		}

	}
	return intervals
}

func isOverlap(i1, i2 interval) bool {
	// assume r in (start, end) is  start < end, implice interval = end - start > 0
	/*
		    i1 is overlay i2 if

			case1:

		         s     e
			i1:  |-----|
			           s    e
			i2:        |----|

				i1.e >= i2.s(i2.e > i2.s thus i2.e > i1.e)  p1

			or case2:
			    s     e
			i1: |-----|
			       s       e
			i2:    |-------|

				i1.e > i2.s && i1.e < i2.e  p1

			combin p1 and p2:
				i1.e >= i2.s && i1.e < i2.e

			i2.s <= i1.e < i2.e

			!!!property is not complete!!!
	*/
	if i1.end >= i2.start && i1.end <= i2.end && i1.start <= i2.start || i1.start >= i2.start {
		return true
	}
	return false
}

func main() {
	// [[1,3],[2,6],[8,10],[15,18]]
	t := []interval{
		{1, 4},
		{4, 5},
		{8, 10},
		{15, 18},
	}

	t2 := []Interval{
		{1, 3},
		{2, 5},
		{5, 10},
		{18, 10},
	}

	t = merge(t) // invalid version
	fmt.Println(t)

	t2 = merge2(t2) // bug in {1, 10} ,{18, 10}, but can solve problem on leetcode, haha
	fmt.Println(t2)

}
