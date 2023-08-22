package main

import (
	"unicode"
)


type TokenType int

const (
	Return TokenType = iota
	IntLiteral
	OpenParen
	CloseParen
)

type Token struct {
	Type TokenType
	Value string
}


func Tokenize(str string) []Token {
	var tokens []Token
	var buffer string
	var i int

	peek := func() string {
		if i < len(str) {
			return string(str[i])
		}
		return ""
	}

	consume := func() rune {
		i++
		return rune(str[i-1])
	}

	for peekValue := peek(); peekValue != ""; peekValue = peek() {
		if unicode.IsLetter(rune(peekValue[0])) {
			buffer = string(consume())
			for next := peek(); next != "" && unicode.IsLetter(rune(next[0])); {
				buffer += string(consume())
				next = peek()
			}

			if buffer == "exit" {
				tokens = append(tokens, Token{Type: Return})
			}
			buffer = ""
		} else if unicode.IsDigit(rune(peekValue[0])) {
			buffer = string(consume())
			for next := peek(); next != "" && unicode.IsDigit(rune(next[0])); {
				buffer += string(consume())
				next = peek()
			}
			tokens = append(tokens, Token{Type: IntLiteral, Value: buffer})
			buffer = ""
		} else if peekValue == "(" {
			consume()
			tokens = append(tokens, Token{Type: OpenParen})
		} else if peekValue == ")" {
			consume()
			tokens = append(tokens, Token{Type: CloseParen})
		} else if peekValue == " " {
			consume()
		} else {
			consume() 
		}
	}

	return tokens
}


