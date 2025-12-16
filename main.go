package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))

	http.Handle("/", http.FileServer(http.Dir("templates")))

	addr := ":8080"
	log.Printf("Serveur démarré — http://localhost%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
