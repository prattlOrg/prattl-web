package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageVariables struct {
	Path   string
	Params string
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.String(), "/")
		var filteredParts []string

		for _, part := range parts {
			if strings.TrimSpace(part) != "" {
				filteredParts = append(filteredParts, strings.TrimSpace(part))
			}
		}
		fmt.Println(parts)

		// payment routes do not need to come from htmx elements
		if len(filteredParts) > 0 && filteredParts[0] == "payment" {
		} else {
			hxReq := r.Header.Get("Hx-Request")

			if hxReq == "" {
				urlParts := strings.Split(r.URL.String(), "?")
				path := urlParts[0]
				params := ""

				if path == "/" {
					path = "/landing"
				}

				if len(urlParts) > 1 {
					params = string('?') + urlParts[1]
				}

				pageVariables := PageVariables{
					Path:   path,
					Params: params,
				}

				tmpl := template.Must(template.ParseFiles(IndexTemplates...))

				err := tmpl.Execute(w, pageVariables)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

		}

		h.ServeHTTP(w, r)
	})
}
