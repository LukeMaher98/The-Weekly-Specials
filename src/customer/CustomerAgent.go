package customer

import (
	"math"
	"math/rand"
	"src/floorStaff"
	"src/item"
	"time"
)

type CustomerAgent struct {
	items            []item.ItemAgent
	impairmentFactor float64
	couponItem       float64
	withChildren     bool
	loyaltyCard      bool
	baseAmicability  float64
	preferredPayment float64 //checkout or cashier don't have payment options rn?
	//customer age 		int       not implemented yet but could later otherwise item age rating has no meaning
	emergencyLeaveChance float64
	emergencyLeave       bool
	competence           float64
	trolleyLimit         int
	currentTrolleyCount  int
	finishedShop         bool
	inQueue              bool
	floorStaffNearby     floorStaff.FloorStaff
}

func NewCustomer(UpperBound int, LowerBound int) CustomerAgent {
	ca := CustomerAgent{}

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	//static values
	ca.impairmentFactor = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.baseAmicability = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.preferredPayment = math.Round((r.Float64()*2)*100) / 100
	ca.competence = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.withChildren = (r.Intn(2) == 1)
	ca.loyaltyCard = (r.Intn(2) == 1)

	//ca.age = (r.Intn(100-14) + 14)

	ca.trolleyLimit = r.Intn((UpperBound+1)-LowerBound) + LowerBound

	//dynamic values
	ca.emergencyLeaveChance = 0.0
	ca.currentTrolleyCount = 0
	ca.emergencyLeave = false
	ca.finishedShop = false
	ca.inQueue = false

	//generate items
	ca.items = []item.ItemAgent{}

	return ca
}

func (ca *CustomerAgent) PropagateTime(ItemHandlingUpper float64, ItemHandlingLower float64) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	//Add item to trolley
	if ca.currentTrolleyCount < ca.trolleyLimit {
		ca.addItemToTrolley(ItemHandlingUpper, ItemHandlingLower)
	} else {
		ca.finishedShop = true
	}

	//Emergency Leave the store
	ca.emergencyLeaveChance = (math.Round((r.Float64()*1)*100) / 100)
	if ca.emergencyLeaveChance > 0.95 {
		ca.emergencyLeave = true
	}

	//Replace item while in Queue
	if ca.inQueue {
		replaceItem := (math.Round((r.Float64()*1)*100) / 100)
		if replaceItem > 0.95 {
			ca.inQueue = false
		}
	}

}

func (ca *CustomerAgent) SelectQueue(QueueLengths []int) int {
	selectedQueue := 0
	currentQueueLength := QueueLengths[0]

	for i, s := range QueueLengths {
		if s <= currentQueueLength {
			selectedQueue = i
		}
	}

	return selectedQueue
}

func (ca *CustomerAgent) IsFinishedShopping() bool {
	return ca.finishedShop
}

func (ca *CustomerAgent) IsJoingQueue() {
	ca.inQueue = true
}

func (ca *CustomerAgent) IsLeavingQueue() bool {
	return ca.inQueue
}

func (ca *CustomerAgent) EmergencyDeparture() bool {
	return ca.emergencyLeave
}

func (ca *CustomerAgent) GetCustomerItems() []item.ItemAgent {
	return ca.items
}

func (ca *CustomerAgent) addItemToTrolley(ItemHandlingUpper float64, ItemHandlingLower float64) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var itemSkipped = 0.0
	var isImpaired = ((ca.impairmentFactor) > (math.Round(((r.Float64()*(0.5))+0.4)*100) / 100))
	var helpedMultiplier = 0.0

	//isHelped := ca.floorStaffNearby.Occupied

	if ca.withChildren && isImpaired {
		itemSkipped = math.Round((r.Float64()*1)*100) / 100
	} else if ca.withChildren || isImpaired {
		itemSkipped = math.Round((r.Float64()*0.8)*100) / 100
	} else {
		itemSkipped = math.Round((r.Float64()*0.4)*100) / 100
	}

	//just waiting for carl in order to see if I need to change this
	/*	if isHelped {
		if ca.BaseAmicability > 0.55 && ca.FloorStaffNearby.BaseHelpfulness > 0.55 {
			helpedMultiplier = ca.Competence * ca.FloorStaffNearby.Competence
		}
	}*/

	itemSkipped = itemSkipped - helpedMultiplier

	if itemSkipped < 0.75 {
		ca.currentTrolleyCount++
		ca.items = append(ca.items, *item.NewItem(ItemHandlingUpper, ItemHandlingLower))
	}
}
