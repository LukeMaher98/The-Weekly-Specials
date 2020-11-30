package item

import (
	"math"
	"math/rand"
	"time"
)

type ItemAgent struct {
	eighteenPlus bool
	price        float64
	handling     float64
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewItem(UpperBound float64, LowerBound float64) ItemAgent {
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
	eighteenPlus := false

	if (math.Round((r.Float64()*(1))*100) / 100) > 0.8 {
		eighteenPlus = true
	}

	return eighteenPlus
}