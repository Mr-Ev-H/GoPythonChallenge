package utility

import (
	"fmt"
	"strconv"
	"strings"
)

// Func taking an input entry id, returns id of next element
type ListNavigator func(string) string

// Given a walk function, traverses a set of nodes, returning the next id by basic parsing of text.
// Used by a number of challenges which is why its been extracted for reuse
func ListWalker(navigationMethod ListNavigator, startId int) string {
	var nextId int64 = int64(startId)
	nextBody := ""
	result := ""

	for result == "" {
		nextBody = navigationMethod(strconv.FormatInt(nextId, 10))
		fmt.Println(nextBody)
		var err error

		if nextBody == "Yes. Divide by two and keep going." {
			nextId = nextId / 2
		} else if strings.HasSuffix(nextBody, ".html") {
			result = nextBody
		} else if nextBody == "Collect the comments." {
			result = nextBody
		} else if nextBody == "that's it." {
			result = nextBody
		} else {
			words := strings.Split(nextBody, " ")

			lastWord := words[len(words)-1]

			nextId, err = strconv.ParseInt(lastWord, 0, 0)

			if err != nil {
				panic(err)
			}
		}
	}

	return result
}
