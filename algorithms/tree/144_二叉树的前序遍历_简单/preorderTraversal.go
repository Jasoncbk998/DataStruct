/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/20 6:54 下午
 **/
package main

//递归
func preorderTraversal2(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

//模拟栈
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left //不断找左节点
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
