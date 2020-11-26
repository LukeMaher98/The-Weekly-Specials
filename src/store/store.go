package store

import (
	"fmt"
	"math"
	"math/rand"
	"src/cashier"
	"src/checkout"
	"src/constants"
	"src/floorStaff"
	"src/manager"
	"time"

	// "src/item"
	"src/customer"
)

type StoreAgent struct {
	baseArrivalRate       float64
	baseInconvenience     float64
	statusRemainingTime   float64
	storeTempStatus       int
	CustomersOnFloor      []customer.CustomerAgent
	CustomersReadyToQueue []customer.CustomerAgent
	CustomerQueues        []constants.CustomerQueue
	FloorStaffFirstShift  []floorStaff.FloorStaff
	FloorStaffSecondShift []floorStaff.FloorStaff
	Checkouts             []checkout.CheckoutAgent
	ManagerFirstShift     manager.ManagerAgent
	ManagerSecondShift    manager.ManagerAgent
	ItemLimitBounds       constants.StoreAttributeBoundsInt
	ItemTimeBounds        constants.StoreAttributeBoundsFloat
}

// CreateStoreAgent : creates 'empty' Scenario for initialisation
func CreateStoreAgent() StoreAgent {
	return StoreAgent{
		0.0,
		0.0,
		0.0,
		0,
		[]customer.CustomerAgent{},
		[]customer.CustomerAgent{},
		[]constants.CustomerQueue{},
		[]floorStaff.FloorStaff{},
		[]floorStaff.FloorStaff{},
		[]checkout.CheckoutAgent{},
		manager.ManagerAgent{},
		manager.ManagerAgent{},
		constants.StoreAttributeBoundsInt{},
		constants.StoreAttributeBoundsFloat{},
	}
}

// CreateInitialisedStoreAgent : populates store agent
func CreateInitialisedStoreAgent(
	arrivalRates constants.StoreAttributeBoundsInt,
	itemLimits constants.StoreAttributeBoundsInt,
	itemTimes constants.StoreAttributeBoundsFloat,
	checkoutCount int,
	cashierShifts constants.StoreShifts,
	floorStaffShifts constants.StoreShifts,
	floorStaffAttributeBounds constants.StaffAttributeBounds,
	cashierAttributeBounds constants.StaffAttributeBounds,
	floorManagerAttributeBounds constants.StaffAttributeBounds,
) StoreAgent {
	newStore := CreateStoreAgent()

	newStore.ItemLimitBounds = itemLimits
	newStore.ItemTimeBounds = itemTimes

	for i := 0; i < checkoutCount; i++ {
		newStore.CustomerQueues = append(newStore.CustomerQueues, constants.CustomerQueue{})
		newStore.Checkouts = append(newStore.Checkouts, checkout.CreateInitialisedCheckoutAgent())
	}

	for i := 0; i < cashierShifts.FirstShiftCount; i++ {
		newStore.Checkouts[i].FirstShiftCashier = cashier.CreateInitialisedCashierAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
			floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound)
	}

	for i := 0; i < cashierShifts.SecondShiftCount; i++ {
		newStore.Checkouts[i].SecondShiftCashier = cashier.CreateInitialisedCashierAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
			floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound)
	}

	for i := 0; i < floorStaffShifts.FirstShiftCount; i++ {
		newStore.FloorStaffFirstShift = append(newStore.FloorStaffFirstShift, floorStaff.CreateInitialisedFloorStaffAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
			floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound))
	}

	for i := 0; i < floorStaffShifts.SecondShiftCount; i++ {
		newStore.FloorStaffSecondShift = append(newStore.FloorStaffSecondShift, floorStaff.CreateInitialisedFloorStaffAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
			floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound))
	}

	newStore.ManagerFirstShift = manager.CreateInitialisedFloorManagerAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
		floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound)

	newStore.ManagerSecondShift = manager.CreateInitialisedFloorManagerAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
		floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound)

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrivalRange := float64(arrivalRates.UpperBound - arrivalRates.LowerBound)
	newStore.baseArrivalRate = math.Round((seed.Float64()*arrivalRange)+float64(arrivalRates.LowerBound)) / float64(60)

	return newStore
}

// PropagateTime : propagates time to agents in store
func (s *StoreAgent) PropagateTime(currentShift int, currentDay int, currentTime float64, externalImpact float64) {
	//if math.Mod(currentTime, 60) == 0 {
		//fmt.Println("current arrival rate: ", getRateOfArrival(s.baseArrivalRate, currentDay, currentTime, externalImpact))
	//}
	/*go?*/ s.checkNewCustomers(getRateOfArrival(s.baseArrivalRate, currentDay, currentTime, externalImpact))
	/*go?*/ s.propagateStore(currentShift)
	/*go?*/ s.propagateCustomerQueues(currentShift)
	/*go?*/ s.propagateConcurrentCheckouts(currentShift)
}

func getRateOfArrival(baseRate float64, currentDay int, currentTime float64, externalImpact float64) float64 {
	rawValue := baseRate
	switch true {
	case currentDay == 0:
		rawValue *= 0.975
		break
	case currentDay == 2 || currentDay == 3:
		rawValue *= 1.025
		break
	case currentDay == 4:
		rawValue *= 1.0375
		break
	case currentDay == 5:
		rawValue *= 1.05
		break
	case currentDay == 6:
		rawValue *= 0.95
		break
	}
	if currentDay < 5 {
		if math.Abs(1020-currentTime) <= 60 {
			rawValue *= 1.05
		}
	} else {
		if math.Abs(900-currentTime) <= 60 {
			rawValue *= 1.05
		}
	}
	rawValue *= (1 + externalImpact)
	if rawValue < 0 {
		rawValue = 0
	} else if rawValue > 1 {
		rawValue = 1
	}
	return rawValue
}

func (s *StoreAgent) checkNewCustomers(rateOfArrival float64) {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64()*1.0 < rateOfArrival {
		s.CustomersOnFloor = append(s.CustomersOnFloor, customer.NewCustomer(s.ItemLimitBounds.UpperBound, s.ItemLimitBounds.LowerBound, s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound))
	}
}

func (s *StoreAgent) propagateCustomerQueues(currentShift int) {
	for index, customer := range s.CustomersOnFloor {
		if customer.IsFinishedShopping() {
			s.CustomersReadyToQueue = append(s.CustomersReadyToQueue, customer)
			s.CustomersOnFloor = append(s.CustomersOnFloor[:index], s.CustomersOnFloor[index+1:]...)
		}
	}

	openCheckoutNum := 0

	for _, checkout := range s.Checkouts {
		if checkout.IsManned(currentShift) {
			openCheckoutNum++
		}
	}

	queueLengths := s.getQueueLengths()

	fmt.Println("Queue: ", queueLengths, "Length: ", len(queueLengths))

	if len(queueLengths) > openCheckoutNum {
		fmt.Println("Queue: ", queueLengths, "Length: ", len(queueLengths))
		for i := openCheckoutNum; i < len(queueLengths); i++ {
			s.CustomerQueues[i].Mutex.Lock()
			for _, customer := range s.CustomerQueues[i].Queue {
				s.CustomersReadyToQueue = append(s.CustomersReadyToQueue, customer)
			}
			s.CustomerQueues = append(s.CustomerQueues[:i], s.CustomerQueues[i+1:]...)
		}
	} else if len(queueLengths) < openCheckoutNum {
		for i := len(queueLengths); i < openCheckoutNum; i++ {
			s.CustomerQueues = append(s.CustomerQueues, constants.CustomerQueue{})
		}
	}

	for _, customer := range s.CustomersReadyToQueue {
		queueIndex := customer.SelectQueue(queueLengths)
		s.CustomerQueues[queueIndex].Mutex.Lock()
		s.CustomerQueues[queueIndex].Queue = append(s.CustomerQueues[queueIndex].Queue, customer)
		s.CustomerQueues[queueIndex].Mutex.Unlock()
	}
}

func (s *StoreAgent) propagateConcurrentCheckouts(currentShift int) {
	for index, checkout := range s.Checkouts {
		if checkout.ProcessingCustomer == false {
			checkout.CurrentCustomerProgress = 0
			s.CustomerQueues[index].Mutex.Lock()
			if len(s.CustomerQueues[index].Queue) > 0 {
				checkout.CurrentCustomer = s.CustomerQueues[index].Queue[0]
				s.CustomerQueues[index].Queue = s.CustomerQueues[index].Queue[1:]
			}
			s.CustomerQueues[index].Mutex.Unlock()
			checkout.ProcessingCustomer = true
			go checkout.ProcessCustomer() // Processes items, calculates processing time, waits that amount of time then sets .ProcessingCustomer false
		}
	}
}

func (s *StoreAgent) propagateStore(currentShift int) {
	for _, customer := range s.CustomersOnFloor {
		customer.PropagateTime()
	}

	if currentShift == 0 {
		for _, staff := range s.FloorStaffFirstShift {
			staff.PropagateTime()
		}
		s.ManagerFirstShift.PropagateTime()
	} else {
		for _, staff := range s.FloorStaffSecondShift {
			staff.PropagateTime()
		}
		s.ManagerSecondShift.PropagateTime()
	}

	for _, checkout := range s.Checkouts {
		checkout.PropagateTime()
	}
}

func (s *StoreAgent) getQueueLengths() []int {
	queueLengths := []int{}
	for _, queue := range s.CustomerQueues {
		queue.Mutex.Lock()
		queueLengths = append(queueLengths, len(queue.Queue))
		queue.Mutex.Unlock()
	}
	return queueLengths
}
