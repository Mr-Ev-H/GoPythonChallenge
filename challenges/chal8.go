package challenges

import (
	"compress/bzip2"
	"fmt"
	"io"
	"strings"
)

// Read and uncompress bzip compressed strings
// http://www.pythonchallenge.com/pc/def/integrity.html
func Solve8() {

	un := "BZh91AY&SYA\xaf\x82\r\x00\x00\x01\x01\x80\x02\xc0\x02\x00 \x00!\x9ah3M\x07<]\xc9\x14\xe1BA\x06\xbe\x084"
	pw := "BZh91AY&SY\x94$|\x0e\x00\x00\x00\x81\x00\x03$ \x00!\x9ah3M\x13<]\xc9\x14\xe1BBP\x91\xf08"

	username := readbzip(un)
	password := readbzip(pw)

	fmt.Println("Next: http://www.pythonchallenge.com/pc/return/good.html") // This is just a link in the page

	fmt.Println("Username: " + username)
	fmt.Println("Password: " + password)

}

func readbzip(bzipIn string) string {
	stringReader := strings.NewReader(bzipIn)

	bzipReader := bzip2.NewReader(stringReader)

	buf := new(strings.Builder)

	_, err := io.Copy(buf, bzipReader)

	if err != nil {
		panic(err)
	}

	return buf.String()
}
