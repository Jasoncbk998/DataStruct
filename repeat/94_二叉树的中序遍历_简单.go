/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/16 11:03
 **/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

/**
按照访问左子树——根节点——右子树的方式遍历这棵树，
而在访问左子树或者右子树的时候我们按照同样的方式遍历，直到遍历完整棵树。
*/
func inorder(root *TreeNode, output *[]int) {
	// 停止条件
	if root != nil {
		//进入递归
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
