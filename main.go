package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatalln("Listening port should be provided")
	}

	port := os.Args[1]

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			fmt.Sprintf(
				"Hello world from app in port %s", port,
			),
		))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
