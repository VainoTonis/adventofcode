package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	listFile, err := os.Open("puzzleInput")
	if err != nil {
		log.Fatal("Failed to open file: ", err)
		return
	}

	fileScanner := bufio.NewScanner(listFile)

	var multiplicationResults int

	for fileScanner.Scan() {
		filteredInput, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
		if err != nil {
			log.Fatal(err)
		}

		numberFilter, err := regexp.Compile(`[0-9]+`)
		if err != nil {
			log.Fatal(err)
		}

		uncorruptedInstructions := filteredInput.FindAllString(fileScanner.Text(), -1)
		for _, uncorruptedInstruction := range uncorruptedInstructions {
			rawMultiplicationInputs := numberFilter.FindAllString(uncorruptedInstruction, -1)
			if len(rawMultiplicationInputs) != 2 {
				log.Fatal("Expected 2 numbers for multiplication")
			}

			var multiplicationInput []int

			for _, rawMultiplicationInput := range rawMultiplicationInputs {
				inputInt, err := strconv.Atoi(rawMultiplicationInput)

				if err != nil {
					log.Fatal("Only INT should be forwarded:", err)
				}

				multiplicationInput = append(multiplicationInput, inputInt)
			}

			multiplicationResults = multiplicationResults + (multiplicationInput[0] * multiplicationInput[1])
		}
	}

	log.Println("Day 3 flag:", multiplicationResults)
}
