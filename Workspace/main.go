package main

import (
	"fmt"
	"agents"
)

func main () {
	var floorStaff = agents.ConstructorFloorStaff(10,0.34,0.56)

	floorStaff.PrintStaff()

	fmt.Println("Occupied Status:", floorStaff.GetOccupied())

	floorStaff.SetOccupied(true)

	fmt.Println("Occupied Status:", floorStaff.GetOccupied())


	var checkout = agents.ConstructorCheckout(true, true)

	fmt.Println("Adult Checkout:", checkout.IsAdultCheckout())

	fmt.Println("Self Checkout:", checkout.IsSelfCheckout())

	fmt.Println("Total Money In Checkout:", checkout.GetMoney())
	
	checkout.AddMoney(0.50)

	fmt.Println("Total Money In Checkout:", checkout.GetMoney())
}