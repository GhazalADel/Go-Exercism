package main

import (
	"fmt"
)

func Gen(char byte) (string, error) {
	if char == 'A' {
		return "A", nil
	}
	start := int(char) - 63
	tmp := start
	startChar := 65
	res := ""
	for i := 0; i < tmp; i++ {
		for j := 0; j < start; j++ {
			res += "."
		}
		res += string(rune(startChar))
		for j := 0; j < start; j++ {
			res += "."
		}
		res += "\n"
		start--
		startChar++
	}
	res += string(char)
	for i := 0; i < 2*tmp-1; i++ {
		res += "."
	}
	res += string(char) + "\n"
	startChar--
	start++
	for i := 0; i < tmp; i++ {
		for j := 0; j < start; j++ {
			res += "."
		}
		res += string(rune(startChar))
		for j := 0; j < start; j++ {
			res += "."
		}
		res += "\n"
		start++
		startChar--
	}
	return res, nil
}

func main() {
	r, _ := Gen('C')
	fmt.Printf("%s", r)
}
