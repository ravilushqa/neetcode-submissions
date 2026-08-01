/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    inOrder := dfsInOrderRecursive(root, nil)

	for i, v := range inOrder {
		if i == 0 {
			continue
		}

		if v <= inOrder[i-1] {
			return false
		}
	}

	return true
}

func dfsInOrderRecursive(node *TreeNode, list []int) []int {
	if node == nil {
		return nil
	}

	if node.Left != nil {
		list = dfsInOrderRecursive(node.Left, list)
	}

	list = append(list, node.Val)

	if node.Right != nil {
		list = dfsInOrderRecursive(node.Right, list)
	}

	return list
}
