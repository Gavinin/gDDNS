package dns

import "gDDNS/internal/local"

type Record struct {
	IP         local.IP
	Domain     string
	RecordType string
}
