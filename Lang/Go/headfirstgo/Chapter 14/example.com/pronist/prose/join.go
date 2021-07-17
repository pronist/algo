package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	var r string

	switch l := len(phrases); l {
	case 0:
		r = ""
	case 1:
		r = phrases[0]
	case 2:
		r = phrases[0] + " and " + phrases[1]
	default:
		r = strings.Join(phrases[:len(phrases)-1], ", ")
		r += ", and "
		r += phrases[len(phrases)-1]
	}
	return r
}
