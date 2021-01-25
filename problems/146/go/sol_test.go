package sol

import (
	"testing"
)

/*
["LRUCache","get","put","get","put","put","get","get"]
[[2],       [2],  [2,6],[1],  [1,5],[1,2],[1],  [2]]

2,6
*/

func TestBasic(t *testing.T) {
	c := Constructor(2)
	t.Log(c.Get(2))
	t.Log(c.order, c.data)

	c.Put(2, 6)
	t.Log(c.order, c.data)

	t.Log(c.Get(1))
	t.Log(c.order, c.data)

	c.Put(1, 5)
	t.Log(c.order, c.data)

	c.Put(1, 2)
	t.Log(c.order, c.data)

	t.Log(c.Get(1))
	t.Log(c.order, c.data)

	t.Log(c.Get(2))
	t.Log(c.order, c.data)

}

/*
["LRUCache","put","put","put","put","get","get"]
[[2],       [2,1],[1,1], [2,3],[4,1],[1],[2]]


[null,null,null,null,null,-1,3]
*/

func TestBasic2(t *testing.T) {
	c := Constructor(2)
	c.Put(2, 1)
	t.Log(c.order, c.data)

	c.Put(1, 1)
	t.Log(c.order, c.data)

	c.Put(2, 3)
	t.Log(c.order, c.data)

	c.Put(4, 1)
	t.Log(c.order, c.data)

	t.Log(c.Get(1))
	t.Log(c.order, c.data)

	t.Log(c.Get(2))
	t.Log(c.order, c.data)

}

/*
["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],       [1,1],[2,2],[1],  [3,3],[2],  [4,4],[1],  [3],  [4]]
Output
[null,null,null,1,null,2,null,1,3,4]
Expected
[null,null,null,1,null,-1,null,-1,3,4]
*/

func TestBasic3(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	t.Log(c.order, c.data)

	c.Put(2, 2)
	t.Log(c.order, c.data)

	t.Log(c.Get(1))
	t.Log(c.order, c.data)

	c.Put(3, 3)
	t.Log(c.order, c.data)

	t.Log(c.Get(2))
	t.Log(c.order, c.data)

	c.Put(4, 4)
	t.Log(c.order, c.data)

	t.Log(c.Get(1))
	t.Log(c.order, c.data)

	t.Log(c.Get(3))
	t.Log(c.order, c.data)

	t.Log(c.Get(4))
	t.Log(c.order, c.data)

}
