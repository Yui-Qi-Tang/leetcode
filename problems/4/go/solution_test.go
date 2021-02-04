package solution

import (
	"testing"
)

type TestForm struct {
	answer float64
	nums1  []int
	nums2  []int
}

var testForms []TestForm = []TestForm{
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

func TestBruteForce(t *testing.T) {
	t.Log("Start testing brute force method...")
	for _, testForm := range testForms {
		result := findMedianSortedArraysBruteForce(testForm.nums1, testForm.nums2)
		if result != testForm.answer {
			t.Fatalf("incorrect result: %0.f(answer: %0.f)\n", result, testForm.answer)
		}
	}
	t.Log(" ...Passed")
}

func TestBinarySearch(t *testing.T) {
	t.Log("Start testing binary search method...")
	for _, testForm := range testForms {
		result := findMedianSortedArrays(testForm.nums1, testForm.nums2)
		if result != testForm.answer {
			t.Fatalf("incorrect result: %0.f(answer: %0.f)\n", result, testForm.answer)
		}
	}
	t.Log(" ...Passed")
}

func TestQuickSelect(t *testing.T) {
	t.Log("Start testing quick method method...")
	for _, testForm := range testForms {
		result := findMedianSortedArraysQuickSelect(testForm.nums1, testForm.nums2)
		if result != testForm.answer {
			t.Fatalf("incorrect result: %0.f(answer: %0.f)\n", result, testForm.answer)
		}
	}
	t.Log(" ...Passed")
}
