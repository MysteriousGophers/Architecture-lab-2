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
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
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

// PostfixToInfix converts a postfix expression to an infix expression.
// It takes a postfix expression as input, constructs a binary tree
// based on the expression, performs an inorder traversal on the tree,
// and returns the resulting infix expression.
//
// Parameters:
//   - expression: A string representing the postfix expression to be converted.
//
// Returns:
//   - string: The resulting infix expression.
//   - error: An error is returned if there is an issue with the conversion process.
//
// Example:
//
//	postfixExpression := "3 4 + 2 *"
//	infixExpression, err := PostfixToInfix(postfixExpression)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Infix Expression:", infixExpression)
//	}
func PostfixToInfix(expression string) (string, error) {
	root := ConstructBinaryTree(expression)
	result := InorderTraversal(root)
	return strings.TrimSpace(result), nil
}
