package floorStaff

import (
	"math"
	"math/rand"
	"time"
)

// Floor Staff agent struct
type FloorStaff struct {
 	Amicability float64
	Competance float64
	Occupied bool
}

// Floor Staff agent constructor
func CreateInitialisedFloorStaffAgent(AmicLB, AmicUB, CompLB, CompUB float64) FloorStaff {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create staff agent 
	staff := FloorStaff {}

	// Randomly initialised variables based on boundings
	staff.Amicability = math.Round(((r.Float64()*(AmicUB-AmicLB))+AmicLB)*100) / 100
	staff.Competance = math.Round(((r.Float64()*(CompUB-CompLB))+CompLB)*100) / 100

	// Initialised False
	staff.Occupied = false

	// Return staff object
	return staff
}

// Placeholder stuff 
func (fs *FloorStaff) PropagateTime() {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	// 50% chance to change occupied status
	if r.Float64() < fs.Amicability {
		fs.Occupied = !fs.Occupied
	}


}

func (fs *FloorStaff) something() {
	
}

// Code by Carl below
// type observer interface {
// 	update(float64)
// }

// func (staff *FloorStaff) update(managerComp float64) {
// 	//do something with manager competence
// }