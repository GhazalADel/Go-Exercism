package binarysearch

func SearchInts(list []int, key int) int {
	left, right := 0, len(list)-1
	for left <= right {
		middle := left + (right-left)/2
		if list[middle] == key {
			return middle
		}
		if list[middle] < key {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
