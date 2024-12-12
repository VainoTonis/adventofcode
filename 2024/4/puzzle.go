package main

import (
	"bufio"
	"log"
	"os"
)

// Moon rune trasnlation
// X = 88
// M = 77
// A = 65
// S = 83

var matrixDirections = []struct {
	row, column int
}{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

var xmasCounter int

func main() {
	listFile, err := os.Open("puzzleInput")
	if err != nil {
		log.Fatal("Failed to open file: ", err)
		return
	}

	fileScanner := bufio.NewScanner(listFile)

	letterMatrix := [][]rune{}

	for fileScanner.Scan() {
		rowOfInput := fileScanner.Text()
		var temporaryRow []rune

		for _, characterRune := range rowOfInput {
			temporaryRow = append(temporaryRow, characterRune)
		}

		letterMatrix = append(letterMatrix, temporaryRow)

	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal("Error scanning the file:", err)
	}

	for rowIndex, row := range letterMatrix {
		for columnIndex, letter := range row {
			if letter == 88 {
				checkSurrondingRunes(letterMatrix, rowIndex, columnIndex)
			}
		}
	}

	log.Println("Part 1 flag:", xmasCounter)
}

func checkSurrondingRunes(letterMatrix [][]rune, row, column int) {
	for _, matrixDirection := range matrixDirections {

		rowToCheck := row + matrixDirection.row
		columnToCheck := column + matrixDirection.column

		if rowToCheck >= 0 && rowToCheck < len(letterMatrix) && columnToCheck >= 0 && columnToCheck < len(letterMatrix[0]) && letterMatrix[rowToCheck][columnToCheck] == 77 {
			rowToCheck := rowToCheck + matrixDirection.row
			columnToCheck := columnToCheck + matrixDirection.column

			if rowToCheck >= 0 && rowToCheck < len(letterMatrix) && columnToCheck >= 0 && columnToCheck < len(letterMatrix[0]) && letterMatrix[rowToCheck][columnToCheck] == 65 {
				rowToCheck := rowToCheck + matrixDirection.row
				columnToCheck := columnToCheck + matrixDirection.column

				if rowToCheck >= 0 && rowToCheck < len(letterMatrix) && columnToCheck >= 0 && columnToCheck < len(letterMatrix[0]) && letterMatrix[rowToCheck][columnToCheck] == 83 {
					xmasCounter++
				}
			}
		}
	}

}
