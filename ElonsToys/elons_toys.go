package elon

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, distance: 0, battery: 100}
}

// for tests
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	count := float64(trackDistance) / float64(car.speed)
	if count*float64(car.batteryDrain) <= float64(car.battery) {
		return true
	}
	return false
}
