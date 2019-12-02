package solution

func findMedianSortedBinarySearch(nums1 []int, nums2 []int) float64 {
	partition := (len(nums1) + len(nums2) + 1) / 2

	partitionNums1 := (len(nums1) - 1) / 2
	partitionNums2 := partition - partitionNums1

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
