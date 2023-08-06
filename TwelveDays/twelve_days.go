package twelve

import "fmt"

var dayToOrdinal = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var dayToGift = map[int]string{
	12: "twelve Drummers Drumming",
	11: "eleven Pipers Piping",
	10: "ten Lords-a-Leaping",
	9:  "nine Ladies Dancing",
	8:  "eight Maids-a-Milking",
	7:  "seven Swans-a-Swimming",
	6:  "six Geese-a-Laying",
	5:  "five Gold Rings",
	4:  "four Calling Birds",
	3:  "three French Hens",
	2:  "two Turtle Doves",
	1:  "a Partridge in a Pear Tree",
}

func Verse(i int) string {
	gifts := ""
	if i != 1 {
		for j := i; j > 1; j-- {
			gifts += dayToGift[j] + ", "
		}
		gifts += "and a Partridge in a Pear Tree."
	} else {
		gifts = "a Partridge in a Pear Tree."
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", dayToOrdinal[i], gifts)
}

func Song() string {
	song := ""
	for i := 1; i < 12; i++ {
		song += Verse(i) + "\n"
	}
	song += Verse(12)
	return song
}
