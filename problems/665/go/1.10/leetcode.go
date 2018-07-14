package main

import "fmt"

func checkPossibility(nums []int) bool {
	// Time complexity: O(n)
    numLen := len(nums) - 1
    var ptr interface{} = nil
    var isInside = func(index int) bool {
        if index < 0 || index > numLen {
            return false
        }
        return true
    }
    for i:=0; i<=numLen; i++ {
        j:= i+1
        // clear idea if occur nums[i] > nums[i+1] twice, just return
        if isInside(j) {
            if nums[i] > nums[j] {
                if ptr != nil {
                    return false
                }
                ptr = i // record the invalid condition occured on index i
            }
        }
    }
    // fmt.Println(ptr)
    // find the valid of nums[ptr-1] <= nums[ptr] <= nums[ptr+1] <= nums[ptr+2] if possibly
    // for now, we have lamma1 : nums[i] > nums[i+1]
    //p is None or p == 0 or p == len(A)-2 or A[p-1] <= A[p+1] or A[p] <= A[p+2]
    //return ptr == nill || ptr == 0, len
    if ptr == nil {
        return true
    }

    // if ptr = 0
    if ptr == 0 {
        return true
    }

    // if ptr = numLen - 2
    if ptr == len(nums) - 2 { //目前的邊界檢查沒問題，用numLen會有問題
        return true
    }

    // otherwise

    //(A[p-1] <= A[p+1] or A[p] <= A[p+2])
    recordIndex := ptr.(int)
    if nums[recordIndex - 1] <= nums[recordIndex+1] || nums[recordIndex] <= nums[recordIndex+2] {
        return true
    }
    return false // default
}
func main() {
	testCases := [][]int{
		{4, 2, 1},
		{2, 4, 3},
		{3, 4, 2, 3},
		{2, 3, 3, 2, 4},
		{1},
		{4, 2, 3},
		// custome your case here
		// ...
	}
	for _, v := range testCases {
	    fmt.Printf("Test case: %v, result: %t\n", v, checkPossibility(v))
	}
}
