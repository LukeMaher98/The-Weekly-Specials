package item

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type ItemAgent struct {
	EighteenPlus bool
	Price        float64
	Handling     float64
}

func NewItem() *ItemAgent {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	ia := ItemAgent{}

	ia.EighteenPlus = (r.Intn(2) == 1)
	ia.Price = math.Round((r.Float64()*10)*100) / 100
	ia.Handling = math.Round((r.Float64()*0.25)*100) / 100

	return &ia
}

//for testing
func PrintItems() {
	for i := 0; i < 10; i++ {
		fmt.Println(*NewItem())
	}
}
