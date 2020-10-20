package solution

// TODO: try to better solution

/*
    Given two sorted arrays nums1 and nums2 of size m and n respectively,
    return the median of the two sorted arrays.

	Follow up: The overall run time complexity should be O(log (m+n))

	Median examples:

	a=[1,3,3,6,7,8,9] the median = 6 (a[len(a)/2] if len(a)%2=1))

	a=[1,2,3,4,5,6,8,9] the median = (4+5)/2 = 4.5 (if len(a)%2=0, (a[len(a)/2]+a[len(a)/2-1])/2)
*/

func findMedianSortedBinarySearch(nums1 []int, nums2 []int) float64 {
	/*
		partition := (len(nums1) + len(nums2) + 1) / 2

		partitionNums1 := (len(nums1) - 1) / 2
		partitionNums2 := partition - partitionNums1*/

	return 0.0
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

func merge(n1, n2 []int) []int {
	if len(n1) == 0 && len(n2) == 0 {
		return []int{}
	}

	m, n := len(n1), len(n2)
	result := make([]int, m+n)

	n1 = append(n1, 1<<63-1)
	n2 = append(n2, 1<<63-1)

	i, j, k := 0, 0, 0
	for k < m+n {
		if n1[i] <= n2[j] {
			result[k] = n1[i]
			i++
		} else {
			result[k] = n2[j]
			j++
		}
		k++
	}

	return result
}
