package legox

import "github.com/go-acme/lego/v4/challenge/dns01"

var (
	GoogleDNS     = "8.8.8.8"
	DefaultNSOpts = SetNSOpts(GoogleDNS)
)

func SetNSOpts(ns ...string) dns01.ChallengeOption {
	return dns01.AddRecursiveNameservers(ns)
}
