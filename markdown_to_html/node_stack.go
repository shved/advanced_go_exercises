package main

import (
	"errors"
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
	closingFlag bool
	kind        NodeKind
	contents    string
}

type nodeStack struct {
	nodes  []Node
	errors []error
}

func validateStack(stack *nodeStack) {
	var nodesByTypes = make(map[NodeKind]map[bool]uint)

	for _, node := range stack.nodes {
		if nodesByTypes[node.kind] == nil {
			nodesByTypes[node.kind] = make(map[bool]uint)
		}
		nodesByTypes[node.kind][node.closingFlag]++
	}

	if nodesByTypes[Bold][false] != nodesByTypes[Bold][true] {
		stack.errors = append(stack.errors, errors.New("invalid bold tags"))
	}

	if nodesByTypes[Italic][false] != nodesByTypes[Italic][true] {
		stack.errors = append(stack.errors, errors.New("invalid italic tags"))
	}

	if nodesByTypes[Strikethrough][false] != nodesByTypes[Strikethrough][true] {
		stack.errors = append(stack.errors, errors.New("invalid strikethrougs tags"))
	}
}

func (ns *nodeStack) Push(n Node) {
	// 0 len
	if len(ns.nodes) < 1 {
		ns.nodes = append(ns.nodes, n)
		return
	}

	if n.kind == Content && ns.Peek().kind == Content {
		// reduce following content tokens into one content node
		lastElem := ns.Pop()
		lastElem.contents = fmt.Sprint(lastElem.contents, n.contents)
		ns.nodes = append(ns.nodes, lastElem)
	} else {
		ns.nodes = append(ns.nodes, n)
	}
}

func (ns *nodeStack) Peek() Node {
	if len(ns.nodes) > 0 {
		return ns.nodes[len(ns.nodes)-1]
	}

	return Node{}
}

func (ns *nodeStack) Pop() (n Node) {
	if len(ns.nodes) > 0 {
		var lastElem Node
		lastElem, ns.nodes = ns.nodes[len(ns.nodes)-1], ns.nodes[:len(ns.nodes)-1]
		return lastElem
	}

	return Node{}
}
