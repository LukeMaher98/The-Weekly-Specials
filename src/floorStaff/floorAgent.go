package floorStaff

import (
	"math"
	"math/rand"
	"time"
)

// Floor Staff agent struct
type FloorStaffAgent struct {
 	Amicability float64
	Competance float64
	Occupied bool
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Floor Staff agent constructor
func CreateInitialisedFloorStaffAgent(AmicLB, AmicUB, CompLB, CompUB float64) FloorStaffAgent {

	// Create staff agent 
	staff := FloorStaffAgent {}

	// Randomly initialised variables based on boundings
	staff.Amicability = math.Round(((r.Float64()*(AmicUB-AmicLB))+AmicLB)*100) / 100
	staff.Competance = math.Round(((r.Float64()*(CompUB-CompLB))+CompLB)*100) / 100

	// Initialised False
	staff.Occupied = false

	// Return staff object
	return staff
}

// Placeholder stuff 
func (fs *FloorStaffAgent) PropagateTime() {
	// 50% chance to change occupied status
	if r.Float64() < 0.5 {
		fs.Occupied = !fs.Occupied
	}

	// Placeholder example 
	if fs.Occupied {
		fs.Amicability += 0.001
		fs.Competance += 0.001
	} else {
		fs.Amicability -= 0.001
		fs.Competance -= 0.001
	}
}