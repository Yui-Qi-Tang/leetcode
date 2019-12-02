package solution

import (
	"testing"
)

type TestForm struct {
	answer float64
	nums1  []int
	nums2  []int
}

func TestBruteForce(t *testing.T) {

	t.Log("Start testing brute force method...")

	testForms := []TestForm{
		TestForm{
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			answer: 2.5,
		},
		TestForm{
			nums1:  []int{1, 2},
			nums2:  []int{3},
			answer: 2.0,
		},
	}

	for i, testForm := range testForms {
		t.Logf("test case: %d...", i)

		result := findMedianSortedArraysBruteForce(testForm.nums1, testForm.nums2)

		if result != testForm.answer {
			t.Fatalf("incorrect result: %0.f(answer: %0.f)\n", result, testForm.answer)
		}
		t.Log("... ok")
	}

	// 	result := findMedianSortedArraysBruteForce(nums1, nums2)

	t.Log(" ...Passed")
}
