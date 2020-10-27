package agents

type checkoutAgent struct {
	selfCheckout bool
	adultCheckout bool 
	assistanceWaitTime float64
	totalMoney float64
	// currentCashier cashier
}

// checkout agent constructor 
func ConstructorCheckout (self, adult bool) (checkoutAgent) {
	var co = checkoutAgent {
		self,
		adult,
		// How is this defined?
		// PLACEHOLDER
		0.5,
		0,
	}
	return co
}

// Getter for selfCheckout 
func (co *checkoutAgent) IsSelfCheckout() (bool) {
	return co.selfCheckout
}

// Getter for adultCheckout 
func (co *checkoutAgent) IsAdultCheckout() (bool) {
	return co.adultCheckout
}

// Add money of an item to the checkout
func (co *checkoutAgent) AddMoney(price float64) {
	co.totalMoney += price
}

// Get the money currently in the checkout
func (co *checkoutAgent) GetMoney() (float64) {
	return co.totalMoney
}

// Set a new cashier to the checkout
// func (co *checkoutAgent) SetCashier (cashier Cashier) {
// 	co.cashier = cashier
// }
