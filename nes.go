package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

var nesnames = make(map[string]bool)
var nesimages = make(map[string]string)
var uniquenesnames []string

type NESListPageData struct {
	Sections    []Section
	CurrentPage string
	Games       []Game
}

type Game struct {
	Page  string
	Image string
	Name  string
}

type Section struct {
	Name  string
	Label string
}

type NESGamePageData struct {
	Name  string
	Image string
	Video string
}

func generateMisterNESGames() {
	generateMisterNESNamesJSON()
	generateMisterNESHTML()
	CopyNESScripts()
}

func generateMisterNESNamesJSON() {

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
		} else {
			if strings.Contains(original, "USA") {
				nesimages[f] = original
			}
		}
	}
}

func generateMisterNESHTML() {

	var tmplBuffer bytes.Buffer

	// Generate NES Games
	if true {
		for _, v := range uniquenesnames {

			nesVideo := nesVideos[v]

			dataGames := NESGamePageData{
				Name:  v,
				Image: nesimages[v],
				Video: nesVideo,
			}
			tmpl := template.Must(template.ParseFiles("mister/nes/game_layout.html", "navigation.html"))
			if err := tmpl.Execute(&tmplBuffer, dataGames); err != nil {
				fmt.Println(err)
			}
			WriteToFile("public/mister/nes/games/"+urlSafe(v)+".html", tmplBuffer.String())
			tmplBuffer.Reset()

		}
	}

	// Generate num.html, a.html, b.html, c.html, ..., z.html, textlist.html
	listFilename := [28]string{"num", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "textlist"}
	for _, v := range listFilename {

		var tempGames []Game
		for _, g := range uniquenesnames {
			// Starting with letter
			if strings.ToLower(g[0:1]) == v {
				temp := Game{urlSafe(g), nesimages[g], g}
				tempGames = append(tempGames, temp)
			}
			// Starting with #
			if v == "num" {
				if _, err := strconv.Atoi(g[0:1]); err == nil {
					temp := Game{urlSafe(g), nesimages[g], g}
					tempGames = append(tempGames, temp)
				}
			}
			// Text List
			if v == "textlist" {
				temp := Game{urlSafe(g), nesimages[g], g}
				tempGames = append(tempGames, temp)
			}
		}

		data := NESListPageData{
			Sections: []Section{{"num", "#"}, {"a", "A"}, {"b", "B"}, {"c", "C"}, {"d", "D"}, {"e", "E"}, {"f", "F"},
				{"h", "H"}, {"i", "I"}, {"j", "J"}, {"k", "K"}, {"l", "L"}, {"m", "M"}, {"n", "N"}, {"o", "O"}, {"p", "P"},
				{"q", "Q"}, {"r", "R"}, {"s", "S"}, {"t", "T"}, {"u", "U"}, {"v", "V"}, {"w", "W"}, {"x", "X"}, {"y", "Y"},
				{"z", "Z"}, {"textlist", "Text List"}},
			CurrentPage: v,
			Games:       tempGames,
		}

		tmpl := template.Must(template.ParseFiles("mister/nes/list_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, data); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/nes/"+strings.ToLower(v)+".html", tmplBuffer.String())
		tmplBuffer.Reset()

	}
}

func urlSafe(name string) string {
	name = strings.Replace(name, " - ", "-", -1)
	name = strings.Replace(name, " ", "-", -1)
	name = strings.Replace(name, "(", "", -1)
	name = strings.Replace(name, ")", "", -1)
	name = strings.Replace(name, ",", "", -1)
	name = strings.Replace(name, "'", "", -1)
	name = strings.Replace(name, ".", "", -1)
	name = strings.Replace(name, "&", "and", -1)
	name = strings.Replace(name, "!", "", -1)
	name = strings.Replace(name, "?", "", -1)
	name = strings.Replace(name, "$", "", -1)
	name = strings.Replace(name, "+", "and", -1)
	name = strings.Replace(name, "[", "", -1)
	name = strings.Replace(name, "]", "", -1)
	return name
}

func CopyNESScripts() {
	err := CopyFile("mister/nes/nes.css", "public/mister/nes/nes.css")
	if err != nil {
		fmt.Println(err)
		return
	}
}
