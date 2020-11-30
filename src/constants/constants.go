package constants

import (
	"src/customer"
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

/*
type CustomerQueue struct {
	Mutex sync.Mutex
	Queue []customer.CustomerAgent
}
*/

type CustomerLost struct {
	Customer customer.CustomerAgent
	Day      int
	Time     float64
	Reason   string
}
