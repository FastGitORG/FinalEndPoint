package main

import "strings"

func replaceHttpToHttps(url string) string {
	if !strings.HasPrefix(url, "http://") {
		return url
	} else {
		return strings.Replace(url, "http://", "https://", 1)
	}
}

func replaceFastGitToGitHub(url string) string {
	url = replaceHttpToHttps(url)

	for k, v := range convLinkMap {
		if strings.HasPrefix(url, k) {
			return strings.Replace(url, k, v, 1)
		}
	}
	return url
}
