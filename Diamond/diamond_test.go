package diamond

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"
	"time"
)

type test struct {
	input rune
	want  string
}

func TestDiamond(t *testing.T) {
	tests := []test{
		{'A', "A\n"},
		{'B', " A \nB B\n A \n"},
		{'C', "  A  \n B B \nC   C\n B B \n  A  \n"},
		{'D', "   A   \n  B B  \n C   C \nD     D\n C   C \n  B B  \n   A   \n"},
	}
	for _, test := range tests {
		got, err := Gen(byte(test.input))

		if err != nil {
			t.Errorf("Gen(byte(%v)) err %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("Gen(byte(%v) got\n%v want\n%v", test.input, got, test.want)
		}
	}
}

var config = &quick.Config{Rand: rand.New(rand.NewSource(time.Now().UnixNano()))}

type correctChar byte

func (c correctChar) Generate(r *rand.Rand, _ int) reflect.Value {
	return reflect.ValueOf(correctChar('A' + r.Intn('Z'-'A'+1)))
}

func checkCorrect(requirement func(byte, []string) bool, keepSeparator bool, t *testing.T) {
	assertion := func(char correctChar) bool {
		d, err := Gen(byte(char))
		if err != nil {
			return false
		}
		separator := strings.Split
		if keepSeparator {
			separator = strings.SplitAfter
		}
		rows := separator(d, "\n")
		if len(rows) < 2 {
			return false
		}
		return requirement(byte(char), rows[:len(rows)-1])
	}
	if err := quick.Check(assertion, config); err != nil {
		t.Error(err)
	}
}

func TestFirstRowContainsOneA(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		return len(rows) > 0 && strings.Count(rows[0], "A") == 1
	}
	checkCorrect(requirement, false, t)
}

func TestLastRowContainsOneA(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		fmt.Printf("row %v count %v\n", rows[len(rows)-1], strings.Count(rows[len(rows)-1], "A"))
		return len(rows) > 0 && strings.Count(rows[len(rows)-1], "A") == 1
	}
	checkCorrect(requirement, false, t)
}

func TestAllRowsIdenticalLettersExceptFirstAndLast(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		for i, row := range rows {
			r := strings.TrimSpace(row)
			if r[0] != r[len(r)-1] {
				return false
			}
			if len(r) < 2 && i != 0 && i != len(rows)-1 {
				return false
			}
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestAllRowsHaveSameTrailingSpaces(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		for _, row := range rows {
			if row == "" {
				return false
			}
			for i, j := 0, len(row)-1; i < j && row[i] == ' '; i, j = i+1, j-1 {
				if row[j] != ' ' {
					return false
				}
			}
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestDiamondIsHorizontallySymmetric(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		for _, row := range rows {
			l := len(row)
			for i := l/2 - 1; i >= 0; i-- {
				if row[i] != row[l-1-i] {
					return false
				}
			}
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestDiamondIsVerticallySymmetric(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
			if rows[i] != rows[j] {
				return false
			}
		}
		return true
	}
	checkCorrect(requirement, true, t)
}

func TestDiamondIsSquare(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		if int(char-'A')*2+1 != len(rows) {
			return false
		}
		for _, row := range rows {
			if len(row) != len(rows) {
				return false
			}
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestDiamondHasItsShape(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		var n int
		for i, row := range rows {
			s := len(strings.TrimSpace(row))
			if i > len(rows)/2 && n <= s {
				return false
			} else if i <= len(rows)/2 && n >= s {
				return false
			}
			n = s
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestTopHalfHasAscendingLetters(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		var start byte = 'A' - 1
		for i := 0; i <= len(rows)/2; i++ {
			s := strings.TrimLeft(rows[i], " ")
			if s == "" || s[0] <= start {
				return false
			}
			start = s[0]
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestBottomHalfHasDescendingLetters(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		var start byte = 'A' - 1
		for i := len(rows) - 1; i > len(rows)/2; i-- {
			s := strings.TrimLeft(rows[i], " ")
			if s == "" || s[0] <= start {
				return false
			}
			start = s[0]
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

func TestDiamondFourCornersAreTriangle(t *testing.T) {
	requirement := func(char byte, rows []string) bool {
		notSpace := func(r rune) bool { return r <= 'Z' && r >= 'A' }
		var n int
		for i, row := range rows {
			s := strings.IndexFunc(row, notSpace)
			e := len(row) - strings.LastIndexFunc(row, notSpace) - 1
			switch {
			case s != e:
				return false
			case i == 0:
				n = s
			default:
				if i > len(rows)/2 && n >= s {
					return false
				} else if i <= len(rows)/2 && n <= s {
					return false
				}
				n = s
			}
		}
		return true
	}
	checkCorrect(requirement, false, t)
}

type wrongChar byte

func (c wrongChar) Generate(r *rand.Rand, _ int) reflect.Value {
	b := rand.Intn(256)
	for ; b >= 'A' && b <= 'Z'; b = r.Intn(256) {
	}
	return reflect.ValueOf(wrongChar(b))
}

func TestCharOutOfRangeShouldGiveError(t *testing.T) {
	assertion := func(char wrongChar) bool {
		_, err := Gen(byte(char))
		return err != nil
	}
	if err := quick.Check(assertion, config); err != nil {
		t.Error(err)
	}
}
