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
			currentTime = elapsed
			scenario.PropagateTime(currentTime)
		}
	}

	scenario.PrintResults()
}

// TO BE ADDRESSED

// - 'Actual values', i.e. dynamically refined equivalents of base values, all comes as the output of functions and therefore do not require a dedicated
//   attribute within the struct of their respective agent

// - The number of checkouts is permanent over the course of the simulation, while the number of cashiers may vary. Therefore checkouts agenst should
//   have their active cashier agent as an attribute, and not the inverse to allow for seemly time propagation

// - All human-representative agent attributes relating to the effieciency with which a staff member does there job should be identified as 'Competance'.
// Likewise all those that relate to how smoothly an a human-representative agent interacts with other such agents should be referred to as 'Amicability'.

// - A time propogation function must be implemented for all agents, with internal agent data persisting through time iterations. At such a point, a sequential
// implementation of the core simulation loop should be possible, following on from which the issue of comcurrency can be addressed.
