package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if !allowedHosts[r.Host] {
			http.Error(w, "503", http.StatusServiceUnavailable)
		}
		http.Redirect(w, r, convFastGitToCloudFlareWorker("https://"+r.Host+"/"+r.RequestURI), 301)

	default:
		http.Error(w, "503", http.StatusServiceUnavailable)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at %v...\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
