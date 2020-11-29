package item

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type ItemAgent struct {
	eighteenPlus bool
	price        float64
	handling     float64
}

func NewItem(UpperBound float64, LowerBound float64) ItemAgent {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	ia := ItemAgent{}

	ia.eighteenPlus = setAgeLimit()
	ia.price = math.Round((r.Float64()*10)*100) / 100
	ia.handling = math.Round(((r.Float64()*((UpperBound+1)-LowerBound))+LowerBound)*100) / 100

	return ia
}

func (ia *ItemAgent) GetItemHandling() float64 {
	return ia.handling
}

func (ia *ItemAgent) GetAgeRating() bool {
	return ia.eighteenPlus
}

func (ia *ItemAgent) GetPrice() float64 {
	return ia.price
}

func setAgeLimit() bool {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	eighteenPlus := false

	if (math.Round((r.Float64()*(1))*100) / 100) > 0.8 {
		eighteenPlus = true
	}

	return eighteenPlus
}

//for testing
func PrintItems() {
	for i := 0; i < 10; i++ {
		fmt.Println(NewItem(55.0, 65.0))
	}
}
