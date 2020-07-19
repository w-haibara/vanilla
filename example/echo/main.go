package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/w-haibara/vanilla"
)

func echoAPIHandler(w http.ResponseWriter, r *http.Request) {
	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	fmt.Println("recv msg:", body.String())

	w.Write(body.Bytes())
}

func main() {
	fmt.Println("--- echo server ---")

	http.HandleFunc("/enc/echo", vanilla.CryptoHandler(echoAPIHandler))

	const appDir = "./page"
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(appDir))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
