package cashier

import (
	"math"
	"math/rand"
	"time"
)

// CashierAgent : the cashier struct
type CashierAgent struct {
	amicability  float64
	competence   float64
	cashHandling float64
	cardHandling float64
	managerBoost float64
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// CreateInitialisedCashierAgent : creates new cashier
func CreateInitialisedCashierAgent(amicLower, amicUpper, compLower, compUpper float64) CashierAgent {
	cashier := CashierAgent{}

	cashier.amicability = math.Round(((r.Float64()*(amicUpper-amicLower))+amicLower)*100) / 100
	cashier.competence = math.Round(((r.Float64()*(compUpper-compLower))+compLower)*100) / 100
	cashier.cashHandling = math.Round(((r.Float64()*0.25)+0.25)*100) / 100
	cashier.cardHandling = math.Round(((r.Float64()*0.25)+0.25)*100) / 100
	cashier.managerBoost = 1.00

	return cashier
}

// TimeToProcess : returns the time to process x items
func (cashier *CashierAgent) TimeToProcess() float64 {
	return 1.2 - ((cashier.competence * cashier.managerBoost) / 2.5)
}

// ManagerPresent : applies a boost to the cashier
func (cashier *CashierAgent) ManagerPresent(boost float64) {
	cashier.managerBoost = boost
}

// ManagerAbsent : reverts the boost to the cashier
func (cashier *CashierAgent) ManagerAbsent() {
	cashier.managerBoost = 1
}

// GetAmicability : returns cashier amicability
func (cashier *CashierAgent) GetAmicability() float64 {
	return cashier.amicability
}
