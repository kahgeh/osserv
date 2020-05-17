package main

import (
	"fmt"
	"log"
	"net/http"

	"kahgeh.com/infra/osserv/pkg/fs"
)

func getVersion(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "1.0.0.")
}

func getMountInfo(w http.ResponseWriter, _ *http.Request) {
	isMounted, _ := fs.IsMounted("/dev/sdf")
	fmt.Fprintf(w, "{\"isMounted\": %t}", isMounted)
}

func main() {
	http.HandleFunc("/version", getVersion)
	http.HandleFunc("/api/fs/_dev_sdf", getMountInfo)
	log.Fatal(http.ListenAndServe(":7000", nil))
}
