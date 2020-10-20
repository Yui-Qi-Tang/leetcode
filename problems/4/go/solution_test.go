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

func TestSol(t *testing.T) {

	t.Log("Start testing binary search method...")

	testForms := []TestForm{
		{
			nums1:  []int{1, 2, 3, 4, 20},
			nums2:  []int{5, 6, 7, 8, 9},
			answer: 5.5,
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			answer: 2.5,
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3},
			answer: 2.0,
		},
		{
			nums1:  []int{0, 0},
			nums2:  []int{0},
			answer: 0.0,
		},
		{
			nums1:  []int{0, 0},
			nums2:  []int{0, 0},
			answer: 0.0,
		},
	}

	for i, testForm := range testForms {
		t.Logf("test case: %d...", i)

		//findMedianSortedBinarySearch(testForm.nums1, testForm.nums2)
		t.Log(findMedianSortedArrays(testForm.nums1, testForm.nums2))

		/*
			if result != testForm.answer {
				t.Fatalf("incorrect result: %0.f(answer: %0.f)\n", result, testForm.answer)
			}*/
		//t.Log("... ok")
	}

	t.Log(" ...Passed")
}
