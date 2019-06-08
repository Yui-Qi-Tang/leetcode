package solution

//import "fmt"

// TreeNode tree node
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func buildTree(values ...int) *TreeNode {
	head := &TreeNode{}
	currentNode := head
    for i, v := range values {
        if i == 0{
			head.Val = v
			head.Left = nil
			head.Right = nil
		} else {
			newNode := &TreeNode{
				v,
				nil,
				nil,
			}
            if v % 2 == 0 {
                currentNode.Right = newNode
			} else {
                currentNode.Left = newNode
			}
			currentNode = newNode
		}		
	}
	//fmt.Println(head)
	return head
}
