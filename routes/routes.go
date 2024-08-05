package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type PageData struct {
	Markdown string
}

func InitializePageRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.Handle("/", middleware(mux))
	return mux
}

var IndexTemplates = []string{
	"public/html/index.html",
	"public/html/layout.html",
}
var ErrTemplates = []string{
	"public/html/index.html",
	"public/html/404.html",
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl := template.Must(template.ParseFiles(ErrTemplates...))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	md, err := convertMarkdown()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl := template.Must(template.ParseFiles(IndexTemplates...))
	tmpl.New("markdown")
	template.Must(tmpl.Lookup("markdown").Parse(md))
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func convertMarkdown() (string, error) {
	source, err := os.ReadFile("public/static/assets/instructions.md")
	if err != nil {
		return "", err
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return "", err
	}
	result := buf.String()
	return result, nil
}
