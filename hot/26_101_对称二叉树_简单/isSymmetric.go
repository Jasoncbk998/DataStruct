/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/26 2:31 下午
 **/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
		1
	2		2
3	  4   4    3
*/
// 从根节点开始,像左和向右递归
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}
