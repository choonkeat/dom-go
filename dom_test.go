package dom_test

import (
	"fmt"
	"testing"

	"github.com/choonkeat/dom-go"
)

func TestAttribute(t *testing.T) {
	got := dom.Div(
		dom.Attrs(
			"class", "greeting",
			"style", "color: red;",
		),
		dom.Text("Hello, world!"),
	)
	want := dom.Node{
		Name: "div",
		Attributes: []dom.Attribute{
			{Name: "class", ValueText: "greeting"},
			{Name: "style", ValueText: "color: red;"},
		},
		Children: []dom.Node{
			{InnerText: "Hello, world!"},
		},
	}
	if got.HTML() != want.HTML() {
		t.Fatalf("want %#v but got %#v", want.HTML(), got.HTML())
	}
}

func Example() {
	fmt.Println(
		// use dom.Element or dom.Div
		dom.Element("div",
			dom.Attrs("class", "1 2 3", "data-foo", `4<'"5"'>6`),
			dom.Text("<oops>789</oops>"),
			dom.Strong(
				dom.Attrs(),
				dom.Text("10"),
			),
		).HTML(),
	)
	// Output: <div class="1 2 3" data-foo="4&lt;&#39;&#34;5&#34;&#39;&gt;6">&lt;oops&gt;789&lt;/oops&gt;<strong>10</strong></div>
}

func ExampleAttrs() {
	fmt.Println(
		dom.Div(
			dom.Attrs(
				"href", "https://google.com",
				"target", "_blank",
			),
		).HTML(),
	)
	// Output: <div href="https://google.com" target="_blank"></div>
}

func ExampleElement() {
	fmt.Println(
		dom.Element("a",
			dom.Attrs(
				"href", "https://google.com",
				"target", "_blank",
			),
			dom.Text("Goo<g>le"),
			dom.Blockquote(
				dom.Attrs(),
				dom.Text("Google"),
			),
		).HTML(),
	)
	// Output: <a href="https://google.com" target="_blank">Goo&lt;g&gt;le<blockquote>Google</blockquote></a>
}

func ExampleText() {
	fmt.Println(
		dom.Text("Goo<g>le").HTML(),
	)
	// Output: Goo&lt;g&gt;le
}