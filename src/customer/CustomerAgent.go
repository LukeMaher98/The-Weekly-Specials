package customer

import (
	"math"
	"math/rand"
	"src/item"
	"time"
)

type CustomerAgent struct {
	Items               []item.ItemAgent
	ImpairmentFactor    float64
	ReplaceItem         float64
	CouponItem          float64
	WithChildren        bool
	LoyaltyCard         bool
	BaseAmicability     float64
	CustomerAmicability float64
	PreferredPayment    float64
	//avaliablePayment        []int
	BaggintTimeSelfCheckout float64
	EmergencyLeave          float64
	SwitchLine              float64
	Competence              float64
	TrolleyLimit            int
	FinishedShop            bool
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewCustomer() CustomerAgent {
	ca := CustomerAgent{}

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	//float values
	ca.ImpairmentFactor = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.ReplaceItem = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.BaseAmicability = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.CustomerAmicability = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.PreferredPayment = math.Round((r.Float64()*2)*100) / 100
	ca.EmergencyLeave = math.Round((r.Float64()*1)*100) / 100
	ca.SwitchLine = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100
	ca.Competence = math.Round(((r.Float64()*(0.5))+0.25)*100) / 100

	//int values
	ca.TrolleyLimit = r.Intn(100) + 1

	//bool values
	ca.WithChildren = (r.Intn(2) == 1)
	ca.LoyaltyCard = (r.Intn(2) == 1)

	//generate items
	ca.Items = GenerateTrolley()

	//discuss this
	ca.FinishedShop = false

	return ca
}

func GenerateTrolley() []item.ItemAgent {
	//This obviously needs to be changed
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var trolley []item.ItemAgent
	var trolleyLimit = r.Intn(100) + 1

	for i := 0; i < trolleyLimit; i++ {
		trolley = append(trolley, *item.NewItem())
	}

	return trolley
}

func (ca *CustomerAgent) PropagateTime() {
	//each time iteration add an item to customers trolley or check if they are done shopping?

}

/*func () ShoppingFinished {

}*/
