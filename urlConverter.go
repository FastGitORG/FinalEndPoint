package main

import "net/url"

func convFastGitToCloudFlareWorker(fgUrl string) string {
	fgUrl = replaceFastGitToGitHub(fgUrl)
	return workerUrl + url.QueryEscape(fgUrl)
}
