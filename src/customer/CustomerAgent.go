package main

import (
	"fmt"
	item "src/customer/item"
)

type CustomerAgent struct {
	items            []item.ItemAgent
	impairmentFactor float64
	/* replaceItem             float64
	couponItem              float64
	withChildren            bool
	loyaltyCard             bool
	baseAmicability         float64
	customerAmicability     float64
	preferredPayment        float64
	avaliablePayment        []int
	baggintTimeSelfCheckout float64
	emergencyLeave          float64
	switchLine              float64 */
}

func newCustomer(items []item.ItemAgent, impairmentFactor float64) *CustomerAgent {

	ca := CustomerAgent{impairmentFactor: impairmentFactor}
	return &ca
}

func (customer *CustomerAgent) AddItem(item item.ItemAgent) {
	customer.items = append(customer.items, item)
}

func main() {
	item1 := *item.NewItem(false, 7.9)
	item2 := *item.NewItem(true, 20.0)
	item3 := *item.NewItem(false, 100.50)
	var s []item.ItemAgent

	customer1 := newCustomer(s, 60.0)
	customer1.AddItem(item1)
	customer1.AddItem(item2)
	customer1.AddItem(item3)
	fmt.Println(*customer1)
	printSlice(customer1.items)

}

func printSlice(s []item.ItemAgent) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
