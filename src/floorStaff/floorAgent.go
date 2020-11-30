package floorStaff

import (
	"math"
	"math/rand"
	"time"
)

// Floor Staff agent struct
type FloorStaff struct {
	managerBoost      float64
	amicability       float64
	competence        float64
	managerOccupied   bool
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Floor Staff agent constructor
func CreateInitialisedFloorStaffAgent(AmicLB, AmicUB, CompLB, CompUB float64) FloorStaff {

	// Create staff agent
	staff := FloorStaff{}

	//Initialised to 1
	staff.managerBoost = 1

	// Randomly initialised variables based on boundings
	staff.amicability = math.Round(((r.Float64()*(AmicUB-AmicLB))+AmicLB)*100) / 100
	staff.competence = math.Round(((r.Float64()*(CompUB-CompLB))+CompLB)*100) / 100
	
	staff.managerOccupied = false

	// Return staff object
	return staff
}

func (fs *FloorStaff) GetAmicability() float64 {
	return fs.amicability * fs.managerBoost
}

func (fs *FloorStaff) GetCompetence() float64 {
	return fs.competence * fs.managerBoost
}

// ManagerPresent : applies a boost to the cashier
func (fs *FloorStaff) ManagerPresent(boost float64) {
	fs.managerBoost = boost + 1.00
	fs.managerOccupied = true
}

// ManagerAbsent : reverts the boost to the cashier
func (fs *FloorStaff) ManagerAbsent() {
	fs.managerBoost = 1
	fs.managerOccupied = false
}
