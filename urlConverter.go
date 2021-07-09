package main

import "net/url"


func convFastGitToCloudFlareWorker(host string, uri string) string {
	newHost, ok := convHostMap[host]
	if !ok {
		newHost = host
	}
	return workerUrl + "https://" + newHost + "/" + url.QueryEscape(uri)
}
