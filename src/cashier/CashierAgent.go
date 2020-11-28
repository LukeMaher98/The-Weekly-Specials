package cashier

import (
	"math"
	"math/rand"
	"time"
)

type CashierAgent struct {
	Amicability  float64
	Competence   float64
	CashHandling float64
	CardHandling float64
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewCashier : creates new cashier
func CreateInitialisedCashierAgent(amicLower, amicUpper, compLower, compUpper float64) CashierAgent {
	cashier := CashierAgent{}

	cashier.Amicability = math.Round(((r.Float64()*(amicUpper-amicLower))+amicLower)*100) / 100
	cashier.Competence = math.Round(((r.Float64()*(compUpper-compLower))+compLower)*100) / 100
	cashier.CashHandling = math.Round(((r.Float64()*0.5)+0.25)*100) / 100
	cashier.CardHandling = math.Round(((r.Float64()*0.5)+0.25)*100) / 100

	return cashier
}
