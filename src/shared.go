package main

// Arguments
var (
	listenAddr string
	workerUrl  string
	isLog      bool
	confPath   string
)

// Key: Value = Source Host: Replaced Host
var (
	convHostMap = map[string]string{
		"raw.fastgit.org":      "raw.githubusercontent.com",
		"download.fastgit.org": "github.com",
		"hub.fastgit.org":      "github.com",
		"archive.fastgit.org":  "github.com",
		"codeload.fastgit.org": "codeload.github.com",
	}
)

const (
	robotsTxt = "User-Agent: *\nDisallow: /"
)

var robotsTxtByte = []byte(robotsTxt)
