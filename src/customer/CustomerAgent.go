package customer

type CustomerAgent struct {
	//Item
	impairmentFactor float64
	//replaceItem float64
	//couponItem float64
	//withChildren bool
	//loyaltyCard bool
	//baseAmicability float64
	//customerAmicability float64
	//preferredPayment float64
	//avaliablePayment int[]
	//baggintTimeSelfCheckout float64
	//emergencyLeave float64
	//switchLine float64
}

func newCustomer(impairmentFactor float64) *CustomerAgent {

	ca := CustomerAgent{impairmentFactor: impairmentFactor}
	return &ca
}
