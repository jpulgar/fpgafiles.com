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
