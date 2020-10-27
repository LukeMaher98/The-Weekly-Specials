package main

import (
	"fmt"
	"agents"
)

func main () {
	var floorStaff = agents.ConstructorFloorStaff(10,0.34,0.56)
	agents.PrintStaff(floorStaff)

	var checkout = agents.ConstructorCheckout(true, true)
	fmt.Println(checkout)

	checkout = agents.AddMoney(checkout, 0.50)

	fmt.Println(checkout)
}