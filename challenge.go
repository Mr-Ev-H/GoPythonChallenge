package main

import (
	"fmt"
	"gopychallenge/challenges"
	"os"
)

func main() {

	var challenge string
	if len(os.Args) < 2 {
		fmt.Println("Must provide exactly 1 challenge to run")
		os.Exit(1)
	} else {
		challenge = os.Args[1]
	}

	switch challenge {
	case "0":
		challenges.Solve0()
	case "1":
		challenges.Solve1()
	case "2":
		challenges.Solve2()
	case "3":
		challenges.Solve3()
	case "3a":
		challenges.Solve3Dyn()
	case "4":
		challenges.Solve4()
	case "5":
		challenges.Solve5()
	case "6":
		challenges.Solve6()
	case "7":
		challenges.Solve7()
	case "8":
		challenges.Solve8()
	case "9":
		challenges.Solve9()
	case "10":
		challenges.Solve10()
	case "11":
		challenges.Solve11()
	case "12":
		challenges.Solve12()
	case "13":
		challenges.Solve13()
	case "14":
		challenges.Solve14()
	case "15":
		challenges.Solve15()
	case "16":
		challenges.Solve16()
	case "17":
		challenges.Solve17()
	case "18":
		challenges.Solve18()
	case "19":
		challenges.Solve19()
	case "20":
		challenges.Solve20()
	case "21":
		fallthrough
	case "22":
		fallthrough
	case "23":
		fallthrough
	case "24":
		fallthrough
	case "25":
		fallthrough
	case "26":
		fallthrough
	case "27":
		fallthrough
	case "28":
		fallthrough
	case "29":
		fallthrough
	case "30":
		fallthrough
	case "31":
		fallthrough
	case "32":
		fallthrough
	case "33":
		fmt.Printf("No solution yet.")
	default:
		fmt.Printf("Unknown challenge: %s \n", challenge)
	}
}
