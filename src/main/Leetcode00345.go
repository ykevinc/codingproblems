package main

func reverseVowels(s string) string {
	vs := make([]int, 0, len(s))
	rs := []rune(s)
	for i, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vs = append(vs, i)
		}
		rs[i] = c
	}
	if len(vs) > 0 {
		for l, r := 0, len(vs)-1; l < r; l, r = l+1, r-1 {
			rs[vs[l]], rs[vs[r]] = rs[vs[r]], rs[vs[l]]
		}
	}
	return string(rs)
}
