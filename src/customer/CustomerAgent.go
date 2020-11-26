package customer

import (
	"math"
	"math/rand"
	"src/floorStaff"
	"src/item"
	"time"
)

type CustomerAgent struct {
	Items                []item.ItemAgent
	ImpairmentFactor     float64
	CouponItem           float64
	WithChildren         bool
	LoyaltyCard          bool
	BaseAmicability      float64
	PreferredPayment     float64 //checkout or cashier don't have payment options rn?
	EmergencyLeaveChance float64
	EmergencyLeave       bool
	Competence           float64
	TrolleyLimit         int
	CurrentTrolleyCount  int
	FinishedShop         bool
	InQueue              bool
	FloorStaffNearby     floorStaff.FloorStaff
}

func NewCustomer(UpperBound int, LowerBound int) *CustomerAgent {
	ca := CustomerAgent{}

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	//static values
	ca.ImpairmentFactor = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.BaseAmicability = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.PreferredPayment = math.Round((r.Float64()*2)*100) / 100
	ca.Competence = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.WithChildren = (r.Intn(2) == 1)
	ca.LoyaltyCard = (r.Intn(2) == 1)

	ca.TrolleyLimit = r.Intn(UpperBound-LowerBound) + LowerBound

	//dynamic values
	ca.EmergencyLeaveChance = 0.0
	ca.CurrentTrolleyCount = 0
	ca.EmergencyLeave = false
	ca.FinishedShop = false
	ca.InQueue = false

	//generate items
	ca.Items = []item.ItemAgent{}

	return &ca
}

func (ca CustomerAgent) PropagateTime() {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	//Add item to trolley
	if ca.CurrentTrolleyCount <= ca.TrolleyLimit {
		ca.addItemToTrolley()
	} else {
		ca.FinishedShop = true
	}

	//Emergency Leave the store
	ca.EmergencyLeaveChance = (math.Round((r.Float64()*1)*100) / 100)
	if ca.EmergencyLeaveChance > 0.95 {
		ca.EmergencyLeave = true
	}

	//Replace item while in Queue
	if ca.InQueue {
		replaceItem := (math.Round((r.Float64()*1)*100) / 100)
		if replaceItem > 0.95 {
			ca.InQueue = false
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

func (ca CustomerAgent) IsFinishedShopping() bool {
	return ca.FinishedShop
}

func (ca CustomerAgent) IsJoingQueue() {
	ca.InQueue = true
}

func (ca CustomerAgent) IsLeavingQueue() bool {
	return ca.InQueue
}

func (ca CustomerAgent) GetCustomerItems() []item.ItemAgent {
	return ca.Items
}

func (ca *CustomerAgent) addItemToTrolley() {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var itemSkipped = 0.0
	var isImpaired = ((ca.ImpairmentFactor) > (math.Round(((r.Float64()*(0.5))+0.4)*100) / 100))
	var helpedMultiplier = 0.0

	isHelped := ca.FloorStaffNearby.Occupied

	if ca.WithChildren && isImpaired {
		itemSkipped = math.Round((r.Float64()*1)*100) / 100
	} else if ca.WithChildren || isImpaired {
		itemSkipped = math.Round((r.Float64()*0.8)*100) / 100
	} else {
		itemSkipped = math.Round((r.Float64()*0.4)*100) / 100
	}

	if isHelped {
		if ca.BaseAmicability > 0.55 && ca.FloorStaffNearby.BaseHelpfulness > 0.55 {
			helpedMultiplier = ca.Competence * ca.FloorStaffNearby.Competence
		}
	}

	itemSkipped = itemSkipped - helpedMultiplier

	if itemSkipped < 0.75 {
		ca.Items = append(ca.Items, *item.NewItem())
	}
}
