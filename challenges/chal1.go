package challenges

import (
	"fmt"
	"strings"
)

// rot2 cipher
// http://www.pythonchallenge.com/pc/def/map.html
func Solve1() {
	cyphertext := "g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj."

	rot2 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+2)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+2)%26
		}
		return r
	}

	fmt.Println(strings.Map(rot2, cyphertext))

	cypherPath := "map"

	answer := strings.Map(rot2, cypherPath)

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%s.html", answer)

	fmt.Println("Next: " + url)
}
