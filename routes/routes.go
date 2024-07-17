package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("in handler")
	tmpl := template.Must(template.ParseFiles(IndexTemplates...))

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/html/pages/404.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
