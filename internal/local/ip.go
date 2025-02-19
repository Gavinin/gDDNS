package local

type IP struct {
	IPv4 string
	IPv6 string
}

type Server interface {
	Query() *IP
}
