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
// or recursively in all children of the target node. It does not return a value since the
// target node is modified in place.
func ReplaceAll(target *dom.Node, matchText string, node dom.Node) {
	switch {
	case target.InnerHTML != "":
		target.InnerHTML = template.HTML(
			strings.ReplaceAll(
				string(target.InnerHTML),
				template.HTMLEscapeString(matchText),
				string(node.HTML()),
			),
		)
	case target.InnerText != "":
		// we're replacing with HTML, so we need to escape the parts
		// and switch to InnerHTML
		parts := strings.Split(target.InnerText, matchText)
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
			ReplaceAll(&child, matchText, node)
			target.Children[i] = child
		}
	}
}
