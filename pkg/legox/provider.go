package legox

type Provider interface {
	ApplyCertificate(domains ...string) (Certificate, error)
}
