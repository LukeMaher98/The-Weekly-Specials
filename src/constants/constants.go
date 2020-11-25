package constants

import (
	"src/customer"
	"sync"
)

type StaffAttributeBounds struct {
	AmicabilityUpperBound float64
	AmicabilityLowerBound float64
	CompetanceUpperBound  float64
	CompetanceLowerBound  float64
}

type StoreAttributeBoundsInt struct {
	UpperBound int
	LowerBound int
}

type StoreAttributeBoundsFloat struct {
	UpperBound float64
	LowerBound float64
}

type StoreShifts struct {
	FirstShiftCount  int
	SecondShiftCount int
}

type CustomerQueue struct {
	Mutex sync.Mutex
	Queue []customer.CustomerAgent
}
