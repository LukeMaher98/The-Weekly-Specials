package scenario

import (
	"fmt"
	"math"
	"src/constants"
	"src/store"
	"strings"
)

type ScenarioAgent struct {
	// User Defined Variables
	ScenarioDuration            int
	ScenarioActive              bool
	startingDay                 int
	startingTime                float64
	openingTime                 float64
	closingTime                 float64
	weatherConditions           float64
	socialConditions            float64
	covidLockdownLevel          int
	checkoutCount               int
	itemLimitBounds             constants.StoreAttributeBoundsInt
	itemTimeBounds              constants.StoreAttributeBoundsFloat
	arrivalBounds               constants.StoreAttributeBoundsInt
	floorStaffShifts            constants.StoreShifts
	cashierShifts               constants.StoreShifts
	floorStaffAttributeBounds   constants.StaffAttributeBounds
	cashierAttributeBounds      constants.StaffAttributeBounds
	floorManagerAttributeBounds constants.StaffAttributeBounds

	store store.StoreAgent

	// Dynamically Defined variables
	currentDay   int
	currentTime  float64
	currentShift int
}

// CreateScenarioAgent : creates 'empty' Scenario for initialisation
func CreateScenarioAgent() ScenarioAgent {
	return ScenarioAgent{-1, false, -1, -1.0, -1.0, -1.0, -2.0, -2.0, -1, 0,
		constants.StoreAttributeBoundsInt{UpperBound: 0, LowerBound: 0},
		constants.StoreAttributeBoundsFloat{UpperBound: 0.0, LowerBound: 0.0},
		constants.StoreAttributeBoundsInt{UpperBound: -1, LowerBound: -1},
		constants.StoreShifts{FirstShiftCount: -1, SecondShiftCount: -1},
		constants.StoreShifts{FirstShiftCount: -1, SecondShiftCount: -1},
		constants.StaffAttributeBounds{AmicabilityUpperBound: -1.0, AmicabilityLowerBound: -1.0, CompetanceUpperBound: -1.0, CompetanceLowerBound: -1.0},
		constants.StaffAttributeBounds{AmicabilityUpperBound: -1.0, AmicabilityLowerBound: -1.0, CompetanceUpperBound: -1.0, CompetanceLowerBound: -1.0},
		constants.StaffAttributeBounds{AmicabilityUpperBound: -1.0, AmicabilityLowerBound: -1.0, CompetanceUpperBound: -1.0, CompetanceLowerBound: -1.0},
		store.CreateStoreAgent(),
		-1, -1.0, -1}
}

// CreateInitialisedScenarioAgent : creates populated Scenario for CLI
func CreateInitialisedScenarioAgent() ScenarioAgent {
	newScenario := CreateScenarioAgent()

	//Uncomment to check items
	//item.PrintItems()

	fmt.Println("Scenario Variables")
	fmt.Println("---------------------")

	read := true
	var defineEmployees bool
	employeeDefinitionCheck := false

	for read {
		if !(newScenario.ScenarioDuration > 0) {
			ScenarioDurationTemp := 0
			fmt.Print("Duration of Simulation in Days [1+]> ")
			fmt.Scanln(&ScenarioDurationTemp)
			if ScenarioDurationTemp > 0 {
				newScenario.ScenarioDuration = ScenarioDurationTemp
			}
		} else if !(newScenario.startingDay >= 0 && newScenario.startingDay <= 6) {
			startingDayTemp := -1
			fmt.Print("Starting Day of Week for Simulation [0-6]> ")
			fmt.Scanln(&startingDayTemp)
			if startingDayTemp > -1 && startingDayTemp < 7 {
				newScenario.startingDay = startingDayTemp
			}
		} else if !(newScenario.startingTime >= 0.0 && newScenario.startingTime <= 24.0) {
			startingTimeTemp := -1.0
			fmt.Print("Starting Time of Day for Simulation [0.0-24.0]> ")
			fmt.Scanln(&startingTimeTemp)
			if startingTimeTemp >= 0.0 && startingTimeTemp <= 24.0 {
				newScenario.startingTime = startingTimeTemp
			}
		} else if !(newScenario.openingTime >= 0.0 && newScenario.openingTime <= 24.0) {
			openingTimeTemp := -1.0
			fmt.Print("Opening Time [0.0-24.0]> ")
			fmt.Scanln(&openingTimeTemp)
			if openingTimeTemp >= 0.0 && openingTimeTemp <= 24.0 {
				newScenario.openingTime = openingTimeTemp
			}
		} else if !(newScenario.closingTime >= newScenario.openingTime && newScenario.closingTime <= 24.0) {
			closingTimeTemp := -1.0
			fmt.Printf("Closing Time [%v-24.0]> ", newScenario.openingTime)
			fmt.Scanln(&closingTimeTemp)
			if closingTimeTemp >= newScenario.openingTime && closingTimeTemp <= 24.0 {
				newScenario.closingTime = closingTimeTemp
			}
		} else if !(newScenario.weatherConditions >= -1.0 && newScenario.weatherConditions <= 1.0) {
			weatherConditionsTemp := -2.0
			fmt.Print("Negative or positive impact of weather conditions [-1.0-1.0]> ")
			fmt.Scanln(&weatherConditionsTemp)
			if weatherConditionsTemp >= -1.0 && weatherConditionsTemp <= 1.0 {
				newScenario.weatherConditions = weatherConditionsTemp
			}
		} else if !(newScenario.socialConditions >= -1.0 && newScenario.socialConditions <= 1.0) {
			socialConditionsTemp := -2.0
			fmt.Print("Negative or positive impact of social conditions (e.g concert, match, local tragedy etc) [-1.0-1.0]> ")
			fmt.Scanln(&socialConditionsTemp)
			if socialConditionsTemp >= -1.0 && socialConditionsTemp <= 1.0 {
				newScenario.socialConditions = socialConditionsTemp
			}
		} else if !(newScenario.covidLockdownLevel >= 0 && newScenario.covidLockdownLevel <= 5) {
			covidLockdownLevelTemp := -1
			fmt.Print("Level of Covid-19 restrictions in place [0-5]> ")
			fmt.Scanln(&covidLockdownLevelTemp)
			if covidLockdownLevelTemp >= 0 && covidLockdownLevelTemp <= 5 {
				newScenario.covidLockdownLevel = covidLockdownLevelTemp
			}
		} else if !(newScenario.checkoutCount >= 1 && newScenario.checkoutCount <= 8) {
			checkoutCountTemp := 0
			fmt.Print("Number of Checkouts [1-8]> ")
			fmt.Scanln(&checkoutCountTemp)
			if checkoutCountTemp >= 1 && checkoutCountTemp <= 8 {
				newScenario.checkoutCount = checkoutCountTemp
			}
		} else if !(newScenario.itemLimitBounds.LowerBound >= 1 && newScenario.itemLimitBounds.LowerBound <= 200) {
			itemLimitLowerBoundTemp := 0
			fmt.Print("Items per Customer Lower Bound [1-200]> ")
			fmt.Scanln(&itemLimitLowerBoundTemp)
			if itemLimitLowerBoundTemp >= 1 && itemLimitLowerBoundTemp <= 200 {
				newScenario.itemLimitBounds.LowerBound = itemLimitLowerBoundTemp
			}
		} else if !(newScenario.itemLimitBounds.UpperBound >= newScenario.itemLimitBounds.LowerBound && newScenario.itemLimitBounds.UpperBound <= 200) {
			itemLimitUpperBoundTemp := 0
			fmt.Printf("Items per Customer Upper Bound [%v-200]> ", newScenario.itemLimitBounds.LowerBound)
			fmt.Scanln(&itemLimitUpperBoundTemp)
			if itemLimitUpperBoundTemp >= newScenario.itemLimitBounds.LowerBound && itemLimitUpperBoundTemp <= 200 {
				newScenario.itemLimitBounds.UpperBound = itemLimitUpperBoundTemp
			}
		} else if !(newScenario.itemTimeBounds.LowerBound >= 0.5 && newScenario.itemTimeBounds.LowerBound <= 6.0) {
			itemTimeLowerBoundTemp := 0.0
			fmt.Print("Time per Product Lower Bound [0.5-6.0]> ")
			fmt.Scanln(&itemTimeLowerBoundTemp)
			if itemTimeLowerBoundTemp >= 0.5 && itemTimeLowerBoundTemp <= 6.0 {
				newScenario.itemTimeBounds.LowerBound = itemTimeLowerBoundTemp
			}
		} else if !(newScenario.itemTimeBounds.UpperBound >= newScenario.itemTimeBounds.LowerBound && newScenario.itemTimeBounds.UpperBound <= 6.0) {
			itemTimeUpperBoundTemp := 0.0
			fmt.Printf("Time per Product Upper Bound [%v-6.0]> ", newScenario.itemTimeBounds.LowerBound)
			fmt.Scanln(&itemTimeUpperBoundTemp)
			if itemTimeUpperBoundTemp >= newScenario.itemTimeBounds.LowerBound && itemTimeUpperBoundTemp <= 6.0 {
				newScenario.itemTimeBounds.UpperBound = itemTimeUpperBoundTemp
			}
		} else if !(newScenario.arrivalBounds.LowerBound >= 0 && newScenario.arrivalBounds.LowerBound <= 60) {
			arrivalLowerBoundTemp := -1
			fmt.Print("Arrival Rate Lower Bound [0-60]> ")
			fmt.Scanln(&arrivalLowerBoundTemp)
			if arrivalLowerBoundTemp >= 0 && arrivalLowerBoundTemp <= 60 {
				newScenario.arrivalBounds.LowerBound = arrivalLowerBoundTemp
			}
		} else if !(newScenario.arrivalBounds.UpperBound >= newScenario.arrivalBounds.LowerBound && newScenario.arrivalBounds.UpperBound <= 60) {
			arrivalUpperBoundTemp := -1
			fmt.Printf("Arrival Rate Upper Bound [%d-60]> ", newScenario.arrivalBounds.LowerBound)
			fmt.Scanln(&arrivalUpperBoundTemp)
			if arrivalUpperBoundTemp >= newScenario.arrivalBounds.LowerBound && arrivalUpperBoundTemp <= 60 {
				newScenario.arrivalBounds.UpperBound = arrivalUpperBoundTemp
			}
		} else if !(newScenario.floorStaffShifts.FirstShiftCount >= 0) {
			firstShiftFloorStaffTemp := -1
			fmt.Print("Number of Floor Staff [First Shift] [0+]> ")
			fmt.Scanln(&firstShiftFloorStaffTemp)
			if firstShiftFloorStaffTemp >= 0 {
				newScenario.floorStaffShifts.FirstShiftCount = firstShiftFloorStaffTemp
			}
		} else if !(newScenario.floorStaffShifts.SecondShiftCount >= 0) {
			secondShiftFloorStaffTemp := -1
			fmt.Print("Number of Floor Staff [Second Shift] [0+]> ")
			fmt.Scanln(&secondShiftFloorStaffTemp)
			if secondShiftFloorStaffTemp >= 0 {
				newScenario.floorStaffShifts.SecondShiftCount = secondShiftFloorStaffTemp
			}
		} else if !(newScenario.cashierShifts.FirstShiftCount >= 0) {
			firstShiftCashiersTemp := -1
			fmt.Print("Number of Cashiers [First Shift] [0+]> ")
			fmt.Scanln(&firstShiftCashiersTemp)
			if firstShiftCashiersTemp >= 0 {
				newScenario.cashierShifts.FirstShiftCount = firstShiftCashiersTemp
			}
		} else if !(newScenario.cashierShifts.SecondShiftCount >= 0) {
			secondShiftCashiersTemp := -1
			fmt.Print("Number of Cashiers [Second Shift] [0+]> ")
			fmt.Scanln(&secondShiftCashiersTemp)
			if secondShiftCashiersTemp >= 0 {
				newScenario.cashierShifts.SecondShiftCount = secondShiftCashiersTemp
			}
		} else if employeeDefinitionCheck == false {
			var defineEmployeesTemp string
			fmt.Print("Define employee characteristics? [Y/N]> ")
			fmt.Scanln(&defineEmployeesTemp)
			if strings.ContainsRune(defineEmployeesTemp, 'Y') && len(defineEmployeesTemp) == 1 {
				defineEmployees = true
				employeeDefinitionCheck = true
			} else if strings.ContainsRune(defineEmployeesTemp, 'N') && len(defineEmployeesTemp) == 1 {
				defineEmployees = false
				employeeDefinitionCheck = true
			}
		} else if defineEmployees == true {
			if newScenario.floorStaffAttributeBounds.AmicabilityLowerBound < 0.0 || newScenario.floorStaffAttributeBounds.AmicabilityLowerBound > 1.0 {
				fsAmicabilityLowerBoundTemp := 0.0
				fmt.Print("Floor Staff Amicability Lower Bound [0.0-1.0]> ")
				fmt.Scanln(&fsAmicabilityLowerBoundTemp)
				if fsAmicabilityLowerBoundTemp >= 0.0 && fsAmicabilityLowerBoundTemp <= 1.0 {
					newScenario.floorStaffAttributeBounds.AmicabilityLowerBound = fsAmicabilityLowerBoundTemp
				}
			} else if newScenario.floorStaffAttributeBounds.AmicabilityUpperBound < newScenario.floorStaffAttributeBounds.AmicabilityLowerBound ||
				newScenario.floorStaffAttributeBounds.AmicabilityUpperBound > 1.0 {
				fsAmicabilityUpperBoundTemp := 0.0
				fmt.Printf("Floor Staff Amicability Upper Bound [%v-1.0]> ", newScenario.floorStaffAttributeBounds.AmicabilityLowerBound)
				fmt.Scanln(&fsAmicabilityUpperBoundTemp)
				if fsAmicabilityUpperBoundTemp >= newScenario.floorStaffAttributeBounds.AmicabilityLowerBound && fsAmicabilityUpperBoundTemp <= 1.0 {
					newScenario.floorStaffAttributeBounds.AmicabilityUpperBound = fsAmicabilityUpperBoundTemp
				}
			} else if newScenario.floorStaffAttributeBounds.CompetanceLowerBound < 0.0 || newScenario.floorStaffAttributeBounds.CompetanceLowerBound > 1.0 {
				fsCompetanceLowerBoundTemp := 0.0
				fmt.Print("Floor Staff Competance Lower Bound [0.0-1.0]> ")
				fmt.Scanln(&fsCompetanceLowerBoundTemp)
				if fsCompetanceLowerBoundTemp >= 0.0 && fsCompetanceLowerBoundTemp <= 1.0 {
					newScenario.floorStaffAttributeBounds.CompetanceLowerBound = fsCompetanceLowerBoundTemp
				}
			} else if newScenario.floorStaffAttributeBounds.CompetanceUpperBound < newScenario.floorStaffAttributeBounds.CompetanceLowerBound ||
				newScenario.floorStaffAttributeBounds.CompetanceUpperBound > 1.0 {
				fsCompetanceUpperBoundTemp := 0.0
				fmt.Printf("Floor Staff Amicability Upper Bound [%v-1.0]> ", newScenario.floorStaffAttributeBounds.CompetanceLowerBound)
				fmt.Scanln(&fsCompetanceUpperBoundTemp)
				if fsCompetanceUpperBoundTemp >= newScenario.floorStaffAttributeBounds.CompetanceLowerBound && fsCompetanceUpperBoundTemp <= 1.0 {
					newScenario.floorStaffAttributeBounds.CompetanceUpperBound = fsCompetanceUpperBoundTemp
				}
			} else if newScenario.cashierAttributeBounds.AmicabilityLowerBound < 0.0 || newScenario.cashierAttributeBounds.AmicabilityLowerBound > 1.0 {
				cAmicabilityLowerBoundTemp := 0.0
				fmt.Print("Cashier Amicability Lower Bound [0.0-1.0]> ")
				fmt.Scanln(&cAmicabilityLowerBoundTemp)
				if cAmicabilityLowerBoundTemp >= 0.0 && cAmicabilityLowerBoundTemp <= 1.0 {
					newScenario.cashierAttributeBounds.AmicabilityLowerBound = cAmicabilityLowerBoundTemp
				}
			} else if newScenario.cashierAttributeBounds.AmicabilityUpperBound < newScenario.cashierAttributeBounds.AmicabilityLowerBound ||
				newScenario.cashierAttributeBounds.AmicabilityUpperBound > 1.0 {
				cAmicabilityUpperBoundTemp := 0.0
				fmt.Printf("Cashier Amicability Upper Bound [%v-1.0]> ", newScenario.cashierAttributeBounds.AmicabilityLowerBound)
				fmt.Scanln(&cAmicabilityUpperBoundTemp)
				if cAmicabilityUpperBoundTemp >= newScenario.cashierAttributeBounds.AmicabilityLowerBound && cAmicabilityUpperBoundTemp <= 1.0 {
					newScenario.cashierAttributeBounds.AmicabilityUpperBound = cAmicabilityUpperBoundTemp
				}
			} else if newScenario.cashierAttributeBounds.CompetanceLowerBound < 0.0 || newScenario.cashierAttributeBounds.CompetanceLowerBound > 1.0 {
				cCompetanceLowerBoundTemp := 0.0
				fmt.Print("Cashier Competance Lower Bound [0.0-1.0]> ")
				fmt.Scanln(&cCompetanceLowerBoundTemp)
				if cCompetanceLowerBoundTemp >= 0.0 && cCompetanceLowerBoundTemp <= 1.0 {
					newScenario.cashierAttributeBounds.CompetanceLowerBound = cCompetanceLowerBoundTemp
				}
			} else if newScenario.cashierAttributeBounds.CompetanceUpperBound < newScenario.cashierAttributeBounds.CompetanceLowerBound ||
				newScenario.cashierAttributeBounds.CompetanceUpperBound > 1.0 {
				cCompetanceUpperBoundTemp := 0.0
				fmt.Printf("Cashier Amicability Upper Bound [%v-1.0]> ", newScenario.cashierAttributeBounds.CompetanceLowerBound)
				fmt.Scanln(&cCompetanceUpperBoundTemp)
				if cCompetanceUpperBoundTemp >= newScenario.cashierAttributeBounds.CompetanceLowerBound && cCompetanceUpperBoundTemp <= 1.0 {
					newScenario.cashierAttributeBounds.CompetanceUpperBound = cCompetanceUpperBoundTemp
				}
			} else if newScenario.floorManagerAttributeBounds.AmicabilityLowerBound < 0.0 || newScenario.floorManagerAttributeBounds.AmicabilityLowerBound > 1.0 {
				fmAmicabilityLowerBoundTemp := 0.0
				fmt.Print("Floor Manager Amicability Lower Bound [0.0-1.0]> ")
				fmt.Scanln(&fmAmicabilityLowerBoundTemp)
				if fmAmicabilityLowerBoundTemp >= 0.0 && fmAmicabilityLowerBoundTemp <= 1.0 {
					newScenario.floorManagerAttributeBounds.AmicabilityLowerBound = fmAmicabilityLowerBoundTemp
				}
			} else if newScenario.floorManagerAttributeBounds.AmicabilityUpperBound < newScenario.floorManagerAttributeBounds.AmicabilityLowerBound ||
				newScenario.floorManagerAttributeBounds.AmicabilityUpperBound > 1.0 {
				fmAmicabilityUpperBoundTemp := 0.0
				fmt.Printf("Floor Manager Amicability Upper Bound [%v-1.0]> ", newScenario.floorManagerAttributeBounds.AmicabilityLowerBound)
				fmt.Scanln(&fmAmicabilityUpperBoundTemp)
				if fmAmicabilityUpperBoundTemp >= newScenario.floorManagerAttributeBounds.AmicabilityLowerBound && fmAmicabilityUpperBoundTemp <= 1.0 {
					newScenario.floorManagerAttributeBounds.AmicabilityUpperBound = fmAmicabilityUpperBoundTemp
				}
			} else if newScenario.floorManagerAttributeBounds.CompetanceLowerBound < 0.0 || newScenario.floorManagerAttributeBounds.CompetanceLowerBound > 1.0 {
				fmCompetanceLowerBoundTemp := 0.0
				fmt.Print("Floor Manager Competance Lower Bound [0.0-1.0]> ")
				fmt.Scanln(&fmCompetanceLowerBoundTemp)
				if fmCompetanceLowerBoundTemp >= 0.0 && fmCompetanceLowerBoundTemp <= 1.0 {
					newScenario.floorManagerAttributeBounds.CompetanceLowerBound = fmCompetanceLowerBoundTemp
				}
			} else if newScenario.floorManagerAttributeBounds.CompetanceUpperBound < newScenario.floorManagerAttributeBounds.CompetanceLowerBound ||
				newScenario.floorManagerAttributeBounds.CompetanceUpperBound > 1.0 {
				fmCompetanceUpperBoundTemp := 0.0
				fmt.Printf("Floor Manager Amicability Upper Bound [%v-1.0]> ", newScenario.floorManagerAttributeBounds.CompetanceLowerBound)
				fmt.Scanln(&fmCompetanceUpperBoundTemp)
				if fmCompetanceUpperBoundTemp >= newScenario.floorManagerAttributeBounds.CompetanceLowerBound && fmCompetanceUpperBoundTemp <= 1.0 {
					newScenario.floorManagerAttributeBounds.CompetanceUpperBound = fmCompetanceUpperBoundTemp
				}
			} else {
				read = false
			}
		} else if defineEmployees == false {
			newScenario.floorStaffAttributeBounds = constants.StaffAttributeBounds{AmicabilityLowerBound: 0.25, AmicabilityUpperBound: 0.75, CompetanceLowerBound: 0.25, CompetanceUpperBound: 0.75}
			newScenario.cashierAttributeBounds = constants.StaffAttributeBounds{AmicabilityLowerBound: 0.25, AmicabilityUpperBound: 0.75,
				CompetanceLowerBound: 0.25, CompetanceUpperBound: 0.75}
			newScenario.floorManagerAttributeBounds = constants.StaffAttributeBounds{AmicabilityLowerBound: 0.25, AmicabilityUpperBound: 0.75, CompetanceLowerBound: 0.25, CompetanceUpperBound: 0.75}
			read = false
		}
	}

	newScenario.store = store.CreateInitialisedStoreAgent(
		newScenario.arrivalBounds,
		newScenario.itemLimitBounds,
		newScenario.itemTimeBounds,
		newScenario.checkoutCount,
		newScenario.cashierShifts,
		newScenario.floorStaffShifts,
		newScenario.floorStaffAttributeBounds,
		newScenario.cashierAttributeBounds,
		newScenario.floorManagerAttributeBounds,
		newScenario.itemLimitBounds,
	)

	return newScenario
}

// PropagateTime : propagates time through simulation
func (s *ScenarioAgent) PropagateTime(elapsed float64) float64 {
	closedTime := 1440 - (s.closingTime*60 - s.openingTime*60)
	elapsedTime := elapsed + 1

	s.currentDay = ((int(elapsedTime+s.startingTime*60) / 1440) + s.startingDay) % 7
	s.currentTime = math.Mod((elapsedTime + s.startingTime*60), 1440.0)
	if s.currentTime >= s.openingTime*60 && s.currentTime <= s.closingTime*60 {
		if s.currentTime < (s.openingTime*60 + ((s.closingTime - s.openingTime) * 30)) {
			s.store.PropagateTime(0, s.currentDay, s.currentTime, s.getEnvironmentalImpactOnArrival())
			fmt.Println("Day of Week: ", s.currentDay, "Time of Day ", s.currentTime, "Current shift: 0")
		} else {
			s.store.PropagateTime(1, s.currentDay, s.currentTime, s.getEnvironmentalImpactOnArrival())
			fmt.Println("Day of Week: ", s.currentDay, "Time of Day ", s.currentTime, "Current shift: 1")
		}
	} else {
		elapsedTime += closedTime - 1
	}

	if int(elapsedTime/1440) >= s.ScenarioDuration {
		s.ScenarioActive = false
	}

	return elapsedTime
}

// PrintResults : prints results of simulation
func (s *ScenarioAgent) PrintResults() {
	fmt.Println("Scenario Results:")
	fmt.Println("------------------")
	//...
}

func (s *ScenarioAgent) getEnvironmentalImpactOnArrival() float64 {
	rawValue := (s.weatherConditions + s.socialConditions) / 8
	rawValue -= 0.05 * float64(s.covidLockdownLevel)
	return rawValue
}

// Activate: begins simulation
func (s *ScenarioAgent) Activate() {
	s.ScenarioActive = true
}
