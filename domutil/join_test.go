package domutil_test

import (
	"html/template"
	"testing"

	"github.com/choonkeat/dom-go"
	"github.com/choonkeat/dom-go/domutil"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		nodes []dom.Node
		want  template.HTML
	}{
		{
			nodes: []dom.Node{
				dom.Div(dom.Attrs("class", "first"), dom.InnerText("Hello, ")),
				dom.Strong(dom.Attrs("class", "second"), dom.InnerText("world")),
				dom.InnerText("! <em>Goodbye</em>"),
				dom.InnerHTML(" <span>world</span>"),
			},
			want: template.HTML(`<div class="first">Hello, </div><strong class="second">world</strong>! &lt;em&gt;Goodbye&lt;/em&gt; <span>world</span>`),
		},
	}
	for _, tt := range tests {
		got := domutil.Join(tt.nodes...).HTML()
		if got != tt.want {
			t.Errorf("\ngot      %q\nbut want %q", got, tt.want)
		}
	}
}
