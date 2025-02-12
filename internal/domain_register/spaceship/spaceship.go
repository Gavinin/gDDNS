package spaceship

import (
	"gDDNS/internal/dns"
)

type SpaceShip struct {
	Key            string
	Secret         string
	TopLevelDomain string
}

func NewSpaceShip() *SpaceShip {
	return &SpaceShip{}
}

func (s SpaceShip) ListDomains() ([]dns.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (s SpaceShip) Update(record dns.Record) error {
	//TODO implement me
	panic("implement me")
}

func (s SpaceShip) Delete(record dns.Record) error {
	//TODO implement me
	panic("implement me")
}
