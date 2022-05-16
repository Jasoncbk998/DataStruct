/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/16 11:03
 **/
package main

func preorderTraversal(root *TreeNode) []int {
	var result []int
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}
