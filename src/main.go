package main

import (
	"src/scenario"
	"time"
)

func main() {
	scenario := scenario.CreateInitialisedScenarioAgent()

	start := time.Now()
	currentTime := 0.0

	scenario.Activate()

	for scenario.ScenarioActive == true {
		elapsed := float64((time.Since(start) / time.Microsecond) / 10000)
		if elapsed != currentTime {
			currentTime = scenario.PropagateTime(currentTime)
		}
	}

	scenario.PrintResults()
}
