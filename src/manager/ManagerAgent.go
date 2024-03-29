package manager

import (
	"math"
	"math/rand"
	"src/cashier"
	"src/checkout"
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

// CreateInitialisedFloorManagerAgent : creates new manager on floor
func CreateInitialisedFloorManagerAgent(amicLower, amicUpper, compLower, compUpper float64, staff []floorStaff.FloorStaff, co []checkout.CheckoutAgent, shift int) ManagerAgent {
	manager := ManagerAgent{}

	manager.amicability = math.Round(((r.Float64()*(amicUpper-amicLower))+amicLower)*100) / 100
	manager.competence = math.Round(((r.Float64()*(compUpper-compLower))-compLower)*100) / 100
	manager.onFloor = true
	manager.floorStaff = staff

	manager.cashiers = getCashiers(co, shift)

	return manager
}

// PropagateTime : propogates time for the manager
func (mngr *ManagerAgent) PropagateTime() {
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

		for i := range mngr.floorStaff {
			if mngr.floorStaff[i].GetAmicability()*mngr.amicability > ((r.Float64()*(0.3))+0.2)*100 {
				mngr.floorStaff[i].ManagerPresent(mngr.competence)
			}
		}
	}
}

// SuperviseCashier : manager supervises a checkout
func (mngr *ManagerAgent) SuperviseCashier() {
	if mngr.onFloor {
		mngr.onFloor = false
		for i := range mngr.floorStaff {
			mngr.floorStaff[i].ManagerAbsent()
		}
	}

	randomIndex := r.Intn(len(mngr.cashiers))
	pick := mngr.cashiers[randomIndex]
	mngr.currentCashier = &pick
	if mngr.currentCashier.GetAmicability()*mngr.amicability > ((r.Float64()*(0.3))+0.2)*100 {
		mngr.currentCashier.ManagerPresent(mngr.competence)
	}
}

func getCashiers(co []checkout.CheckoutAgent, shift int) []cashier.CashierAgent {
	var cash []cashier.CashierAgent

	if shift == 0 {
		for _, c := range co {
			if c.IsManned(shift) {
				cash = append(cash, c.FirstShiftCashier)
			}
		}
	} else if shift == 1 {
		for _, c := range co {
			if c.IsManned(shift) {
				cash = append(cash, c.SecondShiftCashier)
			}
		}
	}

	return cash
}
