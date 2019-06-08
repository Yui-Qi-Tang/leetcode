package solution

import (
	"testing"
	"reflect"
	//"fmt"
)

/*
func showTree(root *TreeNode)  {
    if root == nil {
		fmt.Println("null")
	}

	fmt.Printf("%d\n", root.Val)	
	if root.Left != nil {
        fmt.Printf("<-")
		showTree(root.Left)
	}
	if root.Right != nil {
		fmt.Printf("->")
		showTree(root.Right)
	}

}
*/


func TestBuildTree(t *testing.T){
	t.Log("Build tree...")
	input := []int{1,2,3,7,8,10}
	buildTree(input...)
    // showTree(r)
	t.Log(" ...PASS")
	
}



func TestSolution(t *testing.T) {
	t.Log("Start testing solution... ")
	input := []int{1,2,3}
	answer := []int{1,3,2}
	root := buildTree(input...)
	result := inorderTraversal(root)
	if reflect.DeepEqual(result, answer) {
		t.Log(" ...PASS")
	} else {
		t.Fatal("Error")
		t.Fatal("Answer:", answer)
		t.Fatal("Your result:", result)
	}
	// t.Log(reflect.DeepEqual(result, answer))
}