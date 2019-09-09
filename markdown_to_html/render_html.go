package main

import "bytes"

func renderHtml(nodes []Node) string {
	var tagMapping = make(map[NodeKind]map[bool]string)

	tagMapping[Bold] = make(map[bool]string)
	tagMapping[Bold][false] = "<strong>"
	tagMapping[Bold][true] = "</strong>"

	tagMapping[Italic] = make(map[bool]string)
	tagMapping[Italic][false] = "<em>"
	tagMapping[Italic][true] = "</em>"

	tagMapping[Strikethrough] = make(map[bool]string)
	tagMapping[Strikethrough][false] = "<del>"
	tagMapping[Strikethrough][true] = "</del>"

	var resultingHtml bytes.Buffer

	for _, node := range nodes {
		if node.kind != Content {
			node.contents = tagMapping[node.kind][node.closingFlag]
		}

		resultingHtml.WriteString(node.contents)
	}

	return resultingHtml.String()
}
