package main

import "strings"

func replaceFastGitToGitHub(url string) string {
	for k, v := range convHostMap {
		if strings.HasPrefix(url, k) {
			return strings.Replace(url, k, v, 1)
		}
	}
	return url
}
