package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateMisterPCEGames(generate bool) {
	pceTitleAdded := make(map[string]bool)
	pceImages := make(map[string]string)
	pceGameList := []string{}
	compileMisterConsoleData(pceTitleAdded, &pceGameList, pceImages, pceVideos, "pce")
	pceVideoReport(&pceGameList)
	if generate {
		generateMisterConsoleHTML("TurboGrafx 16 / PC Engine Games", &pceGameList, pceImages, pceVideos, "pce", "pc engine")
	}
}

var pceVideos = map[string]string{}

func pceVideoReport(gameList *[]string) {
	// Calculate Best Video Matches
	var distZero = 0
	var distOne = 0
	var distTwo = 0
	var distThree = 0
	var distFour = 0
	var distFive = 0
	var distMoreThanFive = 0
	for _, v := range *gameList {
		// Only do this process for titles with no video
		if pceVideos[v] == "" {

			tempName := v

			if idx := strings.IndexByte(tempName, '('); idx >= 0 {
				tempName = strings.TrimRight(tempName[:idx], " ")
			}
			if strings.Contains(tempName, ", The") {
				tempName = strings.Replace(tempName, ", The", "", 1)
				tempName = "The " + tempName
			}

			var str1 = []rune(tempName)
			var lowestDistance = 99
			var lowestName = ""
			for _, n := range pceLongplays {
				temptemp := n[:strings.IndexByte(n, '|')]
				var str2 = []rune(temptemp)
				var tempDistance = levenshtein(str1, str2)
				if tempDistance < lowestDistance {
					lowestDistance = tempDistance
					lowestName = n //temptemp
				}
			}
			if lowestDistance == 0 {
				distZero = distZero + 1
			}
			if lowestDistance == 1 {
				distOne = distOne + 1
			}
			if lowestDistance == 2 {
				distTwo = distTwo + 1
			}
			if lowestDistance == 3 {
				distThree = distThree + 1
			}
			if lowestDistance == 4 {
				distFour = distFour + 1
			}
			if lowestDistance == 5 {
				distFive = distFive + 1
			}
			if lowestDistance > 5 {
				distMoreThanFive = distMoreThanFive + 1
			}

			// Cycle through: <= 2, == 3, == 4, == 5, > 5
			//  are correct
			if lowestDistance == 0 {

				//lowestName = lowestName[strings.IndexByte(lowestName, '|')+1:] // for final
				fmt.Println("\"" + v + "\": \"" + lowestName + "\",") // for testing

				// For the rest without matches
				//fmt.Println("\"" + v + "\": \"\",")
				if false {
					fmt.Println(lowestName)
				}
			}
		}
	}
	// Reporting
	fmt.Println("Distance 0: " + strconv.Itoa(distZero))
	fmt.Println("Distance 1: " + strconv.Itoa(distOne))
	fmt.Println("Distance 2: " + strconv.Itoa(distTwo))
	fmt.Println("Distance 3: " + strconv.Itoa(distThree))
	fmt.Println("Distance 4: " + strconv.Itoa(distFour))
	fmt.Println("Distance 5: " + strconv.Itoa(distFive))
	fmt.Println("Distance 5+: " + strconv.Itoa(distMoreThanFive))
}

var pceLongplays = []string{}
