/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/16 13:06
 **/
package main

// https://leetcode.cn/problems/same-tree/
// 两个树 节点分别对应
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
