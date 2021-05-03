package challenges

import (
	"fmt"
	"strings"

	"alexejk.io/go-xmlrpc"
)

// More of a puzzle than a programming challenge - You have to know that lovable Sesame Street character bert is evil (It was a thing back in the 2000's)
// Then you have to perform XMLRPC service discovery on the phonebook/php to get the method call to give you the answer
// This is 2021, XML RPC feels archaic.
// http://www.pythonchallenge.com/pc/return/disproportion.html
func Solve13() {
	answer, _ := LookupNumber("Bert")

	answer = strings.Trim(answer, "5-")
	answer = strings.ToLower(answer)

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/return/%s.html", answer)
	fmt.Println("Next: " + url)
}

func LookupNumber(name string) (string, error) {
	client, _ := xmlrpc.NewClient("http://www.pythonchallenge.com/pc/phonebook.php")

	request :=
		&struct {
			Name string
		}{
			Name: name,
		}

	result := &struct {
		PhoneNumber string
	}{
		PhoneNumber: "",
	}

	defer client.Close()

	err := client.Call("phone", request, result)

	return result.PhoneNumber, err
}
