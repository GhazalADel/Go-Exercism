// Package house is fun.
package house

const testVersion = 1

// Song is a function too.
func Song() string {
	var s string
	for i := 1; i <= 12; i++ {
		s += Verse(i)
		if i != 12 {
			s += "\n\n"
		}
	}
	return s
}

// Verse is a function.
func Verse(i int) string {
	return "This is " + ending(i) + cont(i-1)
}

func cont(i int) string {
	if i == 0 {
		return ""
	}
	return begining(i) + ending(i) + cont(i-1)
}

func begining(i int) string {
	switch i {
	case 11:
		return "that belonged to "
	case 10:
		return "that kept "
	case 9:
		return "that woke "
	case 8:
		return "that married "
	case 7:
		return "that kissed "
	case 6:
		return "that milked "
	case 5:
		return "that tossed "
	case 4:
		return "that worried "
	case 3:
		return "that killed "
	case 2:
		return "that ate "
	case 1:
		return "that lay in "
	}
	return ""
}

func ending(i int) string {
	switch i {
	case 12:
		return "the horse and the hound and the horn\n"
	case 11:
		return "the farmer sowing his corn\n"
	case 10:
		return "the rooster that crowed in the morn\n"
	case 9:
		return "the priest all shaven and shorn\n"
	case 8:
		return "the man all tattered and torn\n"
	case 7:
		return "the maiden all forlorn\n"
	case 6:
		return "the cow with the crumpled horn\n"
	case 5:
		return "the dog\n"
	case 4:
		return "the cat\n"
	case 3:
		return "the rat\n"
	case 2:
		return "the malt\n"
	case 1:
		return "the house that Jack built."
	}
	return ""
}
