package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	listFile, err := os.Open("puzzleInput")
	if err != nil {
		log.Fatal("Failed to open file: ", err)
		return
	}

	fileScanner := bufio.NewScanner(listFile)

	var conditionCounter int

	for fileScanner.Scan() {
		var previousLevel int
		var levelTrend string
		var reportedLevels []int
		var failedChecks bool

		rawLine := strings.Split(fileScanner.Text(), " ")
		for _, entry := range rawLine {
			num, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatal("Expected int got something else: ", err)
			}
			reportedLevels = append(reportedLevels, num)
		}

		for index, reportedLevel := range reportedLevels {
			if index == 0 {
				previousLevel = reportedLevel
				switch {
				case previousLevel < reportedLevels[index+1]:
					levelTrend = "increasing"
				}
				continue
			}

			levelMovement := math.Abs(float64(previousLevel) - float64(reportedLevel))

			positiveLevelMovement := previousLevel < reportedLevel
			positiveLevelTrend := levelTrend == "increasing"
			isLevelMovementClamped := levelMovement != 0 && levelMovement <= 3

			overallPositive := positiveLevelTrend && positiveLevelMovement && isLevelMovementClamped
			overallNegative := !positiveLevelTrend && !positiveLevelMovement && isLevelMovementClamped

			switch {
			case overallPositive || overallNegative:
				previousLevel = reportedLevel
			default:
				failedChecks = true
				break
			}
		}

		switch {
		case failedChecks == true:
			break
		case failedChecks != true:
			conditionCounter++
		}
	}
	log.Println("Day 2 Task 1 flag: ", conditionCounter)
}
