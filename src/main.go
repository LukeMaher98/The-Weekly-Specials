package main

import (
	"src/scenario"
	"time"
)

func main() {
	scenario := scenario.CreateInitialisedScenarioAgent()

	start := time.Now()
	currentTime := scenario.OpeningTime * 60

	scenario.Activate()

	for scenario.ScenarioActive == true {
		elapsed := float64((time.Since(start)/time.Microsecond)/10000) + (scenario.OpeningTime * 60)
		if elapsed != currentTime {
			currentTime = scenario.PropagateTime(currentTime)
		}
	}

	scenario.Store.PrintResults()
}
