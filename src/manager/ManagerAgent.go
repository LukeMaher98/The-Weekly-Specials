package manager

import (
	"math"
	"math/rand"
	"time"
)

type ManagerAgent struct {
	amicability float64
	competence  float64
	onFloor     bool
	//WorkingCheckout checkout.CheckoutAgent
	//SupervisingCashier cashier.CashierAgent

	observerList []observer
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewManager : creates new manager on floor
func NewManager(amicUpper, amicLower, compUpper, compLower float64) ManagerAgent {
	manager := ManagerAgent{}

	manager.amicability = math.Round(((r.Float64()*(amicUpper-amicLower))+amicLower)*100) / 100
	manager.competence = math.Round(((r.Float64()*(compUpper-compLower))-compLower)*100) / 100
	manager.onFloor = true

	return manager
}

// PropogateTime : propogates time for the manager
func (mngr *ManagerAgent) PropogateTime() {
	//check queue lengths, if queue.length > x && checkout empty man checkout

	// 1/4 chance of moving
	if r.Float64() < 0.25 {
		if r.Float64() < 0.5 {
			if mngr.onFloor == true {
				mngr.onFloor = false
			} else {
				mngr.onFloor = true
			}
		} else {
			//go to random checkout
		}
	}
}

func (mngr *ManagerAgent) WorkCheckout() {
	//queues[] = getQueueLengths()
}

func (mngr *ManagerAgent) register(o observer) {
	mngr.observerList = append(mngr.observerList, o)
}

func (mngr *ManagerAgent) notifyAll() {
	for _, observer := range mngr.observerList {
		observer.update(mngr.competence)
	}
}

//clear whole slice at end of shift and repopulate from scratch
func (mngr *ManagerAgent) clearSlice() {
	mngr.observerList = nil
}
