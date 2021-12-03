/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/23 1:38 下午
 **/
package main

/**
中心思想
根节点的左字数反转二叉树，然后再和根节点的右节点进行比较

反转二叉树是第 226 题。判断 2 颗树是否完全相等是第 100 题。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)

}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(invertTree(root.Left), root.Right)
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
