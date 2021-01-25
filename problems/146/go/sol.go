package sol

/*
1 <= capacity <= 3000
0 <= key <= 3000
0 <= value <= 104
At most 3 * 104 calls will be made to get and put.


use link list

*/

type LRUCache struct {
	data  map[int]int
	cap   int
	order []int // [key1, k2, k3...]
}

func Constructor(capacity int) LRUCache {
	return LRUCache{data: make(map[int]int, capacity), cap: capacity}
}

func (lr *LRUCache) Get(key int) int {
	if v, ok := lr.data[key]; ok {
		for i, ov := range lr.order { // O(n)
			if ov == key {
				lr.order = append(lr.order[:i], lr.order[i+1:]...)
				lr.order = append([]int{key}, lr.order...)
			}
		}
		return v
	}
	return -1
}

func (lr *LRUCache) Put(key int, value int) {
	if len(lr.data) < lr.cap {
		lr.data[key] = value
		for _, v := range lr.order { // O(n)
			if v == key {
				return
			}
		}
		lr.order = append(lr.order, key)
		if len(lr.order) > 0 {
			//lr.order[0], lr.order[len(lr.order)-1] = lr.order[len(lr.order)-1], lr.order[0]
			lr.order = append([]int{key}, lr.order[:len(lr.order)-1]...)
		}
	} else {

		if _, exist := lr.data[key]; exist {
			lr.data[key] = value
		}

		delete(lr.data, lr.order[len(lr.order)-1])

		lr.order = append([]int{key}, lr.order[:len(lr.order)-1]...)

	}
}
