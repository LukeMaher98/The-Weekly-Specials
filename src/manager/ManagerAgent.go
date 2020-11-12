package manager

import (
	"math"
	"math/rand"
	"time"
)

type ManagerAgent struct {
	baseHelpfulness float64
	baseMoraleBoost float64

	ActualHelpfulness float64
	ActualMoraleBoost float64
	OnFloor           bool
	//SupervisingCheckout checkout
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewManager() ManagerAgent {
	manager := ManagerAgent{}

	manager.baseHelpfulness = math.Round(((r.Float64()*0.5)+0.25)*100) / 100
	manager.baseMoraleBoost = math.Round(((r.Float64()*1.5)-0.75)*100) / 100

	manager.ActualHelpfulness = manager.baseHelpfulness
	manager.ActualMoraleBoost = manager.baseMoraleBoost
	manager.OnFloor = true
	//manager.SupervisingCheckout = null

	return manager
}
