package main

import (
	"fmt"
)

type NodeKind int

const (
	Content NodeKind = iota
	Bold
	Italic
	Strikethrough
)

type Node struct {
	opening  bool
	kind     NodeKind
	contents string
}

type nodeStack []Node

var inputExample = `
*bold*
_italic_
~strikethroughed~
* doesn't work*
*also ~doesn't~ _work_ like that*
`

func main() {
	// var md string
	// md = strings.Join(os.Args[1:], " ")

	var stack nodeStack
	parse(inputExample, &stack)
	fmt.Println(stack)
}

func parse(text string, stack *nodeStack) {
	for _, char := range text {
		switch char {
		case '*':
			stack.Push(Node{kind: Bold, contents: char})
		case '_':
			stack.Push(Node{kind: Italic, contents: char})
		case '~':
			stack.Push(Node{kind: Strikethrough, contents: char})
		default:
			stack.Push(Node{kind: Content, contents: char})
		}
	}
}

func (ns *nodeStack) Push(n Node) {
	// 0 len
	if len(ns) < 1 {
		ns = append(ns, n)
		return
	}

	switch {
	case n.kind == Content && ns.Peek().kind == Content:
		return
	case n.kind == Content && ns.Peek().kind != Content:
		return
	case n.kind != Content && ns.Peek().kind == Content:
		return
	case n.kind != Content && ns.Peek().kind != Content:
		return
	}

	// // append new content tag
	// if ns.Peek().kind != Content {
	// 	ns = append(ns, n)
	// 	return
	// }
	//
	// // concat with last content
	// if ns.Peek().kind == Content {
	// 	lastElem := ns.Pop()
	// 	lastElem.contents = fmt.Sprint(lastElem.contents, n.contents)
	// 	ns = append(ns, lastElem)
	// }
	// тут должны редьюситься все невалидные теги и копиться контент
	// апдейтиться opening поля тегов должны второй итерацией
}

func (ns *nodeStack) Peek() Node {
	if len(ns) > 0 {
		return ns[len(ns)-1]
	}

	return Node{}
}

func (ns *nodeStack) Pop() (n Node) {
	if len(ns) > 0 {
		lastElem, ns := ns[len(ns)-1], ns[:len(ns)-1]
		return lastElem
	}

	return Node{}
}

func (ns *nodeStack) LastTagKind() int {
	tags := ns.FilterTags
	if len(tags) > 0 {
		return tags.Peek().kind
	}

	return -1
}

func (ns *nodeStack) FilterTags() []Node {
	var tagNodes nodeStack
	for _, node := range ns {
		if node.kind != Content {
			tagNodes = append(tagNodes, node)
		}
	}

	return tagNodes
}
