package domutil

import "github.com/choonkeat/dom-go"

func Join(nodes ...dom.Node) dom.Node {
	return dom.Node{Children: nodes}
}
