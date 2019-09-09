package main

import (
	"fmt"
	"os"
	"strings"
)

// var inputExample = `
// *bold*
// _italic_
// ~strikethroughed~
// *surprisingly*_works_
// * surprisingly also works*
// *and ~also~ _works_ like that*
// `

var closingFlag bool = true

func main() {
	var md string
	md = strings.Join(os.Args[1:], " ")

	var stack nodeStack
	parse(md, &stack)
	fmt.Println(renderHtml(stack.nodes))
}

func parse(text string, stack *nodeStack) {
	for _, char := range text {
		switch char {
		case '*':
			stack.Push(Node{kind: Bold, contents: string(char), closingFlag: toggleClosingFlag()})
		case '_':
			stack.Push(Node{kind: Italic, contents: string(char), closingFlag: toggleClosingFlag()})
		case '~':
			stack.Push(Node{kind: Strikethrough, contents: string(char), closingFlag: toggleClosingFlag()})
		default:
			stack.Push(Node{kind: Content, contents: string(char)})
		}
	}

	validateStack(stack)

	if len(stack.errors) > 0 {
		for _, err := range stack.errors {
			fmt.Errorf("%v\n", err)
		}
		os.Exit(1)
	}
}

func toggleClosingFlag() bool {
	closingFlag = !closingFlag	
	return closingFlag
}

