package main

func ScaleRecipe(quantities []float64, n int) []float64 {
	res := make([]float64, 0)
	for _, v := range quantities {
		res = append(res, v*(float64(n)/2))
	}
	return res
}
func main() {
	ScaleRecipe([]float64{0.6, 300, 1, 0.5, 50, 0.1, 100}, 3)
}
