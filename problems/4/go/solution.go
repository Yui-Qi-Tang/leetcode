package solution

import (
	"math/rand"
	"time"
)

// TODO: try to better solution

/*
    Given two sorted arrays nums1 and nums2 of size m and n respectively,
    return the median of the two sorted arrays.

	Follow up: The overall run time complexity should be O(log (m+n))

	Median examples:

	a=[1,3,3,6,7,8,9] the median = 6 (a[len(a)/2] if len(a)%2=1))

	a=[1,2,3,4,5,6,8,9] the median = (4+5)/2 = 4.5 (if len(a)%2=0, (a[len(a)/2]+a[len(a)/2-1])/2)
*/

// binary search (submit this version for leetcode)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	n1Len := len(nums1)
	n2Len := len(nums2)
	merged := make([]int, n1Len+n2Len)
	end := 1<<63 - 1
	// add end to nums1 & nums2
	nums1 = append(nums1, end)
	nums2 = append(nums2, end)

	i, j, mIndex := 0, 0, 0
	for mIndex < len(merged) {
		if nums1[i] <= nums2[j] {
			merged[mIndex] = nums1[i]
			i++
		} else {
			merged[mIndex] = nums2[j]
			j++
		}
		mIndex++
	}

	if len(merged)%2 == 1 {
		return float64(merged[len(merged)/2])
	}

	m1 := float64(merged[len(merged)/2])
	m2 := float64(merged[len(merged)/2-1])

	return (m1 + m2) / 2.0
}

func findMedianSortedArraysBruteForce(nums1 []int, nums2 []int) float64 {
	// Complexity : O(m+n)
	// Space: O(m+n)
	var merged []int
	var foo int
	foo = 1
	int64Max := foo<<63 - 1

	leftPart := len(nums1)
	rightPart := len(nums2)

	nums1 = append(nums1, int64Max)
	nums2 = append(nums2, int64Max)

	i := 0
	j := 0
	c := 0
	for c < leftPart+rightPart {
		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
		c++
	}

	if len(merged)%2 == 0 {
		middleIndex := len(merged) / 2
		minMedianIndex := middleIndex - 1
		m1 := float64(merged[middleIndex])
		m2 := float64(merged[minMedianIndex])

		return float64((m1 + m2) / 2.0)
	}

	return float64(merged[len(merged)/2.0])
}

func findMedianSortedArraysQuickSelect(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	return medianQuickSelect(nums1)
}

func medianQuickSelect(input []int) float64 {
	if len(input)%2 == 1 {
		return float64(quickSelect(input, len(input)/2))
	}
	return (float64(quickSelect(input, len(input)/2-1)) + float64(quickSelect(input, len(input)/2))) / 2.0
}

func randPivot(i []int) int {
	rand.Seed(time.Now().UnixNano())
	return i[rand.Intn(len(i))]
}

// k: index
func quickSelect(l []int, k int) int {
	if len(l) == 1 && k == 0 {
		return l[0]
	}

	pivot := randPivot(l)

	// partition
	lows := make([]int, 0)   // forall x in l, x < pivot
	highs := make([]int, 0)  // forall x in l, x > pivot
	pivots := make([]int, 0) // forall x in l , x == pivot
	for _, v := range l {
		if v < pivot {
			lows = append(lows, v)
		}
		if v > pivot {
			highs = append(highs, v)
		}
		if v == pivot {
			pivots = append(pivots, v)
		}
	}

	if k < len(lows) {
		return quickSelect(lows, k)
	} else if k < len(lows)+len(pivots) {
		return pivots[0]
	} else {
		return quickSelect(highs, k-len(lows)-len(pivots))
	}
}
