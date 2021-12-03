/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/23 9:06 下午
 **/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
简单题。遍历每个节点的左子树和右子树，累加从左子树到右子树的最大长度。遍历每个节点时，动态更新这个最大长度即可
*/
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	checkDiameter(root, &result)
	return result
}

func checkDiameter(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := checkDiameter(root.Left, result)
	right := checkDiameter(root.Right, result)
	*result = max(*result, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
