package main

const (
	listenAddr = "0.0.0.0:5000"
	workerUrl  = ""
)

var (
	convLinkMap = map[string]string{
		"https://raw.fastgit.org":      "https://raw.githubusercontent.com",
		"https://download.fastgit.org": "https://github.com",
		"https://hub.fastgit.org":      "https://github.com",
	}

	allowedHosts = map[string]bool{
		"raw.fastgit.org":      true,
		"download.fastgit.org": true,
		"hub.fastgit.org":      true,
	}
)
