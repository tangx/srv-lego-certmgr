package certprovider

import "strings"

func ProviderPostion(uri string, pos int) string {
	uriList := strings.Split(uri, "/")
	return uriList[pos]
}

func SortDomains(domains string) string {
	dl := SplitDomains(domains)
	return JoinDomains(dl)
}

func SplitDomains(domains string) []string {
	return strings.Split(domains, ",")
}

func JoinDomains(domains []string) string {
	return strings.Join(domains, ",")
}
