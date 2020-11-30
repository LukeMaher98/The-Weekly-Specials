package customer

import (
	"math"
	"math/rand"
	"src/floorStaff"
	"src/item"
	"time"
)

type CustomerAgent struct {
	items                []item.ItemAgent
	impairmentFactor     float64
	couponItem           float64
	withChildren         bool
	loyaltyCard          bool
	amicability          float64
	age                  int
	cashPreference       bool
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
	lowerBound           int
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
	ca.cashPreference = (r.Intn(2) == 1)
	ca.lowerBound = LowerBound

	ca.trolleyLimit = r.Intn((UpperBound+1)-LowerBound) + LowerBound
	ca.age = (r.Intn(90-14) + 14)

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
		ca.addItemsToTrolley(ItemHandlingUpper, ItemHandlingLower)
	} else {
		ca.finishedShop = true
	}

	//Emergency Leave the store
	ca.emergencyLeaveChance = (math.Round((r.Float64()*1)*100) / 100.00000000000000)
	if ca.emergencyLeaveChance > 0.99999999999999 {
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

	queueLengths := QueueLengths
	adjVal := 0

	if ca.trolleyLimit > 25 && ca.lowerBound < 25 {
		queueLengths = QueueLengths[1:]
		adjVal = 1
	}

	currentQueueLength := QueueLengths[0]
	var queuesOfSameLength []int

	for i, s := range queueLengths {
		if s < currentQueueLength {
			selectedQueue = i
			currentQueueLength = queueLengths[i]
			queuesOfSameLength = nil
		} else if s == currentQueueLength {
			queuesOfSameLength = append(queuesOfSameLength, i)
		}

		if queuesOfSameLength != nil {
			randomIndex := rand.Intn(len(queuesOfSameLength))
			selectedQueue = queuesOfSameLength[randomIndex]
		}
	}

	return selectedQueue + adjVal
}

func (ca *CustomerAgent) GetAge() int {
	return ca.age
}

func (ca *CustomerAgent) GetCashPreference() bool {
	return ca.cashPreference
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

func (ca *CustomerAgent) TimeToProcess() float64 {
	return 1.2 - ((ca.competence) / 2.5)
}

func (ca *CustomerAgent) addItemsToTrolley(ItemHandlingUpper float64, ItemHandlingLower float64) {
	var itemsAddedToTrolley = r.Intn(5) + 1

	for i := 0; i < itemsAddedToTrolley; i++ {
		var isImpaired = ((ca.impairmentFactor) > (math.Round(((r.Float64()*(0.5))+0.4)*100) / 100))
		var itemAddBoost = 0.0
		var chanceItemAdded = 1.0
		var underAge = ca.age < 18
		var abilityToAddAgeRestrictedItem = false

		if !underAge {
			abilityToAddAgeRestrictedItem = true
		}

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
			var newAddedItem = item.NewItem(ItemHandlingUpper, ItemHandlingLower)

			if underAge && newAddedItem.IsAgeRated() {
				abilityToAddAgeRestrictedItem = (math.Round((r.Float64()*1)*100) / 100) > 0.95
			}

			if !newAddedItem.IsAgeRated() {
				ca.items = append(ca.items, newAddedItem)
				ca.currentTrolleyCount++
			} else if abilityToAddAgeRestrictedItem {
				ca.items = append(ca.items, newAddedItem)
				ca.currentTrolleyCount++
			}
		}
	}
}
