package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/choonkeat/dom-go"
)

func body(s string) dom.Node {
	return dom.Element("ul",
		dom.Attrs(),
		dom.Element("li",
			dom.Attrs(),
			dom.Text("dom-go is a Go library for generating HTML."),
		),
		dom.Element("li",
			dom.Attrs(),
			dom.Text("?param is "),
			dom.Text(s),
		),
		dom.Element("li",
			dom.Attrs(),
			dom.A(dom.Attrs("href", "/dom-go?param=123"),
				dom.Strong(dom.Attrs(), dom.Text("dom-go")),
				dom.Text(" with dom-go"),
			),
		),
		dom.Element("li",
			dom.Attrs(),
			dom.A(dom.Attrs("href", "/tmpl?param=456"),
				dom.Strong(dom.Attrs(), dom.Text("dom-go")),
				dom.Text(" with html/template"),
			),
		),
	)
}

func main() {
	http.HandleFunc("/tmpl", func(w http.ResponseWriter, r *http.Request) {
		templateName := "index.html"
		tmpl, err := template.New(templateName).Parse(`
			<!DOCTYPE html>
			<html>
				<head>
					<title>dom-go</title>
				</head>
				<body>
					<h1>dom-go with html/template</h1>
					{{ . }}
				</body>
			</html>
			`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, templateName, body(r.URL.Query().Get("param")).HTML())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(
			dom.Html(
				dom.Attrs(),
				dom.Head(
					dom.Attrs(),
					dom.Title(
						dom.Attrs(),
						dom.Text("Go Web"),
					),
				),
				dom.Body(
					dom.Attrs(),
					dom.H1(
						dom.Attrs(),
						dom.Text("dom-go with dom-go"),
					),
					body(r.URL.Query().Get("param")),
				),
			).HTML(),
		))
	})
	log.Println("Listening on :8080...")
	log.Println(http.ListenAndServe(":8080", nil))
}
