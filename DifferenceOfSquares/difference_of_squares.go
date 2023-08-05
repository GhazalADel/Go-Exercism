package diffsquares

func SquareOfSum(n int) int {
	tot := 0
	for i := 1; i <= n; i++ {
		tot += i
	}
	return tot * tot
}

func SumOfSquares(n int) int {
	tot := 0
	for i := 1; i <= n; i++ {
		tot += i * i
	}
	return tot
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
