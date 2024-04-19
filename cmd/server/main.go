package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	portNumber := flag.Int("port", 8080, "Port number to serve on.")
	flag.Parse()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	fileserver := http.FileServer(http.Dir("assets"))
	router.Handle("/*", fileserver)

	fmt.Printf("SERVING ON http://localhost:%d\n\n", *portNumber)
	http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), router)
}
