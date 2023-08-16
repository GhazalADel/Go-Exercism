package allergies

import (
	"math"
)

func getAllergieValue(item string) uint {
	mapping := map[string]int{
		"eggs":         1,
		"peanuts":      2,
		"shellfish":    3,
		"strawberries": 4,
		"tomatoes":     5,
		"chocolate":    6,
		"pollen":       7,
		"cats":         8,
	}
	return uint(math.Pow(2, float64(mapping[item]-1)))
}
func findMin(target uint) int {
	if target >= 128 {
		return 7
	} else if target >= 64 {
		return 6
	} else if target >= 32 {
		return 5
	} else if target >= 16 {
		return 4
	} else if target >= 8 {
		return 3
	} else if target >= 4 {
		return 2
	} else if target >= 2 {
		return 1
	}
	return 1
}
func canMade(target uint, used map[string]bool) (bool, []string) {
	items := []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
	res := make([]string, 0)
	m := findMin(target)
	for target > 0 {
		if m < 0 {
			break
		}
		if uint(math.Pow(2, float64(m))) > target {
			m--
		} else if _, ok := used[items[m]]; ok {
			m--
		} else {
			used[items[m]] = true
			target -= uint(math.Pow(2, float64(m)))
			m--
		}
	}
	for k, _ := range used {
		res = append(res, k)
	}
	if target == 0 {
		return true, res
	}
	return false, []string{}

}
func Allergies(allergies uint) []string {
	allergies %= 256
	_, res := canMade(allergies, map[string]bool{})
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	if allergies == 0 {
		return false
	}
	val := getAllergieValue(allergen)
	if allergies < val || allergies-val == val {
		return false
	}
	if ok, _ := canMade(allergies-val, map[string]bool{allergen: true}); ok {
		return true
	}
	return false
}
