package item

import (
	"math"
	"math/rand"
	"time"
)

type ItemAgent struct {
	EighteenPlus bool
	Price        float64
	Handling     float64
}

func NewItem(UpperBound float64, LowerBound float64) ItemAgent {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	ia := ItemAgent{}

	ia.EighteenPlus = setAgeLimit()
	ia.Price = math.Round((r.Float64()*10)*100) / 100
	ia.Handling = math.Round(((r.Float64()*(UpperBound-LowerBound))+LowerBound)*100) / 100

	return ia
}

func setAgeLimit() bool {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	eighteenPlus := false

	if (math.Round((r.Float64()*(1))*100) / 100) > 0.8 {
		eighteenPlus = true
	}

	return eighteenPlus
}