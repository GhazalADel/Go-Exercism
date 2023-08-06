package resistorcolorduo

import (
	"strconv"
	"strings"
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	tmp := ""
	for i := 0; i < 2; i++ {
		tmp += strconv.Itoa(ColorCode(colors[i]))
	}
	res, _ := strconv.Atoi(tmp)
	return res
}

func ColorCode(color string) int {
	color = strings.ToLower(color)
	colorCodes := map[string]int{
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
	return colorCodes[color]
}
