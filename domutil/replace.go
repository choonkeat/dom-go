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
		target.InnerHTML = template.HTML(
			strings.ReplaceAll(
				string(target.InnerHTML),
				// we are replacing user text, not html
				// :. escape before matching
				template.HTMLEscapeString(matchText),
				string(node.HTML()),
			),
		)
	case target.InnerText != "":
		// we're replacing with HTML, so we need to escape the parts
		// and switch to InnerHTML
		parts := strings.Split(target.InnerText, matchText)
		if len(parts) == 1 {
			return target
		}
		for i := range parts {
			parts[i] = template.HTMLEscapeString(parts[i])
		}
		target.InnerHTML = template.HTML(strings.Join(parts, string(node.HTML())))
		// we've switched to representing with InnerHTML, so clear InnerText
		target.InnerText = ""
	default:
		newChildren := make([]dom.Node, 0, len(target.Children))
		for _, child := range target.Children {
			newChildren = append(newChildren, ReplaceAll(child, matchText, node))
		}
		target.Children = newChildren
	}
	return target
}
