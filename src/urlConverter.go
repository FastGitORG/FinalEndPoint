package main

import (
	"strings"
)

func convFastGitToCloudFlareWorker(host string, uri string) (string, bool) {
	newHost, ok := convHostMap[removePort(host)]
	if !ok {
		newHost = host
	}
	return workerUrl + "https://" + newHost + uri, ok
}

func removePort(host string) string {
	return strings.Split(host, ":")[0]
}
