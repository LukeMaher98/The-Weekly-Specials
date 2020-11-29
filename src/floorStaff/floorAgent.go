package floorStaff

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Floor Staff agent struct
type FloorStaff struct {
	cleaningTime      float64
	diligenceFactor   float64
	BaseHelpfulness   float64
	actualHelpfulness float64
	Occupied          bool
	Competence        float64
	managerBoost      float64
}

var x = rand.New(rand.NewSource(time.Now().UnixNano()))

// Floor Staff agent constructor
func NewFloorStaff() *FloorStaff {

	// Create staff agent
	staff := FloorStaff{}

	// Randomly initialised variables
	staff.cleaningTime = math.Round(((x.Float64()*(0.75-0.25))+0.25)*100) / 100
	staff.diligenceFactor = math.Round(((x.Float64()*(0.75-0.25))+0.25)*100) / 100
	staff.BaseHelpfulness = math.Round(((x.Float64()*(0.75-0.25))+0.25)*100) / 100

	//Initialised to 1
	staff.managerBoost = 1

	// Dynamically defined variables
	staff.actualHelpfulness = calcActualHelpfulness(staff.diligenceFactor, staff.BaseHelpfulness, staff.managerBoost)

	// Initialised False
	staff.Occupied = false

	// Return staff object
	return &staff
}

// Getter for staff occupied status
func (staff *FloorStaff) GetOccupied() bool {
	return staff.Occupied
}

// Setter for staff occupied status
func (staff *FloorStaff) SetOccupied(val bool) {
	staff.Occupied = val
}

// Dynamically calculate the actual helpfulness of the floor agent
func calcActualHelpfulness(diligenceFactor, baseHelpfulness, managerBoost float64) float64 {
	return (((diligenceFactor + baseHelpfulness) * managerBoost) / 2)
}

// Print floor agent variables
func (staff *FloorStaff) PrintStaff() {
	var ct = staff.cleaningTime
	var df = staff.diligenceFactor
	var bh = staff.BaseHelpfulness
	var ah = staff.actualHelpfulness
	var os = staff.Occupied
	fmt.Printf("Cleaning Time:%.2f, Diligence Factor:%.2f, Base Helpfulness:%.2f, Actual Helpfulness:%.2f, Occupied Status:%t\n", ct, df, bh, ah, os)
}

// ManagerPresent : applies a boost to the cashier
func (staff *FloorStaff) ManagerPresent(boost float64) {
	staff.managerBoost = boost + 1.00
}

// ManagerAbsent : reverts the boost to the cashier
func (staff *FloorStaff) ManagerAbsent() {
	staff.managerBoost = 1
}
