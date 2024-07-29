package domutil

import (
	"html/template"
	"strings"

	"github.com/choonkeat/dom-go"
)

// ReplaceAll allows user content (not html) to be replaced with a html node, while preserving
// the original content. e.g. replacing all occurrences of "{your email}" with a
// `<a href="mailto:...">...</a>` html.
//
// It looks for all occurrences of matchText as InnerText, or escaped(matchText) as InnerHTML,
// or recursively in all children of the target node.
//
// If we want to search and replace html, we can just use strings.ReplaceAll(target.HTML(), ...)
func ReplaceAll(target dom.Node, matchText string, node dom.Node) dom.Node {
	switch {
	case target.InnerHTML != "":
		parts := strings.Split(string(target.InnerHTML), template.HTMLEscapeString(matchText))
		if len(parts) == 1 {
			return target
		}
		newParts := make([]dom.Node, 0, len(parts)+len(parts)-1)
		for i, part := range parts {
			if i != 0 {
				newParts = append(newParts, node)
			}
			newParts = append(newParts, dom.InnerHTML(part))
		}
		return Join(newParts...)
	case target.InnerText != "":
		parts := strings.Split(string(target.InnerText), matchText)
		if len(parts) == 1 {
			return target
		}
		newParts := make([]dom.Node, 0, len(parts)+len(parts)-1)
		for i, part := range parts {
			if i != 0 {
				newParts = append(newParts, node)
			}
			newParts = append(newParts, dom.InnerText(part))
		}
		return Join(newParts...)
	default:
		newChildren := make([]dom.Node, 0, len(target.Children))
		for _, child := range target.Children {
			newChildren = append(newChildren, ReplaceAll(child, matchText, node))
		}
		target.Children = newChildren
	}
	return target
}
