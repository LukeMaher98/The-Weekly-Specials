package customer

import (
	"math"
	"math/rand"
	"src/item"
	"time"
)

type CustomerAgent struct {
	items               []item.ItemAgent
	impairmentFactor    float64
	replaceItem         float64
	couponItem          float64
	withChildren        bool
	loyaltyCard         bool
	baseAmicability     float64
	customerAmicability float64
	preferredPayment    float64
	//avaliablePayment        []int
	baggintTimeSelfCheckout float64
	emergencyLeave          float64
	switchLine              float64
	competence              float64
	trolleyLimit            int
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewCustomer() CustomerAgent {
	ca := CustomerAgent{}

	//float values
	ca.impairmentFactor = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	ca.replaceItem = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	ca.baseAmicability = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	ca.customerAmicability = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	ca.preferredPayment = math.Round((r.Float64()*2)*100) / 100
	ca.emergencyLeave = math.Round((r.Float64()*1)*100) / 100
	ca.switchLine = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100
	ca.competence = math.Round(((r.Float64()*(0.75-0.25))+0.25)*100) / 100

	//int values
	ca.trolleyLimit = r.Intn(100) + 1

	//bool values
	ca.withChildren = (r.Intn(2) == 1)
	ca.loyaltyCard = (r.Intn(2) == 1)

	//generate items
	ca.items = generateTrolley()

	return ca
}

func generateTrolley() []item.ItemAgent {
	var trolley []item.ItemAgent
	var trolleyLimit = r.Intn(100) + 1

	for i := 0; i < trolleyLimit; i++ {
		trolley = append(trolley, *item.NewItem())
	}

	return trolley
}
