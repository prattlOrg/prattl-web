package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"prattl_web/routes"
)

func main() {
	mux := routes.InitializePageRoutes()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public/static"))))
	port := envPortOr("3000")
	fmt.Println("Listening on port ", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
