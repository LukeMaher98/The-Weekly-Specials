package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Checkout Variables")
	fmt.Println("---------------------")

	read := true

	checkoutCount := 0
	productsLowerBound := 0
	productsUpperBound := 0
	timeLowerBound := 0.0
	timeUpperBound := 0.0
	arrivalLowerBound := -1
	arrivalUpperBound := -1

	for read {

		if !(checkoutCount >= 1 && checkoutCount <= 8) {
			checkoutCountTemp := 0
			fmt.Print("Number of Checkouts [1-8]> ")
			fmt.Scanln(&checkoutCountTemp)
			fmt.Print("\n")
			if checkoutCountTemp >= 1 && checkoutCountTemp <= 8 {
				checkoutCount = checkoutCountTemp
			}
			continue
		} else if !(productsLowerBound >= 1 && productsLowerBound <= 200) {
			productsLowerBoundTemp := 0
			fmt.Print("Products per Customer Lower Bound [1-200]> ")
			fmt.Scanln(&productsLowerBoundTemp)
			fmt.Print("\n")
			if productsLowerBoundTemp >= 1 && productsLowerBoundTemp <= 200 {
				productsLowerBound = productsLowerBoundTemp
			}
		} else if !(productsUpperBound >= productsLowerBound && productsUpperBound <= 200) {
			productsUpperBoundTemp := 0
			fmt.Printf("Products per Customer Upper Bound [%d-200]> ", productsLowerBound)
			fmt.Scanln(&productsUpperBoundTemp)
			fmt.Print("\n")
			if productsUpperBoundTemp >= productsLowerBound && productsUpperBoundTemp <= 200 {
				productsUpperBound = productsUpperBoundTemp
			}
		} else if !(timeLowerBound >= 0.5 && timeLowerBound <= 6.0) {
			timeLowerBoundTemp := 0.0
			fmt.Print("Time per Product Lower Bound [0.5-6.0]> ")
			fmt.Scanln(&timeLowerBoundTemp)
			fmt.Print("\n")
			if timeLowerBoundTemp >= 0.5 && timeLowerBoundTemp <= 6.0 {
				timeLowerBound = timeLowerBoundTemp
			}
		} else if !(timeUpperBound >= timeLowerBound && timeUpperBound <= 6.0) {
			timeUpperBoundTemp := 0.0
			fmt.Printf("Time per Product Upper Bound [%v-6.0]> ", timeLowerBound)
			fmt.Scanln(&timeUpperBoundTemp)
			fmt.Print("\n")
			if timeUpperBoundTemp >= timeLowerBound && timeUpperBoundTemp <= 6.0 {
				timeUpperBound = timeUpperBoundTemp
			}
		} else if !(arrivalLowerBound >= 0 && arrivalLowerBound <= 60) {
			arrivalLowerBoundTemp := -1
			fmt.Print("Arrival Rate Lower Bound [0-60]> ")
			fmt.Scanln(&arrivalLowerBoundTemp)
			fmt.Print("\n")
			if arrivalLowerBoundTemp >= 0 && arrivalLowerBoundTemp <= 60 {
				arrivalLowerBound = arrivalLowerBoundTemp
			}
		} else if !(arrivalUpperBound >= arrivalLowerBound && arrivalUpperBound <= 60) {
			arrivalUpperBoundTemp := -1
			fmt.Printf("Arrival Rate Upper Bound [%d-60]> ", arrivalLowerBound)
			fmt.Scanln(&arrivalUpperBoundTemp)
			fmt.Print("\n")
			if arrivalUpperBoundTemp >= arrivalLowerBound && arrivalUpperBoundTemp <= 60 {
				arrivalUpperBound = arrivalUpperBoundTemp
			}
		} else {
			fmt.Printf("Checkout Count: [%d]\n", checkoutCount)
			fmt.Printf("Products per Customer Range: [%d-%d]\n", productsLowerBound, productsUpperBound)
			fmt.Printf("Time per Product Range: [%v-%v]\n", timeLowerBound, timeUpperBound)
			fmt.Printf("Arrival Rate Range: [%d-%d]\n", arrivalLowerBound, arrivalUpperBound)
			fmt.Print("\n")
			fmt.Print("C to reset values, X to exit> ")
			selection, _, _ := reader.ReadRune()
			if selection == 'c' || selection == 'C' {
				checkoutCount = 0
				productsLowerBound = 0
				productsUpperBound = 0
				timeLowerBound = 0.0
				timeUpperBound = 0.0
				arrivalLowerBound = -1
				arrivalUpperBound = -1
				fmt.Print("\n")
			} else if selection == 'x' || selection == 'X' {
				read = false
				fmt.Print("\n")
			}
		}

	}

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
