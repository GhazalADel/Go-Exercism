package lasagna

func PreparationTime(layers []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(layers) * t
}

func Quantities(layers []string) (int, float64) {
	n, s := 0, 0
	for _, v := range layers {
		if v == "noodles" {
			n++
		} else if v == "sauce" {
			s++
		}
	}
	return n * 50, float64(s) * 0.2
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, n int) []float64 {
	res := make([]float64, 0)
	for _, v := range quantities {
		res = append(res, v*(float64(n)/2))
	}
	return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
