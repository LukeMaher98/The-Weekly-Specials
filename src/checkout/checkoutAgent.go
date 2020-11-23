package checkout

import (
	"math"
	"math/rand"
	"time"
)

type CheckoutAgent struct {
	selfCheckout       bool
	adultCheckout      bool
	assistanceWaitTime float64
	totalMoney         float64
	// currentCashier cashier
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// checkout agent constructor
func NewCheckout() *CheckoutAgent {
	co := CheckoutAgent{}

	// Randomly Initialised Variables
	co.selfCheckout = (r.Intn(2) == 1)
	co.adultCheckout = (r.Intn(2) == 1)
	co.assistanceWaitTime = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	co.totalMoney = 0

	return &co
}

// Getter for selfCheckout
func (co *CheckoutAgent) IsSelfCheckout() bool {
	return co.selfCheckout
}

// Getter for adultCheckout
func (co *CheckoutAgent) IsAdultCheckout() bool {
	return co.adultCheckout
}

// Add money of an item to the checkout
func (co *CheckoutAgent) AddMoney(price float64) {
	co.totalMoney += price
}

// Get the money currently in the checkout
func (co *CheckoutAgent) GetMoney() float64 {
	return co.totalMoney
}
