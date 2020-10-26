package morestrings

//ReverseRunes Return reversed string input
func ReverseRunes(s string) string {
	var r = []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		var temp = r[i]
		r[i] = r[j]
		r[j] = temp
	}
	return string(r)
}
