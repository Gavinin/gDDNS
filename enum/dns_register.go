package enum

import "errors"

type DomainRegister string

var UnSupportErr = errors.New("unsupported method")

const (
	SpaceShip DomainRegister = "SpaceShip"
	NameCheap DomainRegister = "NameCheap"
)
