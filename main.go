package main

import (
	"fmt"
	"log"
	"net/http"
)

func TestServer(w http.ResponseWriter, r *http.Request) {
	response := "Failed"
	sc := http.StatusBadRequest
	if r.URL.RawPath != "" {
		sc = http.StatusOK
		response = "Success"
	}
	w.WriteHeader(sc)
	fmt.Fprint(w, response)
}
func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(TestServer)); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
