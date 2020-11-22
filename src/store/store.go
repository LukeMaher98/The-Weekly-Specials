package store

import (
	"math"
	"math/rand"
	"src/constants"
	"time"

	// "src/cashier"
	"src/checkout"
	"src/floorStaff"
	// "src/floorManager"
	// "src/item"
	// "src/customer"
)

type StoreAgent struct {
	baseArrivalRate     float64
	baseInconvenience   float64
	statusRemainingTime float64
	storeTempStatus     int

	//CustomersOnFloor        []customer.CustomerAgent
	//CustomersQueues         [][]customer.CustomerAgent
	FloorStaffFirstShift   	 []floorStaff.FloorStaffAgent
	FloorStaffSecondShift  	 []floorStaff.FloorStaffAgent
	Checkouts 			     []checkout.CheckoutAgent
	//FloorManagerFirstShift  floorManager.FloorManagerAgent
	//FloorManagerSecondShift floorManager.FloorManagerAgent
}

// CreateScenarioAgent : creates 'empty' Scenario for initialisation
func CreateStoreAgent() StoreAgent {
	return StoreAgent{
		0.0,
		0.0,
		0.0,
		0,
		//[]CustomerAgent{},
		//[][]CustomerAgent{},
		[]floorStaff.FloorStaffAgent{},
		[]floorStaff.FloorStaffAgent{},
		[]checkout.CheckoutAgent{},
		//FloorManagerAgent,
		//FloorManagerAgent
	}
}

// CreateInitialisedStoreAgent : populates store agent
func CreateInitialisedStoreAgent(
	arrivalRates constants.StoreAttributeBoundsInt,
	checkoutCount int,
	cashierShifts constants.StoreShifts,
	floorStaffShifts constants.StoreShifts,
	floorStaffAttributeBounds constants.StaffAttributeBounds,
	cashierAttributeBounds constants.StaffAttributeBounds,
	floorManagerAttributeBounds constants.StaffAttributeBounds,
) StoreAgent {
	newStore := CreateStoreAgent()

	for i := 0; i < checkoutCount; i++ {
		//newStore.CustomerQueues = append(newStore.CustomerQueues, []CustomerAgent{})
		newStore.Checkouts = append(newStore.Checkouts, checkout.CreateInitialisedCheckoutAgent())
	}

	for i := 0; i < cashierShifts.FirstShiftCount; i++ {
		//newStore.Checkouts[i].FirstShiftCashier = cashier.CreateInitialisedCashierAgent( ... )
	}

	for i := 0; i < cashierShifts.SecondShiftCount; i++ {
		//newStore.Checkouts[i].SecondShiftCashier = cashier.CreateInitialisedCashierAgent( ... )
	}

	for i := 0; i < floorStaffShifts.FirstShiftCount; i++ {
		newStore.FloorStaffFirstShift = append(newStore.FloorStaffFirstShift, floorStaff.CreateInitialisedFloorStaffAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
			floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound))
	}

	for i := 0; i < floorStaffShifts.SecondShiftCount; i++ {
		newStore.FloorStaffSecondShift = append(newStore.FloorStaffSecondShift, floorStaff.CreateInitialisedFloorStaffAgent(floorStaffAttributeBounds.AmicabilityLowerBound,
			floorStaffAttributeBounds.AmicabilityUpperBound, floorStaffAttributeBounds.CompetanceLowerBound, floorStaffAttributeBounds.CompetanceUpperBound))
	}

	//newStore.FloorManagerFirstShift = floorManager.CreateInitialisedFloorManagerAgent( ... )

	//newStore.FloorManagerSecondShift = floorManager.CreateInitialisedFloorManagerAgent( ... )

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrivalRange := float64(arrivalRates.UpperBound - arrivalRates.LowerBound)
	newStore.baseArrivalRate = math.Round((seed.Float64()*arrivalRange)+float64(arrivalRates.LowerBound)) / float64(60)

	return newStore
}

// PropagateTime : propagates time to agents in store
func (s *StoreAgent) PropagateTime(currentShift int, currentDay int, currentTime float64, externalImpact float64) {
	if math.Mod(currentTime, 60) == 0 {
		// fmt.Println("current arrival rate: ", getRateOfArrival(s.baseArrivalRate, currentDay, currentTime, externalImpact))
		// fmt.Println("First Shift Floor Staff:", s.FloorStaffFirstShift)
		// fmt.Println("Second Shift Floor Staff:", s.FloorStaffSecondShift)
		// for i := 0; i < len(s.Checkouts); i++{
		// 	fmt.Printf("Checkout %d Total: %f\n", i, s.Checkouts[i].TotalMoney)
		// }

	}
	//s.checkNewCustomers(getRateOfArrival(s.baseArrivalRate, currentDay, currentTime, externalImpact))
	//s.checkCustomersReadyToQueue()
	s.propagateConcurrentCheckouts(currentShift) //customers and checkout/cashiers
	s.propagateStoreFloor(currentShift) // customer and floor staff and floor manager
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
	//...
}

func (s *StoreAgent) checkCustomersReadyToQueue() {
	//...
}

func (s *StoreAgent) propagateConcurrentCheckouts(currentShift int) {

	switch true {
	case currentShift == 1:
		for i := 0; i < len(s.Checkouts); i++ {
			s.Checkouts[i].PropagateTime()
		}
		break
	case currentShift == 2:
		for i := 0; i < len(s.Checkouts); i++ {
			s.Checkouts[i].PropagateTime()
		}
		break
	case currentShift == 0:
		break
	}

}

func (s *StoreAgent) propagateStoreFloor(currentShift int) {
	
	switch true {
	case currentShift == 1:
		for i := 0; i < len(s.FloorStaffFirstShift); i++ {
			s.FloorStaffFirstShift[i].PropagateTime()
		}
		break
	case currentShift == 2:
		for i := 0; i < len(s.FloorStaffSecondShift); i++ {
			s.FloorStaffSecondShift[i].PropagateTime()
		}
		break
	case currentShift == 0:
		break
	}

}
