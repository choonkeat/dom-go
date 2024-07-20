package domutil

import (
	"html/template"
	"strings"

	"github.com/choonkeat/dom-go"
)

// ReplaceAll looks for all occurrences of the string s in the target node
// and replaces them with the node. It does not return a value, the target
// node is modified in place.
func ReplaceAll(target *dom.Node, match string, node dom.Node) {
	switch {
	case target.InnerHTML != "":
		target.InnerHTML = template.HTML(strings.ReplaceAll(string(target.InnerHTML), match, string(node.HTML())))
	case target.InnerText != "":
		parts := strings.Split(target.InnerText, match)
		if len(parts) == 1 {
			return
		}
		for i := range parts {
			parts[i] = template.HTMLEscapeString(parts[i])
		}
		target.InnerHTML = template.HTML(strings.Join(parts, string(node.HTML())))
		target.InnerText = ""
	default:
		for i, child := range target.Children {
			ReplaceAll(&child, match, node)
			target.Children[i] = child
		}
	}
}
