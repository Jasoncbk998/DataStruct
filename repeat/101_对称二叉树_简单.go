/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/7 12:10 下午
 **/
package main

import (
	"DataStruct/tools"
)

// 对称二叉树,可以理解为对称
func isSymmetric(root *tools.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// 左边的值等于右边的值
// 左子树向左向右寻找,右子树向左向右寻找
func isMirror(left *tools.TreeNode, right *tools.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isSymmetric_(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMoor(root.Left, root.Right)
}

func isMoor(left *TreeNode, right *TreeNode) bool {
	// 两个都是nil
	if left == nil && right == nil {
		return true
	}
	// 一个是nil,另一个不是nil
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isMoor(left.Left, right.Right) && isMoor(left.Right, right.Left)
}
