package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" {
		w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="refresh" content="3; url=https://doc.fastgit.org">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>FinalEndPoint</title>
</head>
<body>
<p>OK by FinalEndPoint</p>
<p>
    The browser will be redirected to the document page, please wait 3 seconds.
</p>
</body>
</html>
`))
		return
	}

	if r.Method != "GET" {
		http.Error(w, "503 Bad Request - Not Allowed Method", http.StatusServiceUnavailable)
		return
	}

	if strings.HasPrefix(r.RequestURI, "/w2766/") {
		http.Error(w, "402 - You NEED pay!", http.StatusPaymentRequired)
		return
	}

	if r.RequestURI == "/robots.txt" {
		w.Write(robotsTxtByte)
		return
	}

	if isLog {
		log.Println(r.Host + r.RequestURI)
	}

	url, ok := convFastGitToCloudFlareWorker(r.Host, r.RequestURI)
	if ok {
		http.Redirect(w, r, url, 301)
	} else {
		http.Error(w, "503 Bad Request - Not Allowed Url", http.StatusServiceUnavailable)
	}
	return
}

func main() {
	initialise()

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at %v...\n", listenAddr)
	fmt.Printf("Workers Link: %v\n", workerUrl)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}

func initialise() {
	parseArgs()

	if confPath != "" {
		m, ok := parseFromFile(confPath)
		if ok {
			convHostMap = m
		}
	}
}
