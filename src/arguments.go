package main

import (
	"flag"
	"strings"
)

func parseArgs() {
	flag.StringVar(&listenAddr, "addr", "0.0.0.0:5000", "The address which needs to be bound. Default is 0.0.0.0:5000")
	flag.StringVar(&workerUrl, "worker", "https://endpoint.fastgit.org", "The final endpoint place. Default is https://endpoint.fastgit.org")
	flag.BoolVar(&isLog, "log", true, "Enable log")
	flag.StringVar(&confPath, "conv-path", "", "Convert map storage place")

	flag.Parse()

	if !strings.HasSuffix(workerUrl, "/") {
		workerUrl += "/"
	}
}
