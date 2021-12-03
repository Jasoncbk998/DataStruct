/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 12:44 下午
 **/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true
*/

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}
