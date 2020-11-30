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
	amicability      float64
	//customer age 		 int       not implemented yet but could later otherwise item age rating has no meaning
	emergencyLeaveChance float64
	emergencyLeave       bool
	competence           float64
	trolleyLimit         int
	currentTrolleyCount  int
	finishedShop         bool
	inQueue              bool
	FloorStaffNearby     floorStaff.FloorStaff
	initialised          bool
	Occupied             bool
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewCustomer(UpperBound int, LowerBound int) CustomerAgent {
	ca := CustomerAgent{}

	ca.initialised = true

	//static values
	ca.impairmentFactor = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.amicability = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.competence = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.withChildren = (r.Intn(2) == 1)
	ca.loyaltyCard = (r.Intn(2) == 1)

	ca.trolleyLimit = r.Intn((UpperBound+1)-LowerBound) + LowerBound
	//ca.age = (r.Intn(100-14) + 14) will be removed tomorrow if not implemented at checkout

	//dynamic values
	ca.emergencyLeaveChance = 0.0
	ca.currentTrolleyCount = 0
	ca.emergencyLeave = false
	ca.finishedShop = false
	ca.inQueue = false

	// items
	ca.items = []item.ItemAgent{}

	return ca
}

func (ca *CustomerAgent) PropagateTime(ItemHandlingUpper float64, ItemHandlingLower float64) {

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

	// Reset to false after each iteration, floor staff only helps for 1 minute
	ca.Occupied = false

}

func (ca *CustomerAgent) SelectQueue(QueueLengths []int) int {
	selectedQueue := 0
	currentQueueLength := QueueLengths[0]

	for i, s := range QueueLengths {
		if s < currentQueueLength {
			selectedQueue = i
			currentQueueLength = QueueLengths[i]
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

func (ca *CustomerAgent) GetInitialised() bool {
	return ca.initialised
}

func (ca *CustomerAgent) GetAmicability() float64 {
	return ca.amicability
}

func (ca *CustomerAgent) addItemToTrolley(ItemHandlingUpper float64, ItemHandlingLower float64) {
	var isImpaired = ((ca.impairmentFactor) > (math.Round(((r.Float64()*(0.5))+0.4)*100) / 100))
	var itemAddBoost = 0.0
	var chanceItemAdded = 1.0

	if ca.withChildren && isImpaired {
		chanceItemAdded -= math.Round((r.Float64()*0.8)*100) / 100
	} else if ca.withChildren || isImpaired {
		chanceItemAdded -= math.Round((r.Float64()*0.5)*100) / 100
	} else {
		chanceItemAdded -= math.Round((r.Float64()*0.3)*100) / 100
	}

	if ca.Occupied {
		if ca.amicability*ca.FloorStaffNearby.GetAmicability() > ((r.Float64()*(0.3))+0.2)*100 {
			itemAddBoost = 1 + (ca.competence * ca.FloorStaffNearby.GetCompetence())
		}
	} else {
		itemAddBoost = 1 + (ca.competence / 5)
	}

	chanceItemAdded *= itemAddBoost

	if chanceItemAdded > 0.3 {
		ca.currentTrolleyCount++
		ca.items = append(ca.items, item.NewItem(ItemHandlingUpper, ItemHandlingLower))
	}
}
