package clock

import "strconv"

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	h, m = Normalize(h, m)
	return Clock{h: h, m: m}
}

func (c Clock) Add(m int) Clock {
	c.m += m
	c.h, c.m = Normalize(c.h, c.m)
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m -= m
	c.h, c.m = Normalize(c.h, c.m)
	return c
}

func (c Clock) String() string {
	hStr := ""
	mStr := ""
	if len(strconv.Itoa(c.h)) == 1 {
		hStr += "0"
	}
	if len(strconv.Itoa(c.m)) == 1 {
		mStr += "0"
	}
	hStr += strconv.Itoa(c.h)
	mStr += strconv.Itoa(c.m)
	return hStr + ":" + mStr
}
func Normalize(h, m int) (int, int) {
	if m >= 60 {
		h += m / 60
		m %= 60
	}
	if m < 0 {
		m *= -1
		if m%60 == 0 {
			h -= (m / 60)
		} else {
			h -= ((m / 60) + 1)
		}
		tmp := m % 60
		if tmp != 0 {
			m = 60 - tmp
		} else {
			m = 0
		}
	}
	for h >= 24 {
		h -= 24
	}
	for h < 0 {
		h += 24
	}
	return h, m
}
