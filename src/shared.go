package main

const (
	listenAddr = "0.0.0.0:5000"
	workerUrl  = ""
)

var (
	convHostMap = map[string]string{
		"raw.fastgit.org":      "raw.githubusercontent.com",
		"download.fastgit.org": "github.com",
		"hub.fastgit.org":      "github.com",
	}
)
