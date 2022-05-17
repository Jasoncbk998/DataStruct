/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/24 3:04 下午
 **/
package main

import (
	"DataStruct/tools"
	"math"
)

func isBalanced(root *tools.TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return math.Abs(float64(leftHeight-rightHeight)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 不断递归,求得深度
func depth(root *tools.TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(depth(root.Left), depth(root.Right)) + 1
}
