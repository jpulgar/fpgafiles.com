package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var arcadeGameList []ArcadeGame
var arcadeAuthorList []string
var arcadeTitleAdded = make(map[string]bool)
var arcadeAuthorAdded = make(map[string]bool)
var moveList = make(map[string]string)

// XML Structure
type MRA_XML struct {
	XMLName xml.Name `xml:"misterromdescription"`
	SetName string   `xml:"setname"`
	RBF     string   `xml:"rbf"`
	Name    string   `xml:"name"`
	Year    string   `xml:"year"`
	About   struct {
		Author string `xml:"author,attr"`
	} `xml:"about"`
}

type ArcadeGame struct {
	Name    string
	Year    string
	SetName string
	Author  string
}

type ArcadeGamePageData struct {
	Image      string
	Name       string
	Year       string
	ListName   string
	FolderName string
	Author     string
	Moves      template.HTML
	Video      string
	Credit     template.HTML
}

func generateMisterArcadeGames(generate bool) {
	compileMisterArcadeData()
	if generate {
		generateMisterArcadeHTML()
	}
	copyArcadeImages()
}

func compileMisterArcadeData() {

	for _, f := range findAllFiles("assets/arcade/_Arcade", ".mra", "_alternatives") {
		out := MRA_XML{}
		if err := readMRA(f, &out); err != nil {
			fmt.Println(err)
			return
		}
		if out.SetName == "" {
			if out.RBF != "" {
				out.SetName = out.RBF
			}
			// MRA fixes
			if out.Name == "Clean Sweep" {
				out.SetName = "cleanswp"
			}
		}

		game := ArcadeGame{Name: out.Name, SetName: out.SetName, Year: out.Year, Author: out.About.Author}
		// Remove duplicates
		if _, ok := arcadeTitleAdded[out.SetName]; !ok {
			arcadeTitleAdded[out.SetName] = true
			arcadeGameList = append(arcadeGameList, game)
		}
	}

	// Sort arcadeGames by Name
	sort.Slice(arcadeGameList, func(i, j int) bool {
		return arcadeGameList[i].Name < arcadeGameList[j].Name
	})

	// Compile a list of authors
	for _, b := range arcadeGameList {
		if b.Author != "" {
			if _, ok := arcadeAuthorAdded[b.Author]; !ok {
				arcadeAuthorAdded[b.Author] = true
				arcadeAuthorList = append(arcadeAuthorList, b.Author)
			}
		}
	}

	// Write stats.json for Homepage Use
	videosFound := 0
	for _, v := range arcadeVideos {
		if v != "" {
			videosFound++
		}
	}
	videoPercentage := 0.00
	if videosFound != 0 {
		videoPercentage = float64(videosFound) / float64(len(arcadeGameList))
	}

	stats := Stats{nameForFolder("arcade"), len(arcadeGameList), float64(int(videoPercentage*100*100)) / 100, "arcade"}
	prettyJSON, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	WriteToFile("public/mister/arcade/stats.json", string(prettyJSON))

}

func readMRA(filename string, configObject *MRA_XML) error {

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

func readStatFile(filename string, statObject *Stats) error {
	configFilename := filename
	buf, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(buf), &statObject); err != nil {
		return err
	} else {
		return nil
	}
}

func generateMisterArcadeHTML() {

	var tmplBuffer bytes.Buffer

	// Generate num.html, a.html, b.html, c.html, ..., z.html, textlist.html
	listFilename := [28]string{"num", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "textlist"}
	for _, v := range listFilename {

		var tempGames []Game
		for _, g := range arcadeGameList {
			// Starting with letter
			if strings.ToLower(g.Name[0:1]) == v {
				temp := Game{g.SetName, g.SetName + ".png", arcadeVideos[g.SetName], g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
			// Starting with #
			if v == "num" {
				if _, err := strconv.Atoi(g.Name[0:1]); err == nil {
					temp := Game{g.SetName, g.SetName + ".png", arcadeVideos[g.SetName], g.Name + " (" + g.Year + ")"}
					tempGames = append(tempGames, temp)
				}
			}
			// Text List
			if v == "textlist" {
				temp := Game{g.SetName, g.SetName + ".png", arcadeVideos[g.SetName], g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
		}

		data := ListPageData{
			Sections: []Section{{"num", "#"}, {"a", "A"}, {"b", "B"}, {"c", "C"}, {"d", "D"}, {"e", "E"}, {"f", "F"},
				{"h", "H"}, {"i", "I"}, {"j", "J"}, {"k", "K"}, {"l", "L"}, {"m", "M"}, {"n", "N"}, {"o", "O"}, {"p", "P"},
				{"q", "Q"}, {"r", "R"}, {"s", "S"}, {"t", "T"}, {"u", "U"}, {"v", "V"}, {"w", "W"}, {"x", "X"}, {"y", "Y"},
				{"z", "Z"}, {"textlist", "Text List"}},
			CurrentPage: v,
			ListName:    "Arcade Games by Name",
			FolderName:  "arcade",
			Games:       tempGames,
			Credit:      template.HTML(getCredit(("arcade"))),
		}

		tmpl := template.Must(template.ParseFiles("list_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, data); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/arcade/"+strings.ToLower(v)+".html", tmplBuffer.String())
		tmplBuffer.Reset()

	}

	// Generate 1971.html, 1972.html, 1973.html, ..., 2012.html
	listYearFilename := [33]string{"1971", "1972", "1973", "1974", "1975", "1976", "1977", "1978", "1979",
		"1980", "1981", "1982", "1983", "1984", "1985", "1986", "1987", "1988", "1989",
		"1990", "1991", "1992", "1993", "1994", "1995", "1996", "1997", "1998", "1999",
		"2000", "2001", "2004", "2012"}
	for _, v := range listYearFilename {

		var tempGames []Game
		for _, g := range arcadeGameList {
			if g.Year == v {
				temp := Game{g.SetName, g.SetName + ".png", arcadeVideos[g.SetName], g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
		}

		data := ListPageData{
			Sections: []Section{{"1971", "71"}, {"1972", "72"}, {"1973", "73"}, {"1974", "74"}, {"1975", "75"},
				{"1976", "76"}, {"1977", "77"}, {"1978", "78"}, {"1979", "79"}, {"1980", "80"},
				{"1981", "81"}, {"1982", "82"}, {"1983", "83"}, {"1984", "84"}, {"1985", "85"},
				{"1986", "86"}, {"1987", "87"}, {"1988", "88"}, {"1989", "89"}, {"1990", "90"},
				{"1991", "91"}, {"1992", "92"}, {"1993", "93"}, {"1994", "94"}, {"1995", "95"},
				{"1996", "96"}, {"1997", "97"}, {"1998", "98"}, {"1999", "99"}, {"2000", "00"},
				{"2001", "01"}, {"2004", "04"}, {"2012", "12"}},
			CurrentPage: v,
			ListName:    "Arcade Games by Year",
			FolderName:  "arcade",
			Games:       tempGames,
			Credit:      template.HTML(getCredit(("arcade"))),
		}

		tmpl := template.Must(template.ParseFiles("list_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, data); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/arcade/"+strings.ToLower(v)+".html", tmplBuffer.String())
		tmplBuffer.Reset()

	}

	// Generate by Author (arcadeAuthorList)
	for _, v := range arcadeAuthorList {

		var tempGames []Game
		for _, g := range arcadeGameList {
			if g.Author == v {
				temp := Game{g.SetName, g.SetName + ".png", arcadeVideos[g.SetName], g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
		}

		var authorSections []Section
		for _, h := range arcadeAuthorList {
			authorSections = append(authorSections, Section{h, h})
		}

		data := ListPageData{
			Sections:    authorSections,
			CurrentPage: v,
			ListName:    "Arcade Games by Core Author",
			FolderName:  "arcade",
			Games:       tempGames,
			Credit:      template.HTML(getCredit(("arcade"))),
		}

		tmpl := template.Must(template.ParseFiles("list_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, data); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/arcade/"+strings.ToLower(v)+".html", tmplBuffer.String())
		tmplBuffer.Reset()

	}

	// Generate Arcade Games
	for _, v := range arcadeGameList {

		dataGames := ArcadeGamePageData{
			Image:      v.SetName + ".png",
			Name:       v.Name + " (" + v.Year + ")",
			FolderName: "arcade",
			Author:     v.Author,
			ListName:   "Arcade",

			Moves:  template.HTML(moveList[v.SetName]),
			Video:  arcadeVideos[v.SetName],
			Credit: template.HTML(getCredit(("arcade"))),
		}
		tmpl := template.Must(template.ParseFiles("game_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, dataGames); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/arcade/games/"+v.SetName+".html", tmplBuffer.String())
		tmplBuffer.Reset()
	}
}

func generateMisterArcadeCommands() {
	controlsStart := 0
	controlsEnd := controlsStart
	lines := LinesInFile("assets/arcade/command.dat")
	for index, line := range lines {
		if len(line) > 7 {
			if line[0:6] == "$info=" {
				commandsStart := index + 2
				commandsEnd := commandsStart
				inControls := false
				buttonNames := make(map[string]string)
				commonButtonNames := make(map[string]string)

				//fmt.Println("lines[" + strconv.Itoa(index) + "] = " + line[6:])
				sets := strings.Split(line[6:], ",")

				for j := commandsStart; lines[j] != "$end"; j++ {
					if !inControls && lines[j] == "- CONTROLS -" {
						controlsStart = j + 2
						inControls = true
					}
					if inControls && lines[j] == "" {
						controlsEnd = j - 1
						inControls = false
					}
					commandsEnd = j
				}

				// Define Game Specific Buttons
				for jj := controlsStart; jj <= controlsEnd; jj++ {
					if strings.Contains(lines[jj], " : ") {
						temp := strings.Split(lines[jj], " : ")
						temp[1] = strings.Replace(temp[1], " (_K)", "", -1)
						lines[jj] = strings.Replace(lines[jj], " (_K)", "", -1)
						temp[1] = strings.Replace(temp[1], " (_P)", "", -1)
						lines[jj] = strings.Replace(lines[jj], " (_P)", "", -1)
						temp[1] = strings.Replace(temp[1], " (^s)", "", -1)
						lines[jj] = strings.Replace(lines[jj], " (^s)", "", -1)
						buttonNames[temp[0]] = temp[1]
						lines[jj] = strings.Replace(lines[jj], temp[0]+" : ", "", -1)
					}
				}

				// Define Common Buttons
				commonButtonNames["_1"] = "↙️ "
				commonButtonNames["_2"] = "⬇️ "
				commonButtonNames["_3"] = "↘️ "
				commonButtonNames["_4"] = "⬅️ "
				commonButtonNames["_6"] = "➡️ "
				commonButtonNames["_7"] = "↖️ "
				commonButtonNames["_8"] = "⬆️ "
				commonButtonNames["_9"] = "↗️ "
				commonButtonNames["_x"] = "🔄 "
				commonButtonNames["_^"] = "In-Air"
				commonButtonNames[" / "] = " or "
				commonButtonNames["_O"] = "Hold "
				commonButtonNames["^1"] = "[Hold ↙️ ] "
				commonButtonNames["^2"] = "[Hold ⬇️ ] "
				commonButtonNames["^3"] = "[Hold ↘️ ] "
				commonButtonNames["^4"] = "[Hold ⬅️ ] "
				commonButtonNames["^6"] = "[Hold ➡️ ] "
				commonButtonNames["^7"] = "[Hold ↖️ ]"
				commonButtonNames["^8"] = "[Hold ⬆️ ]"
				commonButtonNames["^9"] = "[Hold ↗️ ]"
				commonButtonNames["_+"] = " + "
				commonButtonNames["^*"] = "Tap "
				commonButtonNames["_?"] = "Any Direction "
				commonButtonNames["_P"] = "Punch"
				commonButtonNames["_K"] = "Kick"
				commonButtonNames["^s"] = "Slash"
				commonButtonNames["^T"] = "3xKick"
				commonButtonNames["^U"] = "3xPunch"
				commonButtonNames["^V"] = "2xKick"
				commonButtonNames["^W"] = "2xPunch"

				for _, setname := range sets {
					//fmt.Println(setname)
					for j := commandsStart; j <= commandsEnd; j++ {
						if len(lines[j]) > 1 {

							// Game-specific buttons
							for name, value := range buttonNames {
								if lines[j][0:2] != name {
									lines[j] = strings.Replace(lines[j], name, value, -1)
								}
							}

							// Common buttons
							for name, value := range commonButtonNames {
								if lines[j][0:2] != name {
									lines[j] = strings.Replace(lines[j], name, value, -1)
								}
							}

						}
						lines[j] = strings.Replace(lines[j], "                          ", "<b style='color: #0486ff;'>", 1)
						lines[j] = strings.Replace(lines[j], "─────────────────────────────────────────────────────────────────────────", "<hr>", 1)
						moveList[setname] += lines[j] + "</b><br/>\n"
					}
				}

			}
		}
	}
}

// This is different than other platforms since we only want to copy what is supported
func copyArcadeImages() {

	for _, v := range arcadeGameList {

		// SNAPS
		sourceSize, err := FileSize("assets/arcade/snaps/" + v.SetName + ".png")
		if err != nil {
			fmt.Println(err)
		} else {
			fi, err := os.Stat("public/mister/arcade/snaps/" + v.SetName + ".png")
			if os.IsNotExist(err) {
				fmt.Println("Copying assets/arcade/snaps/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/snaps/"+v.SetName+".png", "public/mister/arcade/snaps/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			} else if sourceSize != fi.Size() {
				fmt.Println("Copying assets/arcade/snaps/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/snaps/"+v.SetName+".png", "public/mister/arcade/snaps/"+v.SetName+".png")
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
			fi, err := os.Stat("public/mister/arcade/titles/" + v.SetName + ".png")
			if os.IsNotExist(err) {
				fmt.Println("Copying assets/arcade/titles/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/titles/"+v.SetName+".png", "public/mister/arcade/titles/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			} else if sourceSize != fi.Size() {
				fmt.Println("Copying assets/arcade/titles/" + v.SetName + ".png")
				copyErr := CopyFile("assets/arcade/titles/"+v.SetName+".png", "public/mister/arcade/titles/"+v.SetName+".png")
				if copyErr != nil {
					fmt.Println(copyErr)
				}
			}
		}
	}

}
