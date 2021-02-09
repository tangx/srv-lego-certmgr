package certgen

import "strings"

func providerPostion(uri string, pos int) string {
	uriList := strings.Split(uri, "/")
	return uriList[pos]
}

func sortDomains(domains string) string {
	dl := splitDomains(domains)
	return joinDomains(dl)
}

func splitDomains(domains string) []string {
	return strings.Split(domains, ",")
}

func joinDomains(domains []string) string {
	return strings.Join(domains, ",")
}
