package main

import (
	"fmt"
	"strings"

	"os"
)

func getMap() map[string]int {
    return map[string]int {
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
}

func solution1() error {
	var err error
	var file []byte
	var repeatingItems []string
	priorityTotal := 0

	itemPriorityWeights := getMap()

	if file, err = os.ReadFile("input.txt"); err != nil {
		return fmt.Errorf("error occured: %w", err)
	}

	fileString := string(file)
	rucksacks := strings.Split(fileString, "\r\n")

	for _, value := range rucksacks {
		rucksackSize := len(value)
		
		firstHalf := value[:rucksackSize/2]
		secondHalf := value[rucksackSize/2:]

		valueIdentified := false
		for _, value1 := range []rune(firstHalf) {
			for _, value2 := range []rune(secondHalf) {
				if valueIdentified == true {
					break
				}

				if string(value1) == string(value2) {
					valueIdentified = true
					repeatingItems = append(repeatingItems, string(value1))
				}
			}
		}
	}
	
	for _, value := range repeatingItems {
		priorityTotal += itemPriorityWeights[value]
	}

	fmt.Printf("total p1: %v", priorityTotal)

	return nil
}

//TBD
func solution2() error {
	var err error
	var file []byte
	var repeatingItems []string
	priorityTotal := 0

	itemPriorityWeights := getMap()

	if file, err = os.ReadFile("input.txt"); err != nil {
		return fmt.Errorf("error occured: %w", err)
	}

	fileString := string(file)
	rucksacks := strings.Split(fileString, "\r\n")

	for _, value := range rucksacks {
		rucksackSize := len(value)
		
		firstHalf := value[:rucksackSize/2]
		secondHalf := value[rucksackSize/2:]

		valueIdentified := false
		for _, value1 := range []rune(firstHalf) {
			for _, value2 := range []rune(secondHalf) {
				if valueIdentified == true {
					break
				}

				if string(value1) == string(value2) {
					valueIdentified = true
					repeatingItems = append(repeatingItems, string(value1))
				}
			}
		}
	}
	
	for _, value := range repeatingItems {
		priorityTotal += itemPriorityWeights[value]
	}

	fmt.Printf("total p2: %v", priorityTotal)

	return nil
}

func main()  {
	if err := solution1(); err != nil {
		fmt.Printf("%v", err)
	}

	if err := solution2(); err != nil {
		fmt.Printf("%v", err)
	}
}
