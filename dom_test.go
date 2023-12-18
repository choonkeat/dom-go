package dom_test

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	"github.com/choonkeat/dom-go"
)

func TestAttribute(t *testing.T) {
	got := dom.Div(
		dom.Attrs(
			"class", "greeting",
			"style", "color: red;",
		),
		dom.InnerText("Hello, world!"),
		dom.InnerHTML("<strong>Hello!</strong>"),
	)
	want := dom.Node{
		Name: "div",
		Attributes: []dom.Attribute{
			{Name: "class", ValueText: "greeting"},
			{Name: "style", ValueText: "color: red;"},
		},
		Children: []dom.Node{
			{InnerText: "Hello, world!"},
			{InnerHTML: "<strong>Hello!</strong>"},
		},
	}
	if got.HTML() != want.HTML() {
		t.Fatalf("want %#v but got %#v", want.HTML(), got.HTML())
	}
}

func TestInput(t *testing.T) {
	t.Parallel()

	got := dom.Input(dom.Attrs("name", "username"))
	want := template.HTML(`<input name="username"/>`)

	if got.HTML() != want {
		t.Fatalf("want %#v but got %#v", want, got.HTML())
	}
}

func TestEmptyElement(t *testing.T) {
	t.Parallel()

	got :=
		dom.Element(
			"", dom.Attrs(),
			dom.InnerText("Good "),
			dom.B(dom.Attrs(), dom.InnerText("morning")),
			dom.InnerText(", world!"),
		)

	want := template.HTML(`Good <b>morning</b>, world!`)

	if got.HTML() != want {
		t.Fatalf("want %#v but got %#v", want, got.HTML())
	}

}

func Example() {
	fmt.Println(
		// use dom.Element or dom.Div
		dom.Element("div",
			dom.Attrs(
				"class", "1 2 3",
				"data-foo", `4<'"5"'>6`,
			),
			dom.InnerText("<oops>789</oops>"),
			dom.P(
				dom.Attrs(),
				dom.InnerHTML("<strong>10</strong>"),
			),
		).HTML(),
	)
	// Output: <div class="1 2 3" data-foo="4&lt;&#39;&#34;5&#34;&#39;&gt;6">&lt;oops&gt;789&lt;/oops&gt;<p><strong>10</strong></p></div>
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
			dom.InnerText("Goo<g>le"),
			dom.InnerHTML("Goo<g>le"),
			dom.Blockquote(
				dom.Attrs(),
				dom.InnerText("Google"),
			),
		).HTML(),
	)
	// Output: <a href="https://google.com" target="_blank">Goo&lt;g&gt;leGoo<g>le<blockquote>Google</blockquote></a>
}

func ExampleInnerText() {
	fmt.Println(
		dom.InnerText("Goo<g>le").HTML(),
	)
	// Output: Goo&lt;g&gt;le
}

func ExampleInnerHTML() {
	fmt.Println(
		dom.InnerHTML("Goo<g>le").HTML(),
	)
	// Output: Goo<g>le
}

func BenchmarkAttrHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dom.Attrs(
			"href", "https://google.com",
		)[0].HTML()
	}
	b.ReportAllocs()
	b.ReportMetric(float64(b.N), "AttrHTML")
}

func BenchmarkNodeHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dom.A(
			dom.Attrs(
				"href", "https://google.com",
				"target", "_blank",
			),
			dom.InnerText("Goo<g>le"),
			dom.Blockquote(
				dom.Attrs(),
				dom.InnerText("Google"),
			),
		).HTML()
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
	b.ResetTimer()
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
