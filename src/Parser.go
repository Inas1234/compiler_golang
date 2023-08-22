package main

import (
	"fmt"
)

type NodeExpr struct {
	int_lit Token
}

type NodeExit struct {
	expr NodeExpr
}

func Parse(tokens []Token) (NodeExit, error) {
	var node NodeExit
	tokens, err := ParseReturn(tokens)
	if err != nil {
		return NodeExit{}, err
	}
	node.expr, tokens, err = ParseExpr(tokens)
	if err != nil {
		return NodeExit{}, err
	}
	return node, nil
}

func ParseReturn(tokens []Token) ([]Token, error) {
	if len(tokens) == 0 || tokens[0].Type != Return {
		return tokens, fmt.Errorf("expected Return token")
	}
	return tokens[1:], nil
}

func ParseExpr(tokens []Token) (NodeExpr, []Token, error) {
	if len(tokens) == 0 || tokens[0].Type != OpenParen {
		return NodeExpr{}, tokens, fmt.Errorf("expected OpenParen token")
	}
	if len(tokens) < 3 || tokens[1].Type != IntLiteral || tokens[2].Type != CloseParen {
		return NodeExpr{}, tokens, fmt.Errorf("expected IntLiteral and CloseParen tokens")
	}
	return NodeExpr{int_lit: tokens[1]}, tokens[3:], nil
}
