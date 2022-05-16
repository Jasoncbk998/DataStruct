/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/16 11:03
 **/
package main

func postorderTraversal(root *TreeNode) []int {
	var result []int
	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}
