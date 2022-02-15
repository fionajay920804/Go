package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// var name = "Leo"
	// name1 := "Sandra"
	// i := 6
	// j := i
	// i = 5
	// c := &i
	// fmt.Println("Hello World!")
	// fmt.Println("I can do it!")
	// fmt.Println(name)
	// fmt.Println(name1)
	// fmt.Println(i)
	// fmt.Println(j)
	// fmt.Println(c)
	root := TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}
	fmt.Println(isValidBST(&root))

}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
