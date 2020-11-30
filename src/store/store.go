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
		floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound,
		newStore.FloorStaffFirstShift, newStore.Checkouts, 0)

	newStore.ManagerSecondShift = manager.CreateInitialisedFloorManagerAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
		floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound,
		newStore.FloorStaffSecondShift, newStore.Checkouts, 1)

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	if arrivalRates.UpperBound != arrivalRates.LowerBound {
		arrivalRange := float64(arrivalRates.UpperBound - arrivalRates.LowerBound)
		newStore.baseArrivalRate = math.Round((seed.Float64()*arrivalRange)+float64(arrivalRates.LowerBound)) / float64(60)
	} else {
		newStore.baseArrivalRate = float64(arrivalRates.LowerBound) / float64(60)
	}

	return newStore
}

// PropagateTime : propagates time to agents in store
func (s *StoreAgent) PropagateTime(currentShift int, currentDay int, currentTime float64, externalImpact float64) {
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
	for customerIndex := range s.CustomersOnFloor {
		customerIndex = customerIndex - removedCount
		if s.CustomersOnFloor[customerIndex].EmergencyDeparture() {
			s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{
				Customer: s.CustomersOnFloor[customerIndex],
				Day:      currentDay,
				Time:     currentTime,
				Reason:   "Emergency",
			})
			s.removeFloorCustomer(customerIndex)
			removedCount++
			continue
		} else if s.CustomersOnFloor[customerIndex].IsFinishedShopping() {
			s.CustomersReadyToQueue = append(s.CustomersReadyToQueue, s.CustomersOnFloor[customerIndex])
			s.removeFloorCustomer(customerIndex)
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
		}
		s.CustomerQueues = s.CustomerQueues[:openCheckoutNum]
		queueLengths = queueLengths[:openCheckoutNum]
	} else if len(queueLengths) < openCheckoutNum {
		for i := len(queueLengths); i < openCheckoutNum; i++ {
			s.CustomerQueues = append(s.CustomerQueues, []customer.CustomerAgent{})
			queueLengths = append(queueLengths, 0)
		}
	}

	removedCustomers := 0
	for customerIndex, customer := range s.CustomersReadyToQueue {
		customerIndex = customerIndex - removedCustomers
		if customer.EmergencyDeparture() == true {
			s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{
				Customer: customer,
				Day:      currentDay,
				Time:     currentTime,
				Reason:   "Emergency",
			})
			s.removeReadyCustomer(customerIndex)
			continue
		} else {
			queueIndex := customer.SelectQueue(queueLengths)
			if len(s.CustomerQueues[queueIndex]) < 6 {
				s.CustomerQueues[queueIndex] = append(s.CustomerQueues[queueIndex], customer)
			} else {
				s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{
					Customer: customer,
					Day:      currentDay,
					Time:     currentTime,
					Reason:   "Queues too long",
				})
			}
			s.removeReadyCustomer(customerIndex)
			removedCustomers++
		}
	}

	for queueIndex := range s.CustomerQueues {
		removedCustomers = 0
		for customerIndex := range s.CustomerQueues[queueIndex] {
			customerIndex = customerIndex - removedCustomers
			if s.CustomerQueues[queueIndex][customerIndex].EmergencyDeparture() == true {
				s.CustomersLost = append(s.CustomersLost, constants.CustomerLost{
					Customer: s.CustomerQueues[queueIndex][customerIndex],
					Day:      currentDay,
					Time:     currentTime,
					Reason:   "Emergency",
				})
				s.removeQueueCustomer(queueIndex, customerIndex)
				removedCustomers++
				continue
			} else if s.CustomerQueues[queueIndex][customerIndex].IsLeavingQueue() {
				s.CustomersOnFloor = append(s.CustomersReadyToQueue, s.CustomerQueues[queueIndex][customerIndex])
				s.removeQueueCustomer(queueIndex, customerIndex)
				removedCustomers++
			}
		}
	}
}

func (s *StoreAgent) propagateConcurrentCheckouts(currentShift int, currentDay int, currentTime float64) {
	for checkoutIndex := range s.Checkouts {
		if s.Checkouts[checkoutIndex].IsManned(currentShift) && s.Checkouts[checkoutIndex].ProcessingCustomer == false {
			s.Checkouts[checkoutIndex].CurrentCustomerProgress = 0
			if len(s.CustomerQueues[checkoutIndex]) > 0 {
				s.Checkouts[checkoutIndex].CurrentCustomer = s.CustomerQueues[checkoutIndex][0]
				s.CustomerQueues[checkoutIndex] = s.CustomerQueues[checkoutIndex][1:]
			}
			s.Checkouts[checkoutIndex].ProcessingCustomer = true
			go s.Checkouts[checkoutIndex].ProcessCustomer(currentShift)
		}
	}
}

func (s *StoreAgent) propagateStore(currentShift int, currentDay int, currentTime float64) {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	for customerIndex := range s.CustomersOnFloor {
		s.CustomersOnFloor[customerIndex].PropagateTime(s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound)
	}

	if currentShift == 0 {
		for i := range s.FloorStaffFirstShift {
			if len(s.CustomersOnFloor) != 0 {
				rand := r.Intn(len(s.CustomersOnFloor))
				if s.CustomersOnFloor[rand].GetAmicability()*s.FloorStaffFirstShift[i].GetAmicability() > ((r.Float64()*(0.3))+0.2)*100 {
					s.CustomersOnFloor[rand].FloorStaffNearby = s.FloorStaffFirstShift[i]
					s.CustomersOnFloor[rand].Occupied = true
				}
			}
		}
		s.ManagerFirstShift.PropagateTime()
	} else {
		for i := range s.FloorStaffSecondShift {
			if len(s.CustomersOnFloor) != 0 {
				rand := r.Intn(len(s.CustomersOnFloor))
				if s.CustomersOnFloor[rand].GetAmicability()*s.FloorStaffSecondShift[i].GetAmicability() > ((r.Float64()*(0.3))+0.2)*100 {
					s.CustomersOnFloor[rand].FloorStaffNearby = s.FloorStaffSecondShift[i]
					s.CustomersOnFloor[rand].Occupied = true
				}
			}
		}
		s.ManagerSecondShift.PropagateTime()
	}

	for customerIndex := range s.CustomersReadyToQueue {
		s.CustomersReadyToQueue[customerIndex].PropagateTime(s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound)
	}

	for queueIndex, queue := range s.CustomerQueues {
		for customerIndex := range queue {
			s.CustomerQueues[queueIndex][customerIndex].PropagateTime(s.ItemTimeBounds.UpperBound, s.ItemTimeBounds.LowerBound)
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

func (s *StoreAgent) removeFloorCustomer(customerIndex int) {
	if customerIndex == len(s.CustomersOnFloor) {
		s.CustomersOnFloor = append(s.CustomersOnFloor[:], s.CustomersOnFloor[:customerIndex]...)
	} else {
		s.CustomersOnFloor = append(s.CustomersOnFloor[:customerIndex], s.CustomersOnFloor[customerIndex+1:]...)
	}
}

func (s *StoreAgent) removeReadyCustomer(customerIndex int) {
	if customerIndex == len(s.CustomersReadyToQueue) {
		s.CustomersReadyToQueue = append(s.CustomersReadyToQueue[:], s.CustomersReadyToQueue[:customerIndex]...)
	} else {
		s.CustomersReadyToQueue = append(s.CustomersReadyToQueue[:customerIndex], s.CustomersReadyToQueue[customerIndex+1:]...)
	}
}

func (s *StoreAgent) removeQueueCustomer(queueIndex int, customerIndex int) {
	if customerIndex == len(s.CustomerQueues[queueIndex]) {
		s.CustomerQueues[queueIndex] = append(s.CustomerQueues[queueIndex][:], s.CustomerQueues[queueIndex][:customerIndex]...)
	} else {
		s.CustomerQueues[queueIndex] = append(s.CustomerQueues[queueIndex][:customerIndex], s.CustomerQueues[queueIndex][customerIndex+1:]...)
	}
}

func (s *StoreAgent) ResetDay() {
	s.CustomersOnFloor = []customer.CustomerAgent{}
	s.CustomersReadyToQueue = []customer.CustomerAgent{}
	s.CustomerQueues = [][]customer.CustomerAgent{}
	for _, checkout := range s.Checkouts {
		checkout.ProcessingCustomer = false
	}
}

func (s *StoreAgent) PrintResults() {
	totalCustomersProcessed := 0
	totalMonetaryIntake := 0.0
	fmt.Println("Scenario Results:")
	fmt.Println("------------------")
	for index, checkout := range s.Checkouts {
		fmt.Println("[Checkout ", index, "] Customers Processed: ", checkout.CustomersProcessed, " Monetary Intake: ", checkout.TotalMoney)
		totalCustomersProcessed += checkout.CustomersProcessed
		totalMonetaryIntake += checkout.TotalMoney
	}
	fmt.Println("\n[Total] Customers Processed: ", totalCustomersProcessed, " Monetary Intake: ", totalMonetaryIntake)
	emergencyLostCustomers := 0
	inconvenienceLostCustomers := 0
	for _, lostCustomer := range s.CustomersLost {
		if lostCustomer.Reason == "Emergency" {
			emergencyLostCustomers++
		} else {
			inconvenienceLostCustomers++
		}
	}
	fmt.Println("\nCustomers Lost: ", len(s.CustomersLost))
	fmt.Println("Customers Lost due to Emergency: ", emergencyLostCustomers)
	fmt.Println("Customers Lost due to Inconvenience", inconvenienceLostCustomers)
}
