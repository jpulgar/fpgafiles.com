package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var neogeoGameList []ArcadeGame

// XML Structure
type ROMSETS_XML struct {
	XMLName xml.Name `xml:"romsets"`
	Romset  []struct {
		SetName string `xml:"name,attr"`
		Name    string `xml:"altname,attr"`
		Year    string `xml:"year,attr"`
	} `xml:"romset"`
}

func generateMisterNeoGeoGames() {
	compileMisterNeoGeoData()
	generateMisterNeoGeoHTML()
	copyNeoGeoImages()
}

func compileMisterNeoGeoData() {

	out := ROMSETS_XML{}
	if err := readNeoGeoRomsets("assets/arcade/neogeo-romsets.xml", &out); err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range out.Romset {
		game := ArcadeGame{Name: v.Name, SetName: strings.Split(v.SetName, ",")[0], Year: v.Year}
		neogeoGameList = append(neogeoGameList, game)
	}
	// Sort neogeoGameList by Name
	sort.Slice(neogeoGameList, func(i, j int) bool {
		return neogeoGameList[i].Name < neogeoGameList[j].Name
	})

}

func readNeoGeoRomsets(filename string, configObject *ROMSETS_XML) error {

	configFilename := filename
	buf, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal([]byte(buf), &configObject); err != nil {
		return err
	} else {
		return nil
	}
}

func generateMisterNeoGeoHTML() {

	var tmplBuffer bytes.Buffer

	// Generate num.html, a.html, b.html, c.html, ..., z.html, textlist.html
	listFilename := [28]string{"num", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "textlist"}
	for _, v := range listFilename {

		var tempGames []Game
		for _, g := range neogeoGameList {
			// Starting with letter
			if strings.ToLower(g.Name[0:1]) == v {
				temp := Game{g.SetName, g.SetName + ".png", g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
			// Starting with #
			if v == "num" {
				if _, err := strconv.Atoi(g.Name[0:1]); err == nil {
					temp := Game{g.SetName, g.SetName + ".png", g.Name + " (" + g.Year + ")"}
					tempGames = append(tempGames, temp)
				}
			}
			// Text List
			if v == "textlist" {
				temp := Game{g.SetName, g.SetName + ".png", g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
		}

		data := ListPageData{
			Sections: []Section{{"num", "#"}, {"a", "A"}, {"b", "B"}, {"c", "C"}, {"d", "D"}, {"e", "E"}, {"f", "F"},
				{"h", "H"}, {"i", "I"}, {"j", "J"}, {"k", "K"}, {"l", "L"}, {"m", "M"}, {"n", "N"}, {"o", "O"}, {"p", "P"},
				{"q", "Q"}, {"r", "R"}, {"s", "S"}, {"t", "T"}, {"u", "U"}, {"v", "V"}, {"w", "W"}, {"x", "X"}, {"y", "Y"},
				{"z", "Z"}, {"textlist", "Text List"}},
			CurrentPage: v,
			ListName:    "NEO•GEO Games",
			FolderName:  "neogeo",
			Games:       tempGames,
			Credit:      template.HTML(getCredit(("neogeo"))),
		}

		tmpl := template.Must(template.ParseFiles("list_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, data); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/neogeo/"+strings.ToLower(v)+".html", tmplBuffer.String())
		tmplBuffer.Reset()

	}

	// Generate 1971.html, 1972.html, 1973.html, ..., 2012.html
	// listYearFilename := [33]string{"1971", "1972", "1973", "1974", "1975", "1976", "1977", "1978", "1979",
	// 	"1980", "1981", "1982", "1983", "1984", "1985", "1986", "1987", "1988", "1989",
	// 	"1990", "1991", "1992", "1993", "1994", "1995", "1996", "1997", "1998", "1999",
	// 	"2000", "2001", "2004", "2012"}
	// for _, v := range listYearFilename {

	// 	var tempGames []Game
	// 	for _, g := range neogeoGameList {
	// 		if g.Year == v {
	// 			temp := Game{g.SetName, g.SetName + ".png", g.Name + " (" + g.Year + ")"}
	// 			tempGames = append(tempGames, temp)
	// 		}
	// 	}

	// 	data := ListPageData{
	// 		Sections: []Section{{"1971", "71"}, {"1972", "72"}, {"1973", "73"}, {"1974", "74"}, {"1975", "75"},
	// 			{"1976", "76"}, {"1977", "77"}, {"1978", "78"}, {"1979", "79"}, {"1980", "80"},
	// 			{"1981", "81"}, {"1982", "82"}, {"1983", "83"}, {"1984", "84"}, {"1985", "85"},
	// 			{"1986", "86"}, {"1987", "87"}, {"1988", "88"}, {"1989", "89"}, {"1990", "90"},
	// 			{"1991", "91"}, {"1992", "92"}, {"1993", "93"}, {"1994", "94"}, {"1995", "95"},
	// 			{"1996", "96"}, {"1997", "97"}, {"1998", "98"}, {"1999", "99"}, {"2000", "00"},
	// 			{"2001", "01"}, {"2004", "04"}, {"2012", "12"}},
	// 		CurrentPage: v,
	// 		ListName:    "NEO GEO Games by Year",
	// 		FolderName:  "neogeo",
	// 		Games:       tempGames,
	// 		Credit:      template.HTML(getCredit(("neogeo"))),
	// 	}

	// 	tmpl := template.Must(template.ParseFiles("list_layout.html", "navigation.html"))
	// 	if err := tmpl.Execute(&tmplBuffer, data); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	WriteToFile("public/mister/neogeo/"+strings.ToLower(v)+".html", tmplBuffer.String())
	// 	tmplBuffer.Reset()

	// }

	// Generate Neo Geo Games
	for _, v := range neogeoGameList {

		dataGames := ArcadeGamePageData{
			Image:      v.SetName + ".png",
			Name:       v.Name + " (" + v.Year + ")",
			FolderName: "neogeo",
			ListName:   "NEO GEO",

			Moves:  template.HTML(moveList[v.SetName]),
			Video:  arcadeVideos[v.SetName],
			Credit: template.HTML(getCredit(("neogeo"))),
		}
		tmpl := template.Must(template.ParseFiles("game_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, dataGames); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/neogeo/games/"+v.SetName+".html", tmplBuffer.String())
		tmplBuffer.Reset()
	}
}

func copyNeoGeoImages() {

	for _, v := range neogeoGameList {

		// SNAPS
		sourceSize, err := FileSize("assets/arcade/snaps/" + v.SetName + ".png")
		if err != nil {
			fmt.Println(err)
		} else {
			fi, err := os.Stat("public/mister/neogeo/snaps/" + v.SetName + ".png")
			if os.IsNotExist(err) {
				fmt.Println("Copying assets/arcade/snaps/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/snaps/"+v.SetName+".png", "public/mister/neogeo/snaps/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			} else if sourceSize != fi.Size() {
				fmt.Println("Copying assets/arcade/snaps/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/snaps/"+v.SetName+".png", "public/mister/neogeo/snaps/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			}
		}

		// TITLES
		sourceSize, err = FileSize("assets/arcade/titles/" + v.SetName + ".png")
		if err != nil {
			fmt.Println(err)
		} else {
			fi, err := os.Stat("public/mister/neogeo/titles/" + v.SetName + ".png")
			if os.IsNotExist(err) {
				fmt.Println("Copying assets/arcade/titles/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/titles/"+v.SetName+".png", "public/mister/neogeo/titles/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			} else if sourceSize != fi.Size() {
				fmt.Println("Copying assets/arcade/titles/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/titles/"+v.SetName+".png", "public/mister/neogeo/titles/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			}
		}
	}

}
