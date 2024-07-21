package domutil_test

import (
	"html/template"
	"testing"

	"github.com/choonkeat/dom-go"
	"github.com/choonkeat/dom-go/domutil"
)

func TestReplaceAll(t *testing.T) {
	// <b class="text-xs">universe</b>
	replaceNode := dom.B(dom.Attrs("class", "text-xs"), dom.InnerText("universe"))

	tests := []struct {
		given dom.Node
		match string
		want  template.HTML
	}{
		{
			// simple scenario showing off the basic scenario this function is designed for
			given: dom.InnerText("hello world"),
			match: "world",
			want:  `hello <b class="text-xs">universe</b>`,
		},
		{
			// the `match` text is escaped before matching InnerHTML
			given: dom.InnerHTML("hello <em>&lt;strong&gt;world&lt;/strong&gt;!</em> <strong>world</strong>!"),
			match: "<strong>world</strong>",
			want:  `hello <em><b class="text-xs">universe</b>!</em> <strong>world</strong>!`,
		},
		{
			// the `match` text matches InnerText as-is
			given: dom.InnerText("hello <em>&lt;strong&gt;world&lt;/strong&gt;!</em> <strong>world</strong>!"),
			match: "<strong>world</strong>",
			want:  `hello &lt;em&gt;&amp;lt;strong&amp;gt;world&amp;lt;/strong&amp;gt;!&lt;/em&gt; <b class="text-xs">universe</b>!`,
		},
		{
			// we recursively look into all children of the target node
			given: dom.Div(
				dom.Attrs(
					"class", "my-world",
				),
				dom.P(
					dom.Attrs(),
					dom.InnerText("hello <em>world!</em>"),
					dom.InnerHTML(", or <span>world</span>"),
				),
				dom.InnerText("hello world?"),
				dom.InnerHTML(", or <em>world</em>"),
			),
			match: "world",
			want: dom.Div(
				dom.Attrs(
					"class", "my-world",
				),
				dom.P(
					dom.Attrs(),
					dom.InnerText("hello <em>"),
					replaceNode, // inserted here
					dom.InnerText("!</em>"),
					dom.InnerHTML(", or <span>"),
					replaceNode, // inserted here
					dom.InnerHTML("</span>"),
				),
				dom.InnerText("hello "),
				replaceNode, // inserted here
				dom.InnerText("?"),
				dom.InnerHTML(", or <em>"),
				replaceNode, // inserted here
				dom.InnerHTML("</em>"),
			).HTML(),
		},
	}

	for _, test := range tests {
		oldHTML := test.given.HTML()
		actual := domutil.ReplaceAll(test.given, test.match, replaceNode)
		if got, want := actual.HTML(), test.want; got != want {
			t.Errorf("\ngot      %q\nbut want %q", got, want)
		}
		if oldHTML != test.given.HTML() {
			t.Errorf("ReplaceAll modified the given node")
		}
	}
}
