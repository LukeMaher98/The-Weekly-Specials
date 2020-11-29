package store

import (
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
	CustomerQueues        [][]customer.CustomerAgent
	CustomersLost         []constants.CustomerLost
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
		[][]customer.CustomerAgent{},
		[]constants.CustomerLost{},
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
		newStore.Checkouts = append(newStore.Checkouts, checkout.CreateInitialisedCheckoutAgent())
	}

	for i := 0; i < cashierShifts.FirstShiftCount; i++ {
		newStore.CustomerQueues = append(newStore.CustomerQueues, []customer.CustomerAgent{})
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
	// for i := range s.CustomerQueues {
	// 	if(len(s.CustomerQueues[i]) > 0){
	// 		fmt.Println("Floor Customers:", len(s.CustomersOnFloor))
	// 		fmt.Println("Lost Customers:", len(s.CustomersLost))
	// 		fmt.Println("RTQ Customers:", len(s.CustomersReadyToQueue))
	// 		for i := range s.CustomerQueues {
	// 			fmt.Println("Queue:", i, len(s.CustomerQueues[i]))
	// 		}
	// 	}
	// }

	s.checkNewCustomers(getRateOfArrival(s.baseArrivalRate, currentDay, currentTime, externalImpact))
	s.propagateStore(currentShift, currentDay, currentTime)
	s.propagateCustomerQueues(currentShift, currentDay, currentTime)
	s.propagateConcurrentCheckouts(currentShift, currentDay, currentTime)
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
		s.CustomersOnFloor = append(s.CustomersOnFloor, customer.NewCustomer(s.ItemLimitBounds.UpperBound, s.ItemLimitBounds.LowerBound))
	}
}

func (s *StoreAgent) propagateCustomerQueues(currentShift int, currentDay int, currentTime float64) {
	removedCount := 0
	for i := range s.CustomersOnFloor {
		i = i - removedCount
		if s.CustomersOnFloor[i].EmergencyDeparture() {
			s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{Customer: s.CustomersOnFloor[i], Day: currentDay, Time: currentTime, Reason: "Emergency"})
			s.CustomersOnFloor = append(s.CustomersOnFloor[:i], s.CustomersOnFloor[i+1:]...)
			removedCount++
			continue
		} else if s.CustomersOnFloor[i].IsFinishedShopping() {
			s.CustomersReadyToQueue = append(s.CustomersReadyToQueue, s.CustomersOnFloor[i])
			s.CustomersOnFloor = append(s.CustomersOnFloor[:i], s.CustomersOnFloor[i+1:]...)
			removedCount++
		}
	}

	openCheckoutNum := 0

	for _, checkout := range s.Checkouts {
		if checkout.IsManned(currentShift) {
			openCheckoutNum++
		}
	}

	queueLengths := s.getQueueLengths()

	if len(queueLengths) > openCheckoutNum {
		for i := openCheckoutNum; i < len(queueLengths); i++ {
			for _, customer := range s.CustomerQueues[i] {
				s.CustomersReadyToQueue = append(s.CustomersReadyToQueue, customer)
			}
			s.CustomerQueues = append(s.CustomerQueues[:i], s.CustomerQueues[i+1:]...)
		}
	} else if len(queueLengths) < openCheckoutNum {
		for i := len(queueLengths); i < openCheckoutNum; i++ {
			s.CustomerQueues = append(s.CustomerQueues, []customer.CustomerAgent{})
		}
	}

	removedCustomers := 0
	for index, customer := range s.CustomersReadyToQueue {
		index = index - removedCustomers
		if customer.EmergencyDeparture() {
			s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{Customer: customer, Day: currentDay, Time: currentTime, Reason: "Emergency"})
			s.CustomersReadyToQueue = append(s.CustomersReadyToQueue[:index], s.CustomersReadyToQueue[index+1:]...)
			continue
		} else {
			queueIndex := customer.SelectQueue(queueLengths)
			if len(s.CustomerQueues[queueIndex]) < 6 {
				s.CustomerQueues[queueIndex] = append(s.CustomerQueues[queueIndex], customer)
			} else {
				s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{Customer: customer, Day: currentDay, Time: currentTime, Reason: "Queues too long"})
			}
			if (index == len(s.CustomersReadyToQueue)) {
				s.CustomersReadyToQueue = append(s.CustomersReadyToQueue[:], s.CustomersReadyToQueue[:index]...)
			} else {
				s.CustomersReadyToQueue = append(s.CustomersReadyToQueue[:index], s.CustomersReadyToQueue[index+1:]...)
			}
			removedCustomers++
		}
	}

	for i := range s.CustomerQueues {
		removedCustomers = 0
		for j := range s.CustomerQueues[i] {
			j = j - removedCustomers
			if s.CustomerQueues[i][j].EmergencyDeparture() {
				s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{Customer: s.CustomerQueues[i][j], Day: currentDay, Time: currentTime, Reason: "Emergency"})
				if (j == len(s.CustomerQueues[i])) {
					s.CustomerQueues[i] = append(s.CustomerQueues[i][:], s.CustomerQueues[i][:j]...)
				} else {
					s.CustomerQueues[i] = append(s.CustomerQueues[i][:j], s.CustomerQueues[i][j+1:]...)
				}
				removedCustomers++
				continue
			} else if s.CustomerQueues[i][j].IsLeavingQueue() {
				s.CustomersOnFloor = append(s.CustomersReadyToQueue, s.CustomerQueues[i][j])
				if (j == len(s.CustomerQueues[i])) {
					s.CustomerQueues[i] = append(s.CustomerQueues[i][:], s.CustomerQueues[i][:j]...)
				} else {
					s.CustomerQueues[i] = append(s.CustomerQueues[i][:j], s.CustomerQueues[i][j+1:]...)
				}
				removedCustomers++
			}
		}
	}
}

func (s *StoreAgent) propagateConcurrentCheckouts(currentShift int, currentDay int, currentTime float64) {
	for i := range s.Checkouts {
		if s.Checkouts[i].ProcessingCustomer == false {
			s.Checkouts[i].CurrentCustomerProgress = 0
			if len(s.CustomerQueues[i]) > 0 {
				s.Checkouts[i].CurrentCustomer = s.CustomerQueues[i][0]
				s.CustomerQueues[i] = s.CustomerQueues[i][1:]
			}
			//fmt.Println(s.Checkouts[i].CurrentCustomer)
			s.Checkouts[i].ProcessingCustomer = true
			go s.Checkouts[i].ProcessCustomer(s.ItemTimeBounds)
			// print(s.Checkouts[i].TotalMoney)
		}
	}
}

func (s *StoreAgent) propagateStore(currentShift int, currentDay int, currentTime float64) {
	for i := range s.CustomersOnFloor {
		s.CustomersOnFloor[i].PropagateTime(s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound)
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

	for i := range s.CustomersReadyToQueue {
		s.CustomersReadyToQueue[i].PropagateTime(s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound)
	}

	for i, queue := range s.CustomerQueues {
		for j := range queue {
			s.CustomerQueues[i][j].PropagateTime(s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound)
		}
	}
}

func (s *StoreAgent) getQueueLengths() []int {
	queueLengths := []int{}
	for _, queue := range s.CustomerQueues {
		queueLengths = append(queueLengths, len(queue))
	}
	return queueLengths
}
