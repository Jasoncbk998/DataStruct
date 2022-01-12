/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/23 9:17 下午
 **/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
简单题。采用深搜的思路，分别从根节点开始同时遍历两个二叉树，并将对应的节点进行合并。两个二叉树的对应节点可能存在以下三种情况：
如果两个二叉树的对应节点都为空，则合并后的二叉树的对应节点也为空；
如果两个二叉树的对应节点只有一个为空，则合并后的二叉树的对应节点为其中的非空节点；
如果两个二叉树的对应节点都不为空，则合并后的二叉树的对应节点的值为两个二叉树的对应节点的值之和，此时需要显性合并两个节点。
对一个节点进行合并之后，还要对该节点的左右子树分别进行合并。用递归实现即可。
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1

}
