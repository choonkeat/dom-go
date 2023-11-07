// Construct HTML elements in Go. Can be used together with `html/template` templates.
//
// Use Element, Attrs, and Text. There are also helpers for every html element
//
// The structs are fully exposed for the convenience of asserting values during tests
// but should not be used directly otherwise.
//
// Example:
//
//	got := dom.Div(
//		dom.Attrs(
//			"class", "greeting",
//			"style", "color: red;",
//		),
//		dom.Text("Hello, world!"),
//	)
//	want := dom.Node{
//		Name: "div",
//		Attributes: []dom.Attribute{
//			{Name: "class", ValueText: "greeting"},
//			{Name: "style", ValueText: "color: red;"},
//		},
//		Children: []dom.Node{
//			{InnerText: "Hello, world!"},
//		},
//	}
//	if got.HTML() != want.HTML() {
//		t.Fatalf("want %#v but got %#v", want.HTML(), got.HTML())
//	}
package dom

import (
	"html/template"
	"strings"
)

// Attrs is the defacto helper function to provide `[]Attribute` to a `Node`
func Attrs(keyvalues ...string) []Attribute {
	var attrs []Attribute
	kvlen := len(keyvalues)
	for i := 0; i < kvlen; i += 2 {
		attrs = append(attrs, Attribute{
			Name:      keyvalues[i],
			ValueText: keyvalues[i+1],
		})
	}
	return attrs
}

// Element is the defacto helper function to construct a Node. You can also use the
// helper functions for every html element, e.g. A(), Div(), Input(), etc.
func Element(name string, attrs []Attribute, children ...Node) Node {
	return Node{
		Name:       name,
		Attributes: attrs,
		Children:   children,
	}
}

// Text is a helper function to construct a Node that has no tag no attributes and no children.
func Text(s string) Node {
	return Node{
		InnerText: s,
	}
}

// Attribute represents a HTML attribute.
//
// This struct is fully exported for the convenience of asserting values during tests
// but should not be used directly otherwise.
type Attribute struct {
	Name string

	// conceptually a union type `template.HTMLAttr | string`
	ValueHTML template.HTMLAttr
	ValueText string
}

// HTML returns the HTML representation of the attribute.
func (a Attribute) HTML() template.HTML {
	valueHTML := a.ValueHTML
	if valueHTML == "" {
		valueHTML = template.HTMLAttr(template.HTMLEscapeString(a.ValueText))
	}
	return template.HTML(template.HTMLEscapeString(a.Name) + "=\"" + string(valueHTML) + "\"")
}

// Node represents a HTML element.
//
// This struct is fully exported for the convenience of asserting values during tests
// but should not be used directly otherwise.
type Node struct {
	Name       string
	Attributes []Attribute

	// conceptually a union type `[]Node | template.HTML | string`
	Children  []Node
	InnerHTML template.HTML
	InnerText string
}

// HTML returns the HTML representation of the node.
func (e Node) HTML() template.HTML {
	if e.Name == "" {
		if e.InnerHTML != "" {
			return e.InnerHTML
		}
		return template.HTML(template.HTMLEscapeString(e.InnerText))
	}

	var attrsHTML strings.Builder
	for _, attr := range e.Attributes {
		attrsHTML.WriteString(string(" " + attr.HTML()))
	}

	var childrenHTML strings.Builder
	childrenHTML.WriteString(string(e.InnerHTML))
	if childrenHTML.Len() == 0 {
		childrenHTML.WriteString(template.HTMLEscapeString(e.InnerText))
	}
	if childrenHTML.Len() == 0 {
		for _, child := range e.Children {
			childrenHTML.WriteString(string(child.HTML()))
		}
	}
	tagName := template.HTMLEscapeString(e.Name)
	return template.HTML("<" + tagName + attrsHTML.String() + ">" + childrenHTML.String() + "</" + tagName + ">")
}

// Helper functions for every html element, using Element() and Text() helpers.

// A returns a Node with name "a".
func A(attrs []Attribute, children ...Node) Node {
	return Element("a", attrs, children...)
}

// Abbr returns a Node with name "abbr".
func Abbr(attrs []Attribute, children ...Node) Node {
	return Element("abbr", attrs, children...)
}

// Address returns a Node with name "address".
func Address(attrs []Attribute, children ...Node) Node {
	return Element("address", attrs, children...)
}

// Area returns a Node with name "area".
func Area(attrs []Attribute, children ...Node) Node {
	return Element("area", attrs, children...)
}

// Article returns a Node with name "article".
func Article(attrs []Attribute, children ...Node) Node {
	return Element("article", attrs, children...)
}

// Aside returns a Node with name "aside".
func Aside(attrs []Attribute, children ...Node) Node {
	return Element("aside", attrs, children...)
}

// Audio returns a Node with name "audio".
func Audio(attrs []Attribute, children ...Node) Node {
	return Element("audio", attrs, children...)
}

// B returns a Node with name "b".
func B(attrs []Attribute, children ...Node) Node {
	return Element("b", attrs, children...)
}

// Base returns a Node with name "base".
func Base(attrs []Attribute, children ...Node) Node {
	return Element("base", attrs, children...)
}

// Bdi returns a Node with name "bdi".
func Bdi(attrs []Attribute, children ...Node) Node {
	return Element("bdi", attrs, children...)
}

// Bdo returns a Node with name "bdo".
func Bdo(attrs []Attribute, children ...Node) Node {
	return Element("bdo", attrs, children...)
}

// Blockquote returns a Node with name "blockquote".
func Blockquote(attrs []Attribute, children ...Node) Node {
	return Element("blockquote", attrs, children...)
}

// Body returns a Node with name "body".
func Body(attrs []Attribute, children ...Node) Node {
	return Element("body", attrs, children...)
}

// Br returns a Node with name "br".
func Br(attrs []Attribute) Node {
	return Element("br", attrs)
}

// Button returns a Node with name "button".
func Button(attrs []Attribute, children ...Node) Node {
	return Element("button", attrs, children...)
}

// Canvas returns a Node with name "canvas".
func Canvas(attrs []Attribute, children ...Node) Node {
	return Element("canvas", attrs, children...)
}

// Caption returns a Node with name "caption".
func Caption(attrs []Attribute, children ...Node) Node {
	return Element("caption", attrs, children...)
}

// Cite returns a Node with name "cite".
func Cite(attrs []Attribute, children ...Node) Node {
	return Element("cite", attrs, children...)
}

// Code returns a Node with name "code".
func Code(attrs []Attribute, children ...Node) Node {
	return Element("code", attrs, children...)
}

// Col returns a Node with name "col".
func Col(attrs []Attribute, children ...Node) Node {
	return Element("col", attrs, children...)
}

// Colgroup returns a Node with name "colgroup".
func Colgroup(attrs []Attribute, children ...Node) Node {
	return Element("colgroup", attrs, children...)
}

// Data returns a Node with name "data".
func Data(attrs []Attribute, children ...Node) Node {
	return Element("data", attrs, children...)
}

// Datalist returns a Node with name "datalist".
func Datalist(attrs []Attribute, children ...Node) Node {
	return Element("datalist", attrs, children...)
}

// Dd returns a Node with name "dd".
func Dd(attrs []Attribute, children ...Node) Node {
	return Element("dd", attrs, children...)
}

// Del returns a Node with name "del".
func Del(attrs []Attribute, children ...Node) Node {
	return Element("del", attrs, children...)
}

// Details returns a Node with name "details".
func Details(attrs []Attribute, children ...Node) Node {
	return Element("details", attrs, children...)
}

// Dfn returns a Node with name "dfn".
func Dfn(attrs []Attribute, children ...Node) Node {
	return Element("dfn", attrs, children...)
}

// Dialog returns a Node with name "dialog".
func Dialog(attrs []Attribute, children ...Node) Node {
	return Element("dialog", attrs, children...)
}

// Div returns a Node with name "div".
func Div(attrs []Attribute, children ...Node) Node {
	return Element("div", attrs, children...)
}

// Dl returns a Node with name "dl".
func Dl(attrs []Attribute, children ...Node) Node {
	return Element("dl", attrs, children...)
}

// Dt returns a Node with name "dt".
func Dt(attrs []Attribute, children ...Node) Node {
	return Element("dt", attrs, children...)
}

// Em returns a Node with name "em".
func Em(attrs []Attribute, children ...Node) Node {
	return Element("em", attrs, children...)
}

// Embed returns a Node with name "embed".
func Embed(attrs []Attribute, children ...Node) Node {
	return Element("embed", attrs, children...)
}

// Fieldset returns a Node with name "fieldset".
func Fieldset(attrs []Attribute, children ...Node) Node {
	return Element("fieldset", attrs, children...)
}

// Figcaption returns a Node with name "figcaption".
func Figcaption(attrs []Attribute, children ...Node) Node {
	return Element("figcaption", attrs, children...)
}

// Figure returns a Node with name "figure".
func Figure(attrs []Attribute, children ...Node) Node {
	return Element("figure", attrs, children...)
}

// Footer returns a Node with name "footer".
func Footer(attrs []Attribute, children ...Node) Node {
	return Element("footer", attrs, children...)
}

// Form returns a Node with name "form".
func Form(attrs []Attribute, children ...Node) Node {
	return Element("form", attrs, children...)
}

// H1 returns a Node with name "h1".
func H1(attrs []Attribute, children ...Node) Node {
	return Element("h1", attrs, children...)
}

// H2 returns a Node with name "h2".
func H2(attrs []Attribute, children ...Node) Node {
	return Element("h2", attrs, children...)
}

// H3 returns a Node with name "h3".
func H3(attrs []Attribute, children ...Node) Node {
	return Element("h3", attrs, children...)
}

// H4 returns a Node with name "h4".
func H4(attrs []Attribute, children ...Node) Node {
	return Element("h4", attrs, children...)
}

// H5 returns a Node with name "h5".
func H5(attrs []Attribute, children ...Node) Node {
	return Element("h5", attrs, children...)
}

// H6 returns a Node with name "h6".
func H6(attrs []Attribute, children ...Node) Node {
	return Element("h6", attrs, children...)
}

// Head returns a Node with name "head".
func Head(attrs []Attribute, children ...Node) Node {
	return Element("head", attrs, children...)
}

// Header returns a Node with name "header".
func Header(attrs []Attribute, children ...Node) Node {
	return Element("header", attrs, children...)
}

// Hr returns a Node with name "hr".
func Hr(attrs []Attribute) Node {
	return Element("hr", attrs)
}

// Html returns a Node with name "html".
func Html(attrs []Attribute, children ...Node) Node {
	return Element("html", attrs, children...)
}

// I returns a Node with name "i".
func I(attrs []Attribute, children ...Node) Node {
	return Element("i", attrs, children...)
}

// Iframe returns a Node with name "iframe".
func Iframe(attrs []Attribute) Node {
	return Element("iframe", attrs)
}

// Img returns a Node with name "img".
func Img(attrs []Attribute) Node {
	return Element("img", attrs)
}

// Input returns a Node with name "input".
func Input(attrs []Attribute) Node {
	return Element("input", attrs)
}

// Ins returns a Node with name "ins".
func Ins(attrs []Attribute, children ...Node) Node {
	return Element("ins", attrs, children...)
}

// Kbd returns a Node with name "kbd".
func Kbd(attrs []Attribute, children ...Node) Node {
	return Element("kbd", attrs, children...)
}

// Label returns a Node with name "label".
func Label(attrs []Attribute, children ...Node) Node {
	return Element("label", attrs, children...)
}

// Legend returns a Node with name "legend".
func Legend(attrs []Attribute, children ...Node) Node {
	return Element("legend", attrs, children...)
}

// Li returns a Node with name "li".
func Li(attrs []Attribute, children ...Node) Node {
	return Element("li", attrs, children...)
}

// Link returns a Node with name "link".
func Link(attrs []Attribute) Node {
	return Element("link", attrs)
}

// Main returns a Node with name "main".
func Main(attrs []Attribute, children ...Node) Node {
	return Element("main", attrs, children...)
}

// Map returns a Node with name "map".
func Map(attrs []Attribute, children ...Node) Node {
	return Element("map", attrs, children...)
}

// Mark returns a Node with name "mark".
func Mark(attrs []Attribute, children ...Node) Node {
	return Element("mark", attrs, children...)
}

// Meta returns a Node with name "meta".
func Meta(attrs []Attribute) Node {
	return Element("meta", attrs)
}

// Meter returns a Node with name "meter".
func Meter(attrs []Attribute, children ...Node) Node {
	return Element("meter", attrs, children...)
}

// Nav returns a Node with name "nav".
func Nav(attrs []Attribute, children ...Node) Node {
	return Element("nav", attrs, children...)
}

// Noscript returns a Node with name "noscript".
func Noscript(attrs []Attribute, children ...Node) Node {
	return Element("noscript", attrs, children...)
}

// Object returns a Node with name "object".
func Object(attrs []Attribute, children ...Node) Node {
	return Element("object", attrs, children...)
}

// Ol returns a Node with name "ol".
func Ol(attrs []Attribute, children ...Node) Node {
	return Element("ol", attrs, children...)
}

// Optgroup returns a Node with name "optgroup".
func Optgroup(attrs []Attribute, children ...Node) Node {
	return Element("optgroup", attrs, children...)
}

// Option returns a Node with name "option".
func Option(attrs []Attribute, children ...Node) Node {
	return Element("option", attrs, children...)
}

// Output returns a Node with name "output".
func Output(attrs []Attribute, children ...Node) Node {
	return Element("output", attrs, children...)
}

// P returns a Node with name "p".
func P(attrs []Attribute, children ...Node) Node {
	return Element("p", attrs, children...)
}

// Param returns a Node with name "param".
func Param(attrs []Attribute, children ...Node) Node {
	return Element("param", attrs, children...)
}

// Picture returns a Node with name "picture".
func Picture(attrs []Attribute, children ...Node) Node {
	return Element("picture", attrs, children...)
}

// Pre returns a Node with name "pre".
func Pre(attrs []Attribute, children ...Node) Node {
	return Element("pre", attrs, children...)
}

// Progress returns a Node with name "progress".
func Progress(attrs []Attribute, children ...Node) Node {
	return Element("progress", attrs, children...)
}

// Q returns a Node with name "q".
func Q(attrs []Attribute, children ...Node) Node {
	return Element("q", attrs, children...)
}

// Rp returns a Node with name "rp".
func Rp(attrs []Attribute, children ...Node) Node {
	return Element("rp", attrs, children...)
}

// Rt returns a Node with name "rt".
func Rt(attrs []Attribute, children ...Node) Node {
	return Element("rt", attrs, children...)
}

// Ruby returns a Node with name "ruby".
func Ruby(attrs []Attribute, children ...Node) Node {
	return Element("ruby", attrs, children...)
}

// S returns a Node with name "s".
func S(attrs []Attribute, children ...Node) Node {
	return Element("s", attrs, children...)
}

// Samp returns a Node with name "samp".
func Samp(attrs []Attribute, children ...Node) Node {
	return Element("samp", attrs, children...)
}

// Script returns a Node with name "script".
func Script(attrs []Attribute, children ...Node) Node {
	return Element("script", attrs, children...)
}

// Section returns a Node with name "section".
func Section(attrs []Attribute, children ...Node) Node {
	return Element("section", attrs, children...)
}

// Select returns a Node with name "select".
func Select(attrs []Attribute, children ...Node) Node {
	return Element("select", attrs, children...)
}

// Small returns a Node with name "small".
func Small(attrs []Attribute, children ...Node) Node {
	return Element("small", attrs, children...)
}

// Source returns a Node with name "source".
func Source(attrs []Attribute) Node {
	return Element("source", attrs)
}

// Span returns a Node with name "span".
func Span(attrs []Attribute, children ...Node) Node {
	return Element("span", attrs, children...)
}

// Strong returns a Node with name "strong".
func Strong(attrs []Attribute, children ...Node) Node {
	return Element("strong", attrs, children...)
}

// Style returns a Node with name "style".
func Style(attrs []Attribute, children ...Node) Node {
	return Element("style", attrs, children...)
}

// Sub returns a Node with name "sub".
func Sub(attrs []Attribute, children ...Node) Node {
	return Element("sub", attrs, children...)
}

// Summary returns a Node with name "summary".
func Summary(attrs []Attribute, children ...Node) Node {
	return Element("summary", attrs, children...)
}

// Sup returns a Node with name "sup".
func Sup(attrs []Attribute, children ...Node) Node {
	return Element("sup", attrs, children...)
}

// Table returns a Node with name "table".
func Table(attrs []Attribute, children ...Node) Node {
	return Element("table", attrs, children...)
}

// Tbody returns a Node with name "tbody".
func Tbody(attrs []Attribute, children ...Node) Node {
	return Element("tbody", attrs, children...)
}

// Td returns a Node with name "td".
func Td(attrs []Attribute, children ...Node) Node {
	return Element("td", attrs, children...)
}

// Template returns a Node with name "template".
func Template(attrs []Attribute, children ...Node) Node {
	return Element("template", attrs, children...)
}

// Textarea returns a Node with name "textarea".
func Textarea(attrs []Attribute, children ...Node) Node {
	return Element("textarea", attrs, children...)
}

// Tfoot returns a Node with name "tfoot".
func Tfoot(attrs []Attribute, children ...Node) Node {
	return Element("tfoot", attrs, children...)
}

// Th returns a Node with name "th".
func Th(attrs []Attribute, children ...Node) Node {
	return Element("th", attrs, children...)
}

// Thead returns a Node with name "thead".
func Thead(attrs []Attribute, children ...Node) Node {
	return Element("thead", attrs, children...)
}

// Time returns a Node with name "time".
func Time(attrs []Attribute, children ...Node) Node {
	return Element("time", attrs, children...)
}

// Title returns a Node with name "title".
func Title(attrs []Attribute, children ...Node) Node {
	return Element("title", attrs, children...)
}

// Tr returns a Node with name "tr".
func Tr(attrs []Attribute, children ...Node) Node {
	return Element("tr", attrs, children...)
}

// Track returns a Node with name "track".
func Track(attrs []Attribute) Node {
	return Element("track", attrs)
}

// U returns a Node with name "u".
func U(attrs []Attribute, children ...Node) Node {
	return Element("u", attrs, children...)
}

// Ul returns a Node with name "ul".
func Ul(attrs []Attribute, children ...Node) Node {
	return Element("ul", attrs, children...)
}

// Var returns a Node with name "var".
func Var(attrs []Attribute, children ...Node) Node {
	return Element("var", attrs, children...)
}

// Video returns a Node with name "video".
func Video(attrs []Attribute, children ...Node) Node {
	return Element("video", attrs, children...)
}

// Wbr returns a Node with name "wbr".
func Wbr(attrs []Attribute) Node {
	return Element("wbr", attrs)
}

// SVG helper functions

// SVG returns a Node with name "svg".
func SVG(attrs []Attribute, children ...Node) Node {
	return Element("svg", attrs, children...)
}

// Circle returns a Node with name "circle".
func Circle(attrs []Attribute, children ...Node) Node {
	return Element("circle", attrs, children...)
}

// Ellipse returns a Node with name "ellipse".
func Ellipse(attrs []Attribute, children ...Node) Node {
	return Element("ellipse", attrs, children...)
}

// Line returns a Node with name "line".
func Line(attrs []Attribute, children ...Node) Node {
	return Element("line", attrs, children...)
}

// Path returns a Node with name "path".
func Path(attrs []Attribute, children ...Node) Node {
	return Element("path", attrs, children...)
}

// Polygon returns a Node with name "polygon".
func Polygon(attrs []Attribute, children ...Node) Node {
	return Element("polygon", attrs, children...)
}

// Polyline returns a Node with name "polyline".
func Polyline(attrs []Attribute, children ...Node) Node {
	return Element("polyline", attrs, children...)
}

// Rect returns a Node with name "rect".
func Rect(attrs []Attribute, children ...Node) Node {
	return Element("rect", attrs, children...)
}

// TextSVG returns a Node with name "text".
func TextSVG(attrs []Attribute, children ...Node) Node {
	return Element("text", attrs, children...)
}

// Tspan returns a Node with name "tspan".
func Tspan(attrs []Attribute, children ...Node) Node {
	return Element("tspan", attrs, children...)
}

// Use returns a Node with name "use".
func Use(attrs []Attribute, children ...Node) Node {
	return Element("use", attrs, children...)
}

// View returns a Node with name "view".
func View(attrs []Attribute, children ...Node) Node {
	return Element("view", attrs, children...)
}

// ForeignObject returns a Node with name "foreignObject".
func ForeignObject(attrs []Attribute, children ...Node) Node {
	return Element("foreignObject", attrs, children...)
}
