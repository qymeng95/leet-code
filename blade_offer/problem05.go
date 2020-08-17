package blade_offer

import "strings"

func ReplaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func TestReplaceSpace() {
	var s = "We are happy."
	res := ReplaceSpace(s)

	println(res)
}
