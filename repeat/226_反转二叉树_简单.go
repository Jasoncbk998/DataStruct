/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/17 11:43
 **/
package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
