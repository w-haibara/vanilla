package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func echoAPIHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, r.Body)
}

func main() {
	/*
	 * Title View (genelated by http://patorjk.com/software/taag/#p=display&h=0&v=0&f=Epic&t=vanilla%0A)
	 */
	fmt.Println(`
|\     /|(  ___  )( (    /|\__   __/( \      ( \      (  ___  )
| )   ( || (   ) ||  \  ( |   ) (   | (      | (      | (   ) |
| |   | || (___) ||   \ | |   | |   | |      | |      | (___) |
( (   ) )|  ___  || (\ \) |   | |   | |      | |      |  ___  |
 \ \_/ / | (   ) || | \   |   | |   | |      | |      | (   ) |
  \   /  | )   ( || )  \  |___) (___| (____/\| (____/\| )   ( |
   \_/   |/     \||/    )_)\_______/(_______/(_______/|/     \|`)

	http.HandleFunc("/enc/echo", crypto(echoAPIHandler))

	const appDir = "./page"
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(appDir))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
