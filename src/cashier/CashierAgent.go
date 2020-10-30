package cashier

import (
	"math"
	"math/rand"
	"time"
)

type CashierAgent struct {
	cleaningTime     float64
	diligenceFactor  float64
	expedienceFactor float64
	baseTimeCash     float64
	baseTimeCard     float64
	baseTimeItem     float64

	ActualTimeCash float64
	ActualTimeCard float64
	ActualTimeItem float64
}

func NewCashier() CashierAgent {
	rand.Seed(time.Now().UTC().UnixNano())
	cashier := CashierAgent{}

	cashier.cleaningTime = math.Round(((rand.Float64()*0.5)+0.25)*100) / 100
	cashier.diligenceFactor = math.Round(((rand.Float64()*0.5)+0.25)*100) / 100
	cashier.expedienceFactor = math.Round(((rand.Float64()*0.5)+0.25)*100) / 100
	cashier.baseTimeCash = math.Round(((rand.Float64()*0.5)+0.25)*100) / 100
	cashier.baseTimeCard = math.Round(((rand.Float64()*0.5)+0.25)*100) / 100
	cashier.baseTimeItem = math.Round(((rand.Float64()*0.5)+0.25)*100) / 100

	cashier.ActualTimeCash = cashier.baseTimeCash
	cashier.ActualTimeCard = cashier.baseTimeCard
	cashier.ActualTimeItem = cashier.baseTimeItem

	return cashier
}
