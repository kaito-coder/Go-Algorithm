package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxD := 0
	findHeight(root, &maxD)
	return maxD
}
func findHeight(root *TreeNode, maxD *int) int {
	if root == nil {
		return 0
	}
	left := findHeight(root.Left, maxD)
	right := findHeight(root.Right, maxD)
	localMax := left + right
	*maxD = max(*maxD, localMax)
	return max(left, right) + 1

}
func main() {

}
