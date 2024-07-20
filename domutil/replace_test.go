package domutil_test

import (
	"html/template"
	"testing"

	"github.com/choonkeat/dom-go"
	"github.com/choonkeat/dom-go/domutil"
)

func TestReplaceAll(t *testing.T) {
	replaceNode := dom.B(dom.Attrs("class", "abc def123"), dom.InnerText("universe"))

	tests := []struct {
		given dom.Node
		match string
		want  template.HTML
	}{
		{
			given: dom.InnerHTML("hello <em>&lt;strong&gt;world&lt;/strong&gt;!</em>"),
			match: "<strong>world</strong>",
			want:  `hello <em><b class="abc def123">universe</b>!</em>`,
		},
		{
			given: dom.InnerText("hello <em><strong>world</strong>!</em>"),
			match: "<strong>world</strong>",
			want:  `hello &lt;em&gt;<b class="abc def123">universe</b>!&lt;/em&gt;`,
		},
		{
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
		domutil.ReplaceAll(&test.given, test.match, replaceNode)
		actual := test.given
		if got, want := actual.HTML(), test.want; got != want {
			t.Errorf("\ngot      %q\nbut want %q", got, want)
		}
		if oldHTML == test.given.HTML() {
			t.Errorf("ReplaceAll did not modify the given node")
		}
	}
}
