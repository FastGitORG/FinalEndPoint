package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.RequestURI == "/" {
			w.Write([]byte("OK by FinalEndPoint"))
			return
		}
		url, ok := convFastGitToCloudFlareWorker(r.Host, r.RequestURI)
		if ok {
			http.Redirect(w, r, url, 301)
		} else {
			http.Error(w, "503 Bad Request - Not Allowed Url", http.StatusServiceUnavailable)
		}
		return
	default:
		http.Error(w, "503 Bad Request - Not Allowed Method", http.StatusServiceUnavailable)
		return
	}
}

func main() {
	parseArgs()

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at %v...\n", listenAddr)
	fmt.Printf("Workers Link: %v\n", workerUrl)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
