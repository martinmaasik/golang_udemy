package stringutil

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// pmst sisu on selles, et samas packages saab kasutada lowercase asju (Reverse kutsub reverseTwo mis on lowercase),
// samal ajal hiljem teises kaustas olev main.go hakkab kasutama Reverse-i mis pole samas paketis (kuigi Reverse omakorda kutsub reverseTwo)
