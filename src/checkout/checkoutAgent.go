package checkout

import (
	"math"
	"math/rand"
	"src/cashier"
	"time"
)

type CheckoutAgent struct {
	SelfCheckout bool
	AdultCheckout bool 
	AssistanceWaitTime float64
	TotalMoney float64
	FirstShiftCashier cashier.CashierAgent
	SecondShiftCashier cashier.CashierAgent
}

var t = rand.New(rand.NewSource(time.Now().UnixNano()))

// checkout agent constructor 
func CreateInitialisedCheckoutAgent() CheckoutAgent {
	co := CheckoutAgent{}

	// Randomly Initialised Variables
	co.SelfCheckout = false
	co.AdultCheckout = (t.Intn(2) == 1)
	co.AssistanceWaitTime = math.Round(((t.Float64()*(0.75-0.25))+0.25)*100) / 100
	co.TotalMoney = 0
	co.FirstShiftCashier = cashier.CashierAgent{}
	co.SecondShiftCashier = cashier.CashierAgent{}
	
	return co
}

func (co *CheckoutAgent) PropagateTime() {
	// I presume item handling and such will be handled mostly by the cashier and here is only for record keeping and spillage chances?
	// Placeholder stuff, 1/1000 chance of spillage and do nothing on that occurence 
	if t.Float64() < 0.001 {

	} else {
		co.TotalMoney += t.Float64()
	}
}
