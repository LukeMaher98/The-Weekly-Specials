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
func IsSelfCheckout(co checkoutAgent) (bool) {
	return co.selfCheckout
}

// Getter for adultCheckout 
func IsAdultCheckout(co checkoutAgent) (bool) {
	return co.adultCheckout
}

// Add money of an item to the checkout
func AddMoney(co checkoutAgent, price float64) (checkoutAgent) {
	co.totalMoney += price
	return co
}

// Get the money currently in the checkout
func GetMoney(co checkoutAgent) (float64) {
	return co.totalMoney
}

// Set a new cashier to the checkout
// func SetCashier (co checkoutAgent, cashier Cashier) (checkoutAgent) {
// 	co.cashier = cashier
// 	return co
// }
