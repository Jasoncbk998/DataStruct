/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/16 13:25
 **/
package main

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}
