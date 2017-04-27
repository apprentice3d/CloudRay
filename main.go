package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Started server on port 3009 ...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up!"))
	})

	log.Fatal(http.ListenAndServe(":3009", nil))

}
