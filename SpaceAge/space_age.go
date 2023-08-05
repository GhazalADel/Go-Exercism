package space

import "math"

type Planet string

const (
	EarthYear = 31557600
	Mercury   = 0.2408467
	Venus     = 0.61519726
	Earth     = 1.0
	Mars      = 1.8808158
	Jupiter   = 11.862615
	Saturn    = 29.447498
	Uranus    = 84.016846
	Neptune   = 164.79132
)

func Age(sec float64, p Planet) float64 {
	var age float64
	switch p {
	case "Mercury":
		age = Mercury
	case "Venus":
		age = Venus
	case "Earth":
		age = Earth
	case "Mars":
		age = Mars
	case "Jupiter":
		age = Jupiter
	case "Saturn":
		age = Saturn
	case "Uranus":
		age = Uranus
	case "Neptune":
		age = Neptune
	default:
		age = math.NaN()
	}
	return sec / EarthYear / age
}
