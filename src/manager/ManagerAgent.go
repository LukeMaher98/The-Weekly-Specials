package manager

import (
	"math"
	"math/rand"
	"src/cashier"
	"src/floorStaff"
	"time"
)

// ManagerAgent : the manager struct
type ManagerAgent struct {
	amicability    float64
	competence     float64
	onFloor        bool
	currentCashier *cashier.CashierAgent
	floorStaff     []floorStaff.FloorStaff
	cashiers       []cashier.CashierAgent
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewManager : creates new manager on floor
func NewManager(amicUpper, amicLower, compUpper, compLower float64, staff []floorStaff.FloorStaff, cashiers []cashier.CashierAgent) ManagerAgent {
	manager := ManagerAgent{}

	manager.amicability = math.Round(((r.Float64()*(amicUpper-amicLower))+amicLower)*100) / 100
	manager.competence = math.Round(((r.Float64()*(compUpper-compLower))-compLower)*100) / 100
	manager.onFloor = true
	manager.floorStaff = staff
	manager.cashiers = cashiers

	return manager
}

// PropogateTime : propogates time for the manager
func (mngr *ManagerAgent) PropogateTime() {
	// 1/4 chance of moving
	if r.Float64() < 0.25 {
		if r.Float64() < 0.5 {
			mngr.WorkTheFloor()
		} else {
			mngr.SuperviseCashier()
		}
	}
}

// WorkTheFloor : moves manager to floor or keeps them there
func (mngr *ManagerAgent) WorkTheFloor() {
	if !mngr.onFloor {
		mngr.onFloor = true
		mngr.currentCashier.ManagerAbsent()
		mngr.currentCashier = nil

		for _, staff := range mngr.floorStaff {
			staff.ManagerPresent(mngr.competence)
		}
	}
}

// SuperviseCashier : manager supervises a checkout
func (mngr *ManagerAgent) SuperviseCashier() {
	if mngr.onFloor {
		mngr.onFloor = false
		for _, staff := range mngr.floorStaff {
			staff.ManagerAbsent()
		}
	}

	randomIndex := r.Intn(len(mngr.cashiers))
	pick := mngr.cashiers[randomIndex]
	mngr.currentCashier = &pick
	mngr.currentCashier.ManagerPresent(mngr.amicability)
}
