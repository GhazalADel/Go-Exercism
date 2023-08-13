package yacht

func count(dice []int, target int) int {
	count := 0
	for _, v := range dice {
		if v == target {
			count++
		}
	}
	return count
}
func Score(dice []int, category string) int {
	res := 0
	if category == "ones" {
		return count(dice, 1)
	} else if category == "twos" {
		return count(dice, 2) * 2
	} else if category == "threes" {
		return count(dice, 3) * 3
	} else if category == "fours" {
		return count(dice, 4) * 4
	} else if category == "fives" {
		return count(dice, 5) * 5
	} else if category == "sixes" {
		return count(dice, 6) * 6
	} else if category == "full house" {
		mapping := make(map[int]int)
		for _, v := range dice {
			mapping[v]++
		}
		hasThree := false
		hasTwo := false
		for _, v := range mapping {
			if v == 3 {
				hasThree = true
			}
			if v == 2 {
				hasTwo = true
			}
		}
		if hasThree && hasTwo {
			return Score(dice, "choice")
		}
	} else if category == "four of a kind" {
		t := -1
		for i := 1; i <= 6; i++ {
			if count(dice, i) >= 4 {
				t = i
				break
			}
		}
		if t != -1 {
			res = t * 4
		}
	} else if category == "little straight" {
		tmp := 0
		for i, v := range dice {
			tmp ^= (i + 1)
			tmp ^= v
		}
		if tmp == 0 {
			res = 30
		}
	} else if category == "big straight" {
		tmp := 0
		for i, v := range dice {
			tmp ^= (i + 2)
			tmp ^= v
		}
		if tmp == 0 {
			res = 30
		}
	} else if category == "choice" {
		for _, v := range dice {
			res += v
		}
	} else if category == "yacht" {
		for i := 1; i <= 6; i++ {
			if count(dice, i) == 5 {
				return 50
			}
		}
	}

	return res
}
