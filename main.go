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
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		return DescribeNumber(float64(i.(int)))
	case float64:
		return DescribeNumber(i.(float64))
	}
	return "Return to sender"
}
func main() {
	fmt.Println(DescribeNumber(2))
	fmt.Println(DescribeNumber(2.787))
	fmt.Println(DescribeAnything(42))
}
