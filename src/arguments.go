package main

import "flag"

func parseArgs() {
	flag.StringVar(&listenAddr, "addr", "0.0.0.0:5000", "The address which needs to be bound. Default is 0.0.0.0:5000")
	flag.StringVar(&workerUrl, "worker", "", "The final endpoint place. Default is empty")
	flag.Parse()
}
