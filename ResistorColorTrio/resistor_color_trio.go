package resistorcolortrio

import (
	"math"
	"strconv"
)

var resistor = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	the_zerooos := int(math.Pow(10, float64(resistor[colors[2]])))
	if the_zerooos == 0 {
		the_zerooos = 1
	}
	rvalue := (resistor[colors[0]]*10 + resistor[colors[1]]) * the_zerooos
	suffix := ""
	switch {
	case rvalue > 1000000000:
		rvalue /= 1000000000
		suffix = " giga"
	case rvalue > 1000000:
		rvalue /= 1000000
		suffix = " mega"
	case rvalue > 1000:
		rvalue /= 1000
		suffix = " kilo"
	default:
		suffix = " "
	}
	return strconv.Itoa(rvalue) + suffix + "ohms"
}
