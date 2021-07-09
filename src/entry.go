package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		url, ok := convFastGitToCloudFlareWorker(r.Host, r.RequestURI)
		if ok {
			http.Redirect(w, r, url, 301)
		} else {
			http.Error(w, "503", http.StatusServiceUnavailable)
		}
		return
	default:
		http.Error(w, "503", http.StatusServiceUnavailable)
		return
	}
}

func main() {
	parseArgs()

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at %v...\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
