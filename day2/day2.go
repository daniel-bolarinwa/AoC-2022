package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	X1 int = 1
	Y1     = 2
	Z1     = 3

	X2 int = 0
	Y2     = 3
	Z2     = 6
)

func solution1() error {
	var err error
	var file []byte

	if file, err = os.ReadFile("input.txt"); err != nil {
		return fmt.Errorf("error occured: %w", err)
	}

	fileString := string(file)
	competitionRounds := strings.Split(fileString, "\r\n")

	userScore := 0
	for _, value := range competitionRounds {
		oppUserDecisions := strings.Fields(value)

		if oppUserDecisions[0] == "A" && oppUserDecisions[1] == "Z" {
			userScore += (0 + Z1)
		}

		if oppUserDecisions[0] == "A" && oppUserDecisions[1] == "Y" {
			userScore += (6 + Y1)
		}

		if oppUserDecisions[0] == "B" && oppUserDecisions[1] == "X" {
			userScore += (0 + X1)
		}

		if oppUserDecisions[0] == "B" && oppUserDecisions[1] == "Z" {
			userScore += (6 + Z1)
		}

		if oppUserDecisions[0] == "C" && oppUserDecisions[1] == "Y" {
			userScore += (0 + Y1)
		}

		if oppUserDecisions[0] == "C" && oppUserDecisions[1] == "X" {
			userScore += (6 + X1)
		}

		if oppUserDecisions[0] == "C" && oppUserDecisions[1] == "Z" {
			userScore += (3 + Z1)
		}

		if oppUserDecisions[0] == "B" && oppUserDecisions[1] == "Y" {
			userScore += (3 + Y1)
		}

		if oppUserDecisions[0] == "A" && oppUserDecisions[1] == "X" {
			userScore += (3 + X1)
		}
	}
	fmt.Printf("day2 part1: %v \n", userScore)

	return nil
}

func solution2() error {
	var err error
	var file []byte

	if file, err = os.ReadFile("input.txt"); err != nil {
		return fmt.Errorf("error occured: %w", err)
	}

	fileString := string(file)
	competitionRounds := strings.Split(fileString, "\r\n")

	userScore := 0
	for _, value := range competitionRounds {
		oppUserDecisions := strings.Fields(value)

		if oppUserDecisions[0] == "A" && oppUserDecisions[1] == "Z" {
			userScore += (2 + Z2)
		}

		if oppUserDecisions[0] == "A" && oppUserDecisions[1] == "Y" {
			userScore += (1 + Y2)
		}

		if oppUserDecisions[0] == "A" && oppUserDecisions[1] == "X" {
			userScore += (3 + X2)
		}

		if oppUserDecisions[0] == "B" && oppUserDecisions[1] == "X" {
			userScore += (1 + X2)
		}

		if oppUserDecisions[0] == "B" && oppUserDecisions[1] == "Z" {
			userScore += (3 + Z2)
		}

		if oppUserDecisions[0] == "B" && oppUserDecisions[1] == "Y" {
			userScore += (2 + Y2)
		}

		if oppUserDecisions[0] == "C" && oppUserDecisions[1] == "Y" {
			userScore += (3 + Y2)
		}

		if oppUserDecisions[0] == "C" && oppUserDecisions[1] == "X" {
			userScore += (2 + X2)
		}

		if oppUserDecisions[0] == "C" && oppUserDecisions[1] == "Z" {
			userScore += (1 + Z2)
		}
	}
	fmt.Printf("day2 part2: %v", userScore)

	return nil
}

func main() {
	var err error
	if err = solution1(); err != nil {
		fmt.Printf("%v", err)
	}

	if err = solution2(); err != nil {
		fmt.Printf("%v", err)
	}
}
