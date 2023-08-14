package robotname

import (
	"errors"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

const maxNumberOfNames = 26 * 26 * 10 * 10 * 10

var robotNames = map[string]bool{}

func (r *Robot) Name() (name string, err error) {
	if len(robotNames) == maxNumberOfNames {
		return "", errors.New("")
	}
	if r.name == "" {
		r.name = generateName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = generateName()
}

func generateName() string {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	for {
		pre := ""
		for i := 0; i < 2; i++ {
			n := y1.Intn(26) + 65
			pre += string(rune(n))
		}
		suf := ""
		for i := 0; i < 3; i++ {
			n := y1.Intn(10) + 48
			suf += string(rune(n))
		}
		name := pre + suf
		if !robotNames[name] {
			robotNames[name] = true
			return name
		}
	}
}
