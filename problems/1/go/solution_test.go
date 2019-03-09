package twosum2

import (
	"reflect"
	"testing"
)

func TestInput1(t *testing.T) {
	answer := []int{0, 1}
	target := 9
	nums := []int{2, 7, 11, 15}
	result := solution(target, nums)
	t.Log(reflect.DeepEqual(result, answer))
}

func TestInput2(t *testing.T) {
	answer := []int{1, 2}
	target := 6
	nums := []int{3, 2, 4}
	result := solution(target, nums)
	t.Log(reflect.DeepEqual(result, answer))
}
