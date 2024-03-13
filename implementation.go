package lab2

import (
	"strings"
)

type BinaryTreeNode struct {
	Value string
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func ConstructBinaryTree(postfixExpr string) *BinaryTreeNode {
	stack := make([]*BinaryTreeNode, 0)

	tokens := strings.Fields(postfixExpr)

	for _, token := range tokens {
		node := &BinaryTreeNode{Value: token}

		if isOperator(token) {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			node.Left = left
			node.Right = right
		}
		stack = append(stack, node)
	}
	return stack[0]
}

func isOperator(token string) bool {
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	return operators[token]
}

func InorderTraversal(root *BinaryTreeNode) string {
	if root != nil {
		left := InorderTraversal(root.Left)
		right := InorderTraversal(root.Right)
		return left + root.Value + " " + right
	}
	return ""
}

// TODO: document this function.
// PostfixToInfix converts
func PostfixToInfix(expression string) (string, error) {
	root := ConstructBinaryTree(expression)
	result := InorderTraversal(root)
	return strings.TrimSpace(result), nil
}
