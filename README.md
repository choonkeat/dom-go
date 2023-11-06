# DOM

Construct HTML elements in Go. Can be used together with `html/template` templates.

# Example

```go
elem := dom.Element("div",
    dom.Attrs("class", "1 2 3", "data-foo", `4<'"5"'>6`),
    dom.Text("<oops>789</oops>"),
    dom.Strong(
        dom.Attrs(),
        dom.Text("10"),
    ),
)
```

representing HTML

```html
<div class="1 2 3" data-foo="4&lt;&#39;&#34;5&#34;&#39;&gt;6">
  &lt;oops&gt;789&lt;/oops&gt;
  <strong>10</strong>
</div>
```

notice the text values are html safe

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
