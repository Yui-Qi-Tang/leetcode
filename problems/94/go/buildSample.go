package solution

// TreeNode tree node
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 這是個模擬，因為問題沒有寫到建立樹的性質。
// HINT: 建立樹不是從頭開始
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
	return head
}
