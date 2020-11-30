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

func (co *CheckoutAgent) IsManned(currentShift int) bool {
	Manned := false

	if currentShift == 0 {
		if (cashier.CashierAgent{}) == co.FirstShiftCashier {
			Manned = false
		} else {
			Manned = true
		}
	} else {
		if (cashier.CashierAgent{}) == co.SecondShiftCashier {
			Manned = false
		} else {
			Manned = true
		}
	}

	return Manned
}

func (co *CheckoutAgent) ProcessCustomer(ItemTimeBounds constants.StoreAttributeBoundsFloat, shift int) {

	// Only process actual customers
	if co.CurrentCustomer.GetInitialised() == true {
		for co.ProcessingCustomer == true {
			customerTotal := 0.0
			for _, item := range co.CurrentCustomer.GetCustomerItems() {
				//Age rated check goes here lenny : ) ?
				co.CurrentCustomerProgress += item.GetItemHandling()
				customerTotal += item.GetPrice()
			}
			co.TotalMoney += customerTotal

			sleepTime := time.Millisecond
			if shift == 0 {
				sleepTime = time.Duration(int(co.CurrentCustomerProgress * co.FirstShiftCashier.TimeToProcess()))
			} else if shift == 1 {
				sleepTime = time.Duration(int(co.CurrentCustomerProgress * co.SecondShiftCashier.TimeToProcess()))
			}

			time.Sleep(sleepTime * time.Millisecond)

			co.CustomersProcessed++
		}
	}

	co.CurrentCustomer = customer.CustomerAgent{}
	co.ProcessingCustomer = false
}
