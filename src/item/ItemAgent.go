package main

import "fmt"

type ItemAgent struct {
	eighteenPlus bool
	price        float64
}

func newItem(eighteenPlus bool, price float64) *ItemAgent {

	ia := ItemAgent{eighteenPlus: eighteenPlus, price: price}
	return &ia
}

func (item *ItemAgent) GetEighteenPlus() bool {
	return item.eighteenPlus
}

func main() {

	fmt.Println(ItemAgent{eighteenPlus: true, price: 6.6})
	fmt.Println(*newItem(true, 6.6))

	item1 := newItem(false, 7.9)
	fmt.Println("Hello")

	fmt.Println(item1.GetEighteenPlus())

}
