/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/10 1:38 下午
 **/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
输入：p = [1,2,3], q = [1,2,3]
输出：true
输入：p = [1,2], q = [1,null,2]
输出：false
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//都是nil 则true
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		//val不相等,则nil
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		//一个树 nil 另一个没有nil 则false
		return false
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
