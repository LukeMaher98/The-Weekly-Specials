package checkout

import (
	"src/cashier"
	"src/customer"
	"time"
)

type CheckoutAgent struct {
	ProcessingCustomer bool
	CurrentCustomerProgress float64
	TotalMoney         float64
	FirstShiftCashier  cashier.CashierAgent
	SecondShiftCashier cashier.CashierAgent
	CurrentCustomer    customer.CustomerAgent
	CustomersProcessed int
}

// checkout agent constructor
func CreateInitialisedCheckoutAgent() CheckoutAgent {
	co := CheckoutAgent{}

	// Randomly Initialised Variables
	co.ProcessingCustomer = false
	co.CurrentCustomerProgress = 0
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

func (co *CheckoutAgent) ProcessCustomer(shift int) {

	// Only process actual customers
	if co.CurrentCustomer.GetInitialised() == true {
		for co.ProcessingCustomer == true {
			customerTotal := 0.0
			for _, item := range co.CurrentCustomer.GetCustomerItems() {
				if co.CurrentCustomer.GetAge() < 18 && item.IsAgeRated() {
					// Skips item 
				} else {
					co.CurrentCustomerProgress += item.GetItemHandling()
					customerTotal += item.GetPrice()
				}
			}
			co.TotalMoney += customerTotal

			sleepTime := time.Millisecond
			if shift == 0 {
				sleepTime = time.Duration(int((co.CurrentCustomerProgress/30) * co.FirstShiftCashier.TimeToProcess()))
			} else if shift == 1 {
				sleepTime = time.Duration(int((co.CurrentCustomerProgress/30) * co.SecondShiftCashier.TimeToProcess()))
			}

			if co.CurrentCustomer.GetCashPreference() {
				sleepTime += time.Duration(1.2 - ((co.FirstShiftCashier.GetAmicability() * co.CurrentCustomer.GetAmicability()) / 2.5)) / 10
			} else {
				sleepTime += time.Duration(1.2 - ((co.FirstShiftCashier.GetAmicability() * co.CurrentCustomer.GetAmicability()) / 2.5)) / 30
			}

			time.Sleep(sleepTime * time.Millisecond)

			co.CustomersProcessed++
		}
	}

	co.CurrentCustomer = customer.CustomerAgent{}
	co.ProcessingCustomer = false
}

func (co *CheckoutAgent) ProcessSelf(shift int) {

	// Only process actual customers
	if co.CurrentCustomer.GetInitialised() == true {
		for co.ProcessingCustomer == true {
			customerTotal := 0.0
			for _, item := range co.CurrentCustomer.GetCustomerItems() {
				if co.CurrentCustomer.GetAge() < 18 && item.IsAgeRated() {
					// Skips item 
				} else {
					co.CurrentCustomerProgress += item.GetItemHandling()
					customerTotal += item.GetPrice()
				}
			}
			co.TotalMoney += customerTotal

			sleepTime := time.Millisecond
			if shift == 0 {
				sleepTime = time.Duration(int((co.CurrentCustomerProgress/30) * co.CurrentCustomer.TimeToProcess()))
			} else if shift == 1 {
				sleepTime = time.Duration(int((co.CurrentCustomerProgress/30) * co.CurrentCustomer.TimeToProcess()))
			}

			if co.CurrentCustomer.GetCashPreference() {
				sleepTime += time.Duration(1) / 10
			} else {
				sleepTime += time.Duration(1) / 30
			}

			time.Sleep(sleepTime * time.Millisecond)

			co.CustomersProcessed++
		}
	}

	co.CurrentCustomer = customer.CustomerAgent{}
	co.ProcessingCustomer = false
}
