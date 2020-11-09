package floorStaff

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Floor Staff agent struct
type floorStaff struct {
 	cleaningTime float64
	diligenceFactor float64
	baseHelpfulness float64
	actualHelpfulness float64
	occupied bool
}

var x = rand.New(rand.NewSource(time.Now().UnixNano()))

// Floor Staff agent constructor
func NewFloorStaff () *floorStaff {

	// Create staff agent 
	staff := floorStaff {}

	// Randomly initialised variables 
	staff.cleaningTime = math.Round(((x.Float64()*(0.75-0.25))+0.25)*100) / 100
	staff.diligenceFactor = math.Round(((x.Float64()*(0.75-0.25))+0.25)*100) / 100
	staff.baseHelpfulness = math.Round(((x.Float64()*(0.75-0.25))+0.25)*100) / 100

	// Dynamically defined variables 
	staff.actualHelpfulness = calcActualHelpfulness(staff.diligenceFactor, staff.baseHelpfulness)

	// Initialised False
	staff.occupied = false

	// Return staff object
	return &staff
}

// Getter for staff occupied status
func (staff *floorStaff) GetOccupied() (bool) {
	return staff.occupied
}

// Setter for staff occupied status
func (staff *floorStaff) SetOccupied(val bool) {
	staff.occupied = val
}

// Dynamically calculate the actual helpfulness of the floor agent
func calcActualHelpfulness (diligenceFactor, baseHelpfulness float64) (float64) {
	return ((diligenceFactor + baseHelpfulness) / 2)
}

// Print floor agent variables 
func (staff *floorStaff) PrintStaff() {
	var ct = staff.cleaningTime
	var df = staff.diligenceFactor
	var bh = staff.baseHelpfulness
	var ah = staff.actualHelpfulness
	var os = staff.occupied
	fmt.Printf("Cleaning Time:%.2f, Diligence Factor:%.2f, Base Helpfulness:%.2f, Actual Helpfulness:%.2f, Occupied Status:%t\n", ct, df, bh, ah, os)
}