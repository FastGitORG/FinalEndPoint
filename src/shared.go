package main

// Arguments
var (
	listenAddr string
	workerUrl  string
)

// Key: Value = Source Host: Replaced Host
var (
	convHostMap = map[string]string{
		"raw.fastgit.org":      "raw.githubusercontent.com",
		"download.fastgit.org": "github.com",
		"hub.fastgit.org":      "github.com",
	}
)
