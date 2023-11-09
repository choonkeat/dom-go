# DOM

Construct HTML elements in Go. Can be used together with `html/template` templates.

# Example

```go
elem := dom.Element("div",
    dom.Attrs(
        "class", "1 2 3",
        "data-foo", `4<'"5"'>6`,
    ),
    dom.InnerText("<oops>789</oops>"),
    dom.P(
        dom.Attrs(),
        dom.InnerHTML("<strong>10</strong>"),
    ),
)
```

representing HTML

```html
<div class="1 2 3" data-foo="4&lt;&#39;&#34;5&#34;&#39;&gt;6">
    &lt;oops&gt;789&lt;/oops&gt;
    <p><strong>10</strong></p>
</div>
```

Indented for illustrative purpose; there are no newlines introduced.

Notice the text values added via `InnerText` are html safe and `InnerHTML` trusts your raw html

## Usage (Standalone)

```go
fmt.Println(elem.HTML())
```

or in a `http.HandlerFunc`

```go
w.Write([]byte(elem.HTML()))
```

## Usage (html/template)

```go
tmpl.ExecuteTemplate(w, "index.html", elem.HTML())
```

with `index.html` content being

```html
<!DOCTYPE html>
<html>
    <head><title>dom-go</title></head>
    <body>{{ . }}</body>
</html>
```
