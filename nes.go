package main

import (
	"fmt"
	"strconv"
	"strings"
)

// var arcadeSets []string
//var nesGameInfo = make(map[string]NESEntry)

// JSON Structure
// type NESEntry struct {
// 	Name  string `json:"name"`
// 	Image string `json:"image"`
// }

// // HTML Structure
// type IndexPageData struct {
// 	Alphabet1 []string
// 	Alphabet2 []string
// 	Alphabet3 []string
// 	Years     []string
// }

// type ArcadeGamePageData struct {
// 	ID     string
// 	Name   string
// 	Year   string
// 	Author string
// 	Moves  template.HTML
// 	Video  string
// }

func generateMisterNESGames() {
	generateMisterNESNamesJSON()
}

func generateMisterNESNamesJSON() {
	//nesentries := []NESEntry{}
	nesnames := make(map[string]bool)
	nesimages := make(map[string]string)
	var uniquenesnames []string

	for _, f := range findAllFiles("assets/nes/titles", ".png", "") {

		original := f
		original = strings.Replace(original, "assets/nes/titles/", "", 1)
		f = strings.Replace(f, "assets/nes/titles/", "", 1)
		f = strings.Replace(f, ".png", "", 1)
		if idx := strings.IndexByte(f, '('); idx >= 0 {
			f = strings.TrimRight(f[:idx], " ")
		}
		if idx := strings.IndexByte(f, '['); idx >= 0 {
			f = strings.TrimRight(f[:idx], " ")
		}
		f = strings.TrimRight(f, " ")

		// Only add unique titles once
		if _, ok := nesnames[f]; !ok {
			nesnames[f] = true
			uniquenesnames = append(uniquenesnames, f)
			nesimages[f] = original
			// nesentry := NESEntry{Name: f, Image: original}
			// nesentries = append(nesentries, nesentry)
		} else {
			if strings.Contains(original, "USA") {
				nesimages[f] = original
			}
		}
	}

	// for _, v := range uniquenesnames {
	//fmt.Println(v + " --> " + nesimages[v]) // Show All

	// if strings.Contains(nesimages[v], "USA") { // Show USA released games
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	// if strings.Contains(nesimages[v], "(Unl)") { // Show Unlicensed only
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	// if strings.Contains(nesimages[v], "(Australia)") { // Show Australia only
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	// if strings.Contains(nesimages[v], "(Asia)") { // Show Asia only
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	// if strings.Contains(nesimages[v], "(Korea)") { // Show Korea only
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	// if strings.Contains(nesimages[v], "(Europe)") { // Show Europe only
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	// if strings.Contains(nesimages[v], "(Japan)") { // Show Europe only
	// 	fmt.Println(v + " --> " + nesimages[v])
	// }
	//}

	if false {
		return
	}

	// Calculate Best Video Matches
	var distZero = 0
	var distOne = 0
	var distTwo = 0
	var distThree = 0
	var distFour = 0
	var distFive = 0
	var distMoreThanFive = 0
	for _, v := range uniquenesnames {
		// Only do this process for titles with no video
		//if arcadeVideos[arcadeGameInfo[i].SetName] == "" {

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
		for _, n := range nesLongplays {
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

		// List already contains <= 2
		// Now let's do == 3, etc
		if lowestDistance == 3 {

			// TODO: Only fix these
			if strings.Contains(nesimages[v], "USA") {
				if !strings.Contains(nesimages[v], "(Unl)") {
					lowestName = lowestName[strings.IndexByte(lowestName, '|'):]
					fmt.Println("\"" + v + "\": \"" + lowestName + "\",")
				}
			}

		}
		//}
	}
	// Reporting
	fmt.Println("Distance 0: " + strconv.Itoa(distZero))
	fmt.Println("Distance 1: " + strconv.Itoa(distOne))
	fmt.Println("Distance 2: " + strconv.Itoa(distTwo))
	fmt.Println("Distance 3: " + strconv.Itoa(distThree))
	fmt.Println("Distance 4: " + strconv.Itoa(distFour))
	fmt.Println("Distance 5: " + strconv.Itoa(distFive))
	fmt.Println("Distance 5+: " + strconv.Itoa(distMoreThanFive))

	fmt.Println("Unique NES games: " + strconv.Itoa(len(uniquenesnames)))
	// Sort entries
	// sort.Slice(entries, func(i, j int) bool {
	// 	return entries[i].Name < entries[j].Name
	// })

	// prettyJSON, err := json.MarshalIndent(entries, "", "    ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// WriteToFile("mister/arcade/name.json", string(prettyJSON))
	// WriteToFile("public/mister/arcade/name.json", string(prettyJSON))
}
