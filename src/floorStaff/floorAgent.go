package floorStaff

import (
	"math"
	"math/rand"
	"time"
)

// Floor Staff agent struct
type FloorStaff struct {
	managerBoost      float64
	Amicability       float64
	Competance        float64
	OccupyingCustomer bool
	ManagerOccupied   bool
}

// Floor Staff agent constructor
func CreateInitialisedFloorStaffAgent(AmicLB, AmicUB, CompLB, CompUB float64) FloorStaff {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create staff agent
	staff := FloorStaff{}

	//Initialised to 1
	staff.managerBoost = 1

	// Randomly initialised variables based on boundings
	staff.Amicability = math.Round(((r.Float64()*(AmicUB-AmicLB))+AmicLB)*100) / 100
	staff.Competance = math.Round(((r.Float64()*(CompUB-CompLB))+CompLB)*100) / 100

	// Initialised False
	staff.OccupyingCustomer = false
	staff.ManagerOccupied = false

	// Return staff object
	return staff
}

func (fs *FloorStaff) PropagateTime() {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	// 50% chance to change occupied status
	if r.Float64() < (fs.Amicability * fs.managerBoost) {
		fs.OccupyingCustomer = !fs.OccupyingCustomer
	}

}

// ManagerPresent : applies a boost to the cashier
func (fs *FloorStaff) ManagerPresent(boost float64) {
	fs.managerBoost = boost + 1.00
}

// ManagerAbsent : reverts the boost to the cashier
func (fs *FloorStaff) ManagerAbsent() {
	fs.managerBoost = 1
}
