package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"os"
)

func solution() error {
	var fileString string
	var calorieTotals []int
	var err error
	var file []byte

	if file, err = os.ReadFile("input.txt"); err != nil {
		return fmt.Errorf("error occured: %w", err)
	}

	fileString = string(file)

	elves := strings.Split(fileString, "\r\n\r\n")

	for _, value := range elves {
		elfTotalCalories := 0

		elfCalories := strings.Split(value, "\r\n")

		for _, value := range elfCalories {
			stringConverted, _ := strconv.Atoi(value)

			elfTotalCalories += stringConverted
		}

		calorieTotals = append(calorieTotals, elfTotalCalories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calorieTotals)))

	fmt.Printf("part 1: %v \n", calorieTotals[0])
	fmt.Printf("part 2: %v", calorieTotals[0] + calorieTotals[1] + calorieTotals[2])

	return nil
}

func main() {
	if err := solution(); err != nil {
		fmt.Printf("%v", err)
	}
}
