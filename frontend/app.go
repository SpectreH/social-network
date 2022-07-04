package main

import (
	"log"
	"net/http"
)

var PORT string = ":3000"

func main() {
	srv := &http.Server{
		Addr:    PORT,
		Handler: router(),
	}

	log.Printf("Listening on %s...", PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	httpFS := http.FileServer(http.Dir("dist"))
	mux.Handle("/static/", httpFS)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "dist/favicon.ico")
		return
	}

	http.ServeFile(w, r, "dist/index.html")
}
