package dom_test

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"testing"

	"github.com/choonkeat/dom-go"
)

func TestAttribute(t *testing.T) {
	var gotHTML, wantHTML strings.Builder
	dom.Div(
		dom.Attrs(
			"class", "greeting",
			"style", "color: red;",
		),
		dom.Text("Hello, world!"),
	).HTML(&gotHTML)
	dom.Node{
		Name: "div",
		Attributes: []dom.Attribute{
			{Name: "class", ValueText: "greeting"},
			{Name: "style", ValueText: "color: red;"},
		},
		Children: []dom.Node{
			{InnerText: "Hello, world!"},
		},
	}.HTML(&wantHTML)
	if gotHTML.String() != wantHTML.String() {
		t.Fatalf("want %#v but got %#v", wantHTML.String(), gotHTML.String())
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
		).HTML(nil).String(),
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
		).HTML(nil).String(),
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
		).HTML(nil).String(),
	)
	// Output: <a href="https://google.com" target="_blank">Goo&lt;g&gt;le<blockquote>Google</blockquote></a>
}

func ExampleText() {
	fmt.Println(
		dom.Text("Goo<g>le").HTML(nil).String(),
	)
	// Output: Goo&lt;g&gt;le
}

func BenchmarkAttrHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		dom.Attrs(
			"href", "https://google.com",
		)[0].HTML(&sb).String()
	}
	b.ReportAllocs()
	b.ReportMetric(float64(b.N), "AttrHTML")
}

func BenchmarkNodeHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		dom.A(
			dom.Attrs(),
			dom.Text("Goo<g>le"),
			dom.Blockquote(
				dom.Attrs(),
				dom.Text("Google"),
			),
		).HTML(&sb).String()
	}
	b.ReportAllocs()
	b.ReportMetric(float64(b.N), "NodeHTML")
}

func BenchmarkHtmlTemplate(b *testing.B) {
	tmpl, err := template.New("index.html").Parse(`<a href="{{ .Href }}" target="{{ .Target }}">{{ .Text1 }}<blockquote>{{ .Text2 }}</blockquote></a>`)
	if err != nil {
		b.Fatal(err)
	}
	type data struct {
		Href   string
		Target string
		Text1  string
		Text2  string
	}
	for i := 0; i < b.N; i++ {
		d := data{
			Href:   "https://google.com",
			Target: "_blank",
			Text1:  "Goo<g>le",
			Text2:  "Google",
		}
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "index.html", d); err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
	b.ReportMetric(float64(b.N), "HTMLTemplate")
}
