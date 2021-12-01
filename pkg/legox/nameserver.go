package legox

import "github.com/go-acme/lego/v4/challenge/dns01"

const (
	GoogleDNS = "8.8.8.8"
	AliyunDNS = "223.5.5.5"
)

var (
	DefaultNSOpts = SetNSOpts(AliyunDNS, GoogleDNS)
)

func SetNSOpts(ns ...string) dns01.ChallengeOption {
	return dns01.AddRecursiveNameservers(ns)
}
