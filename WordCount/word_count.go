package wordcount

import "strings"

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	res := make(map[string]int)
	tmp := ""
	for _, l := range phrase {
		if (l >= 65 && l <= 90) || (l >= 97 && l <= 122) || (l >= 48 && l <= 57) || l == 39 {
			tmp += string(l)
		} else {
			if tmp != "" {
				if strings.Count(tmp, "'") != len(tmp) {
					for strings.HasPrefix(tmp, "'") {
						tmp = tmp[1:]
					}
					for strings.HasSuffix(tmp, "'") {
						tmp = tmp[:len(tmp)-1]
					}
					if tmp != "" {
						tmp = strings.ToLower(tmp)
						if _, ok := res[tmp]; !ok {
							res[tmp] = 1
						} else {
							res[tmp]++
						}
					}
				}
				tmp = ""
			}
		}
	}
	if tmp != "" {
		if strings.Count(tmp, "'") != len(tmp) {
			for strings.HasPrefix(tmp, "'") {
				tmp = tmp[1:]
			}
			for strings.HasSuffix(tmp, "'") {
				tmp = tmp[:len(tmp)-1]
			}
			if tmp != "" {
				tmp = strings.ToLower(tmp)
				if _, ok := res[tmp]; !ok {
					res[tmp] = 1
				} else {
					res[tmp]++
				}
			}
		}
	}
	return res
}
