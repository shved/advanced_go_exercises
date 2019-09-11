package main

import (
	"fmt"
	"os"
	"strings"
)

// Example of valid input:
// *bold*
// _italic_
// ~strikethroughed~
// *surprisingly*_works_
// * surprisingly also works*
// *and ~also~ _works_ like that*
// *and ~_also_ will~ work like* that
//
var closingFlags = make(map[NodeKind]bool)

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
			stack.Push(Node{kind: Bold, contents: string(char), closingFlag: toggleClosingFlag(Bold)})
		case '_':
			stack.Push(Node{kind: Italic, contents: string(char), closingFlag: toggleClosingFlag(Italic)})
		case '~':
			stack.Push(Node{kind: Strikethrough, contents: string(char), closingFlag: toggleClosingFlag(Strikethrough)})
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

// Toggles closinf flag value for each tag type
func toggleClosingFlag(t NodeKind) bool {
	if _, ok := closingFlags[t]; ok == false {
		closingFlags[t] = true
	}
	closingFlags[t] = !closingFlags[t]
	return closingFlags[t]
}
