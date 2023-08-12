package stringset

// Implement Set as a collection of unique string values.
type Set struct {
	elements map[string]bool
}

//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

func New() Set {
	return Set{elements: make(map[string]bool)}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, v := range l {
		if _, ok := s.elements[string(v)]; !ok {
			s.elements[string(v)] = true
		}
	}
	return s
}

func (s Set) String() string {
	res := "{"
	for k, _ := range s.elements {
		res += "\""
		res += k
		res += "\", "
	}
	if len(s.elements) != 0 {
		res = res[:len(res)-2]
	}
	res += "}"
	return res
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.elements[elem]
	return ok
}

func (s Set) Add(elem string) {
	if !s.Has(elem) {
		s.elements[elem] = true
	}
}

func Subset(s1, s2 Set) bool {
	for k, _ := range s1.elements {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}
	return len(Intersection(s1, s2).elements) == 0
}

func Equal(s1, s2 Set) bool {
	if len(s1.elements) != len(s2.elements) {
		return false
	}
	for k, _ := range s1.elements {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	if len(s1.elements) == 0 {
		return Set{elements: make(map[string]bool)}
	}
	if len(s2.elements) == 0 {
		return Set{elements: make(map[string]bool)}
	}

	res := New()
	for k, _ := range s2.elements {
		if s1.Has(k) {
			res.elements[k] = true
		}
	}
	return res
}

func Difference(s1, s2 Set) Set {
	for k, _ := range s1.elements {
		if s2.Has(k) {
			delete(s1.elements, k)
		}
	}
	return s1
}

func Union(s1, s2 Set) Set {
	if len(s1.elements) == 0 {
		return s2
	}
	if len(s2.elements) == 0 {
		return s1
	}
	for k, _ := range s2.elements {
		if !s1.Has(k) {
			s1.elements[k] = true
		}
	}
	return s1
}
