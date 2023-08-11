package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 && len(l2) != 0 {
		return RelationSublist
	}
	if len(l2) == 0 && len(l1) != 0 {
		return RelationSuperlist
	}
	if len(l1) == len(l2) {
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				return RelationUnequal
			}
		}
		return RelationEqual
	}
	firstIsBigger := true
	if len(l1) < len(l2) {
		firstIsBigger = false
		l1, l2 = l2, l1
	}
	start := l2[0]
	i := 0
	for i < len(l1) {
		if l1[i] == start {
			isSubList := true
			for j := 0; j < len(l2); j++ {
				if l1[i+j] != l2[j] {
					isSubList = false
					break
				}
			}
			if isSubList {
				if firstIsBigger {
					return RelationSuperlist
				} else {
					return RelationSublist
				}
			}
		}
		i++
	}
	return RelationUnequal
}
