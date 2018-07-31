package main1

import (
	g "./goroutines"
	"net/http"
)

func iHandler(w http.ResponseWriter, r *http.Request) {

}

func main1() {
	http.HandleFunc("/", iHandler)
	http.ListenAndServe(":80", nil)
}
