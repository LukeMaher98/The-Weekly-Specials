package checkout

import (
	"math"
	"math/rand"
	"src/cashier"
	"src/constants"
	"src/customer"
	"time"
)

type CheckoutAgent struct {
	SelfCheckout       bool
	AdultCheckout      bool
	ProcessingCustomer bool
	// These two do nothing right now
	CurrentCustomerProgress float64
	AssistanceWaitTime      float64

	TotalMoney         float64
	FirstShiftCashier  cashier.CashierAgent
	SecondShiftCashier cashier.CashierAgent
	CurrentCustomer    customer.CustomerAgent

	CustomersProcessed int
}

// checkout agent constructor
func CreateInitialisedCheckoutAgent() CheckoutAgent {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	co := CheckoutAgent{}

	// Randomly Initialised Variables
	co.SelfCheckout = false
	co.AdultCheckout = (r.Intn(2) == 1)
	co.ProcessingCustomer = false
	co.CurrentCustomerProgress = 0
	co.AssistanceWaitTime = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	co.TotalMoney = 0
	co.FirstShiftCashier = cashier.CashierAgent{}
	co.SecondShiftCashier = cashier.CashierAgent{}

	return co
}

func (co *CheckoutAgent) PropagateTime() {

}

func (co *CheckoutAgent) IsManned(currentShift int) bool {
	Manned := false

	if currentShift == 0 {
		if ((cashier.CashierAgent{}) == co.FirstShiftCashier) {
			Manned = false
		} else {
			Manned = true
		}
	} else {
		if ((cashier.CashierAgent{}) == co.SecondShiftCashier) {
			Manned = false
		} else {
			Manned = true
		}
	}

	return Manned
}

func (co *CheckoutAgent) ProcessCustomer(ItemTimeBounds constants.StoreAttributeBoundsFloat) {

	for _, item := range co.CurrentCustomer.GetCustomerItems() {
		co.CurrentCustomerProgress += item.GetItemHandling()
		co.TotalMoney += item.GetPrice()
	}

	// 1 clock = 60 seconds, dividing by 10 for a [0.5-6] second time to scan per item depending on the handling
	co.CurrentCustomerProgress /= 10

	// Wait until current time is current time + Round(currentCustomerProgress)

	co.CustomersProcessed++
	co.ProcessingCustomer = false

}
