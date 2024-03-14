package lab2

import (
	"fmt"
	"strings"
	"unicode"
)

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
func PostfixToInfix(input string) (string, error) {
	err := validate(input)
	if err != nil {
		return "", err
	}
	root := constructBinaryTree(input)
	result := inorderTraversal(root)
	return result, nil
}

func validate(input string) error {
	if input == "" {
		return fmt.Errorf("empty input")
	}
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("blank input")
	}
	if strings.ContainsFunc(input, func(char rune) bool {
		return !(unicode.IsNumber(char) ||
			unicode.IsSpace(char) ||
			isOperator(string(char)))
	}) {
		return fmt.Errorf("usupported input")
	}
	return nil
}

type BinaryTreeNode struct {
	Value string
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func constructBinaryTree(postfixExpr string) *BinaryTreeNode {
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

func getPriority(token string) int {
	switch token {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 10
}

func inorderTraversal(root *BinaryTreeNode) string {
	if root != nil {
		left := inorderTraversal(root.Left)
		right := inorderTraversal(root.Right)

		if root.Left != nil && getPriority(root.Value) > getPriority(root.Left.Value) {
			left = "(" + left + ")"
		}
		if root.Right != nil && getPriority(root.Value) > getPriority(root.Right.Value) {
			right = "(" + right + ")"
		}

		return strings.TrimSpace(left + " " + root.Value + " " + right)
	}
	return ""
}
