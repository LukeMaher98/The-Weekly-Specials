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
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewItem() *ItemAgent {
	r.Seed(time.Now().UTC().UnixNano())
	ia := ItemAgent{}

	ia.eighteenPlus = (r.Intn(2) == 1)
	ia.price = math.Round((rand.Float64()*10)*100) / 100

	return &ia
}

//for testing
func PrintItems() {
	for i := 0; i < 10; i++ {
		fmt.Println(*NewItem())
	}
}
