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
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
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
