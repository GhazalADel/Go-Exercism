package main

import "fmt"

func Replace(log string, oldRune, newRune rune) string {
	tmp := ""
	for _, char := range log {
		if char == oldRune {
			tmp += string(newRune)
		} else {
			tmp += string(char)
		}
	}
	return tmp
}
func NewVoteCounter(initialVotes int) *int {
	var res *int = &initialVotes
	return res
}
func main() {
	log := "please replace '👎' with '👍'"

	fmt.Println(Replace(log, '👎', '👍'))

	fmt.Println(NewVoteCounter(12))
}
