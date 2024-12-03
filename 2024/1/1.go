package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	listFile, err := os.Open("list")
	if err != nil {
		log.Fatal("Failed to open file: ", err)
		return
	}

	fileScanner := bufio.NewScanner(listFile)
	fileScanner.Split(bufio.ScanWords)

	leftRight := 0

	var leftColumn, rightColumn []int

	for fileScanner.Scan() {
		number, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if leftRight == 0 {
			leftColumn = append(leftColumn, number)
			leftRight = 1
		} else if leftRight == 1 {
			rightColumn = append(rightColumn, number)
			leftRight = 0
		}

	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal("Failure during scanning: ", err)
	}

	sort.Ints(leftColumn)
	fmt.Println(leftColumn)

	sort.Ints(rightColumn)
	fmt.Println(rightColumn)

	var taskFlag float64

	for i := 0; i < len(rightColumn); i++ {
		columnDiff := leftColumn[i] - rightColumn[i]
		taskFlag = taskFlag + math.Abs(float64(columnDiff))
	}

	fmt.Println("Part 1 answer: ", taskFlag)
}
