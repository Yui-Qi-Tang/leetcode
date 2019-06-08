package solution


/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
*/

// recursive visit sequence in tree: left -> current -> right
func inorderTraversal(root *TreeNode) []int {
    var results []int

    if root != nil {
        results = append(results, inorderTraversal(root.Left)...)
        results = append(results, root.Val)
        results = append(results, inorderTraversal(root.Right)...)
    }
    return results
}




