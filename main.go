package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/klog"
)

type logData map[string]interface{}

func testServer(w http.ResponseWriter, r *http.Request) {
	response := "Failed"
	sc := http.StatusBadRequest
	if r.URL.EscapedPath() != "" {
		sc = http.StatusOK
		response = "Success"
	}
	w.WriteHeader(sc)
	fmt.Fprint(w, response)
}

func debugMiddleware(w http.ResponseWriter, r *http.Request) {
	dataJSON, err := json.Marshal(logData{
		"URLPath":              r.URL.Path,
		"Request Body":         r.Body,
		"Request Query Params": r.URL.Query(),
	})
	if err != nil {
		klog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	klog.Info("http-middleware ", string(dataJSON))

	testServer(w, r)
}

func main() {
	klog.InitFlags(nil)
	if err := http.ListenAndServe(":8080", http.StripPrefix("/", http.HandlerFunc(debugMiddleware))); err != nil {
		klog.Fatalf("could not listen on port 8080 %v", err)
	}
}
