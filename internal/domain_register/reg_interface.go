package domain_register

import (
	"gDDNS/internal/dns"
)

type Service interface {
	ListDomains() ([]dns.Record, error)
	Update(record dns.Record) error
	PUT(record dns.Record) error
	Delete(record dns.Record) error
}
