package solution

import (
	"github.com/golang-collections/collections/stack"
)

func outOfRange(x int) bool {
	if x > 1<<31-1 || x < -1<<31 {
		return true
	}
	return false
}

// index, cover
func pivot(i int) (int, bool) {

	if i%2 == 0 {
		return i / 2, false
	}

	return i/2 + 1, true
}

func isPalindrome(n int) bool {
	s := stack.New()

	if outOfRange(n) {
		return false
	}

	if n < 0 {
		return false
	}

	if n < 10 {
		return true
	}

	for n != 0 {
		s.Push(n % 10)
		n /= 10
	}

	nArr := make([]int, s.Len())
	for i := range nArr {
		nArr[i] = s.Pop().(int)
	}

	p, cover := pivot(len(nArr))

	n2Arr := make([]int, 0)
	n2Arr = append(n2Arr, nArr[:p]...)

	if cover {
		nArr = nArr[p-1:]
	} else {
		nArr = nArr[p:]
	}

	for _, v := range n2Arr {
		s.Push(v)
	}

	for _, v := range nArr {
		v2 := s.Pop()
		if v != v2.(int) {
			return false
		}
	}

	return true

}

func isPalindrome2(x int) bool {

	if x > 1<<31-1 || x < -1<<31 {
		return false
	}

	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	arr := make([]int, 0)

	for x != 0 {
		arr = append(arr, x%10)
		x /= 10
	}

	p, cover := pivot(len(arr))

	a1 := make([]int, 0)
	a2 := make([]int, 0)
	a1 = arr[:p]
	if cover {
		a2 = arr[p-1:]
	} else {
		a2 = arr[p:]
	}

	for i, v := range a2 {
		if v != a1[len(a1)-i-1] {
			return false
		}
	}

	return true

}

func isPalindrome3(x int) bool {
	if x > 1<<31-1 || x < -1<<31 {
		return false
	}

	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	reverse := 0

	for x > reverse {
		r := x % 10
		reverse = (reverse * 10) + r
		x /= 10
	}

	//     len(x)%2 ==0 , len(x)%2 !=0
	return x == reverse || x == reverse/10
}
