package scenario

import (
	"fmt"
)

type ScenarioAgent struct {
	startingDay           int
	startingTime          float64
	openingTime           float64
	closingTime           float64
	currentTime           float64
	checkoutCount         int
	productsLowerBound    int
	productsUpperBound    int
	itemTimeLowerBound    float64
	itemTimeUpperBound    float64
	arrivalLowerBound     int
	arrivalUpperBound     int
	firstShiftFloorStaff  int
	secondShiftFloorStaff int
	ScenarioDuration      int
	ScenarioActive        bool
	//store          storeAgent
}

// CreateScenarioAgent : creates 'empty' Scenario for initialisation
func CreateScenarioAgent() ScenarioAgent {
	return ScenarioAgent{-1, -1.0, -1.0, -1.0, -1.0, 0, 0, 0, 0.0, 0.0, -1, -1, -1, -1, 0, false}
}

// CreateInitialisedScenarioAgent : creates populated Scenario for CLI
func CreateInitialisedScenarioAgent() ScenarioAgent {
	newScenario := CreateScenarioAgent()

	fmt.Println("Scenario Variables")
	fmt.Println("---------------------")

	read := true

	for read {
		if !(newScenario.ScenarioDuration > 0) {
			ScenarioDurationTemp := 0
			fmt.Print("Duration of Simulation in Days [1+]> ")
			fmt.Scanln(&ScenarioDurationTemp)
			fmt.Print("\n")
			if ScenarioDurationTemp > 0 {
				newScenario.ScenarioDuration = ScenarioDurationTemp
			}
			continue
		} else if !(newScenario.startingDay >= 0 && newScenario.startingDay <= 6) {
			startingDayTemp := -1
			fmt.Print("Starting Day of Week for Simulation [0-6]> ")
			fmt.Scanln(&startingDayTemp)
			fmt.Print("\n")
			if startingDayTemp >= 0 && startingDayTemp <= 6 {
				newScenario.startingDay = startingDayTemp
			}
			continue
		} else if !(newScenario.startingTime >= 0.0 && newScenario.startingTime <= 24.0) {
			startingTimeTemp := -1.0
			fmt.Print("Starting Time of Day for Simulation [0.0-24.0]> ")
			fmt.Scanln(&startingTimeTemp)
			fmt.Print("\n")
			if startingTimeTemp >= 0.0 && startingTimeTemp <= 24.0 {
				newScenario.startingTime = startingTimeTemp
			}
			continue
		} else if !(newScenario.openingTime >= 0.0 && newScenario.openingTime <= 24.0) {
			openingTimeTemp := -1.0
			fmt.Print("Opening Time [0.0-24.0]> ")
			fmt.Scanln(&openingTimeTemp)
			fmt.Print("\n")
			if openingTimeTemp >= 0.0 && openingTimeTemp <= 24.0 {
				newScenario.openingTime = openingTimeTemp
			}
			continue
		} else if !(newScenario.closingTime >= newScenario.itemTimeLowerBound && newScenario.closingTime <= 24.0) {
			closingTimeTemp := -1.0
			fmt.Printf("Closing Time [%v-24.0]> ", newScenario.openingTime)
			fmt.Scanln(&closingTimeTemp)
			fmt.Print("\n")
			if closingTimeTemp >= newScenario.openingTime && closingTimeTemp <= 24.0 {
				newScenario.closingTime = closingTimeTemp
			}
			continue
		} else if !(newScenario.checkoutCount >= 1 && newScenario.checkoutCount <= 8) {
			checkoutCountTemp := 0
			fmt.Print("Number of Checkouts [1-8]> ")
			fmt.Scanln(&checkoutCountTemp)
			fmt.Print("\n")
			if checkoutCountTemp >= 1 && checkoutCountTemp <= 8 {
				newScenario.checkoutCount = checkoutCountTemp
			}
			continue
		} else if !(newScenario.productsLowerBound >= 1 && newScenario.productsLowerBound <= 200) {
			productsLowerBoundTemp := 0
			fmt.Print("Products per Customer Lower Bound [1-200]> ")
			fmt.Scanln(&productsLowerBoundTemp)
			fmt.Print("\n")
			if productsLowerBoundTemp >= 1 && productsLowerBoundTemp <= 200 {
				newScenario.productsLowerBound = productsLowerBoundTemp
			}
			continue
		} else if !(newScenario.productsUpperBound >= newScenario.productsLowerBound && newScenario.productsUpperBound <= 200) {
			productsUpperBoundTemp := 0
			fmt.Printf("Products per Customer Upper Bound [%v-200]> ", newScenario.productsLowerBound)
			fmt.Scanln(&productsUpperBoundTemp)
			fmt.Print("\n")
			if productsUpperBoundTemp >= newScenario.productsLowerBound && productsUpperBoundTemp <= 200 {
				newScenario.productsUpperBound = productsUpperBoundTemp
			}
			continue
		} else if !(newScenario.itemTimeLowerBound >= 0.5 && newScenario.itemTimeLowerBound <= 6.0) {
			itemTimeLowerBoundTemp := 0.0
			fmt.Print("Time per Product Lower Bound [0.5-6.0]> ")
			fmt.Scanln(&itemTimeLowerBoundTemp)
			fmt.Print("\n")
			if itemTimeLowerBoundTemp >= 0.5 && itemTimeLowerBoundTemp <= 6.0 {
				newScenario.itemTimeLowerBound = itemTimeLowerBoundTemp
			}
			continue
		} else if !(newScenario.itemTimeUpperBound >= newScenario.itemTimeLowerBound && newScenario.itemTimeUpperBound <= 6.0) {
			itemTimeUpperBoundTemp := 0.0
			fmt.Printf("Time per Product Upper Bound [%v-6.0]> ", newScenario.itemTimeLowerBound)
			fmt.Scanln(&itemTimeUpperBoundTemp)
			fmt.Print("\n")
			if itemTimeUpperBoundTemp >= newScenario.itemTimeLowerBound && itemTimeUpperBoundTemp <= 6.0 {
				newScenario.itemTimeUpperBound = itemTimeUpperBoundTemp
			}
			continue
		} else if !(newScenario.arrivalLowerBound >= 0 && newScenario.arrivalLowerBound <= 60) {
			arrivalLowerBoundTemp := -1
			fmt.Print("Arrival Rate Lower Bound [0-60]> ")
			fmt.Scanln(&arrivalLowerBoundTemp)
			fmt.Print("\n")
			if arrivalLowerBoundTemp >= 0 && arrivalLowerBoundTemp <= 60 {
				newScenario.arrivalLowerBound = arrivalLowerBoundTemp
			}
			continue
		} else if !(newScenario.arrivalUpperBound >= newScenario.arrivalLowerBound && newScenario.arrivalUpperBound <= 60) {
			arrivalUpperBoundTemp := -1
			fmt.Printf("Arrival Rate Upper Bound [%d-60]> ", newScenario.arrivalLowerBound)
			fmt.Scanln(&arrivalUpperBoundTemp)
			fmt.Print("\n")
			if arrivalUpperBoundTemp >= newScenario.arrivalLowerBound && arrivalUpperBoundTemp <= 60 {
				newScenario.arrivalUpperBound = arrivalUpperBoundTemp
			}
		} else if !(newScenario.firstShiftFloorStaff >= 0) {
			firstShiftFloorStaffTemp := -1
			fmt.Print("Number of Floor Staff [First Shift] [0+]> ")
			fmt.Scanln(&firstShiftFloorStaffTemp)
			fmt.Print("\n")
			if firstShiftFloorStaffTemp >= 0 {
				newScenario.firstShiftFloorStaff = firstShiftFloorStaffTemp
			}
			continue
		} else if !(newScenario.secondShiftFloorStaff >= 0) {
			secondShiftFloorStaffTemp := -1
			fmt.Print("Number of Floor Staff [Second Shift] [0+]> ")
			fmt.Scanln(&secondShiftFloorStaffTemp)
			fmt.Print("\n")
			if secondShiftFloorStaffTemp >= 0 {
				newScenario.secondShiftFloorStaff = secondShiftFloorStaffTemp
			}
			continue
		} else {
			read = false
		}
	}

	return newScenario
}

// PropagateTime : propagates time through simulation
func (s ScenarioAgent) PropagateTime(elapsedTime float64) {
	//s.store.propagateTime(elapsedTime)
}

// PrintResults : prints results of simulation
func (s ScenarioAgent) PrintResults() {
	fmt.Println("Scenario Results:")
	fmt.Println("------------------")
	//...
}

/*func philos(id int, left, right chan bool, wg *sync.WaitGroup) {
	fmt.Printf("Philosopher # %d wants to eat\n", id)
	<-left
	<-right
	left <- true
	right <- true
	fmt.Printf("Philosopher # %d finished eating\n", id)
	wg.Done()
}
func main() {
	const numPhilos = 5
	var forks [numPhilos]chan bool
	for i := 0; i < numPhilos; i++ {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}
	var wg sync.WaitGroup
	for i := 0; i < numPhilos; i++ {
		wg.Add(1)
		go philos(i, forks[(i-1+numPhilos)%numPhilos], forks[(i+numPhilos)%numPhilos], &wg)
	}
	wg.Wait()
	fmt.Println("Everybody finished eating")
}*/
