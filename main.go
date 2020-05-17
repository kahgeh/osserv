package main

import (
	"fmt"
	"log"
	"net/http"
)

func getVersion(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "1.0.0.")
}

func main() {
	http.HandleFunc("/version", getVersion)
	log.Fatal(http.ListenAndServe(":7000", nil))
}
