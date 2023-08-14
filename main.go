package main

import (
	"errors"
	"fmt"
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
	delete(robotNames, r.name)
	r.name = generateName()
}

func generateName() string {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	for {
		pre := ""
		for i := 0; i < 2; i++ {
			n := y1.Intn(26) + 65
			pre += string(n)
		}
		suf := ""
		for i := 0; i < 3; i++ {
			n := y1.Intn(10) + 48
			suf += string(n)
		}

		name := pre + suf
		if !robotNames[name] {
			robotNames[name] = true
			return name
		}
	}
}

func main() {
	r := Robot{}
	n, err := r.Name()
	fmt.Printf("%v %v", n, err)
}
