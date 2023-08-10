package sieve

func Sieve(limit int) (primes []int) {
	if limit < 2 {
		return []int{}
	}
	nums := make([]int, 0)
	for i := 2; i <= limit; i++ {
		nums = append(nums, i)
	}
	for len(nums) != 0 {
		prime := nums[0]
		nums = nums[1:]

		nums = removeMultiples(nums, prime)
		primes = append(primes, prime)
	}
	return primes
}

func removeMultiples(slice []int, prime int) (filtered []int) {
	if len(slice) == 0 {
		return slice
	}
	multiplyer := 2
	multiple := multiplyer * prime
	for len(slice) != 0 && multiple <= slice[len(slice)-1] {
		slice = remove(slice, multiple)
		multiplyer += 1
		multiple = multiplyer * prime
	}
	return slice
}

func remove(slice []int, value int) (result []int) {
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}
