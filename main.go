package main

import (
	"fmt"
	"os"

	sleepiq "github.com/danpenn/SleepIQ"
)

// Bed Preset Positions
const (
	PositionFavorite = 1
	PositionRead     = 2
	PositionWatchTV  = 3
	PositionFlat     = 4
	PositionZeroG    = 5
	PositionSnore    = 6
)

func main() {
	fmt.Println("Reading environment variable for password")
	var password string
	var username string
	username = os.Getenv("SLEEPIQ_USER")
	password = os.Getenv("SLEEPIQ_PASS")

	// Create a new instance of SleepIQ
	siq := sleepiq.New()

	// Login
	_, err := siq.Login(username, password)
	if err != nil {
		fmt.Println("login failed - ", err)
		return
	}

	// Get information about all the beds
	beds, err := siq.Beds()
	if err != nil {
		fmt.Println("could not get beds - ", err)
		return
	}

	// Display some bed information
	for _, bed := range beds.Beds {
		fmt.Printf("%s (%s)\n", bed.Size, bed.Name)
	}

	// Set the right side of the bed to the 'WatchTV' preset position
	bedStatus, err := siq.ControlBedPosition(beds.Beds[0].BedID, "right", PositionWatchTV)

	// bedStatus, err := siq.ControlBedPosition(beds.Beds[0].BedID, "Right", siq.WatchTV)
	if err != nil {
		fmt.Println("could not set bed position - ", err)
		return
	}

	// Display a confirmation showing the new position
	fmt.Printf("Position: %s", bedStatus.CurrentPositionPreset)
}
