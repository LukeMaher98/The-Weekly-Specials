package main

import (
	"fmt"
	"src/cashier"
)

func main() {
	/*scenario := scenario.CreateInitialisedScenarioAgent()

	start := time.Now()
	currentTime := 0.0
	scenario.ScenarioActive = true

	for scenario.ScenarioActive == true {
		elapsed := float64((time.Since(start) / time.Microsecond) / 10000)
		if elapsed != currentTime {
			fmt.Println(elapsed)
			if int(elapsed/1440) < scenario.ScenarioDuration {
				currentTime = elapsed
				scenario.PropagateTime(currentTime)
			} else {
				scenario.ScenarioActive = false
			}
		}
	}

	scenario.PrintResults()*/

	cashier := cashier.NewCashier()
	fmt.Println(cashier)
	fmt.Println(cashier.ActualTimeCard)
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
