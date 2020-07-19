package main

import (
	"io"
	"log"
	"net/http"

	"github.com/w-haibara/vanilla"
)

func echoAPIHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, r.Body)
}

func main() {
	http.HandleFunc("/enc/echo", CryptoHandler(echoAPIHandler))

	const appDir = "./page"
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(appDir))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
