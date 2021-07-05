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
	"strings"
)

var arcadeSets []string
var moveList = make(map[string]string)
var arcadeGameInfo = make(map[string]Entry)

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

// JSON Structure
type Entry struct {
	Name    string `json:"name"`
	Year    string `json:"year"`
	SetName string `json:"setname"`
	Author  string `json:"author"`
}

// HTML Structure
type IndexPageData struct {
	Alphabet1 []string
	Alphabet2 []string
	Alphabet3 []string
	Years     []string
}

type ArcadeGamePageData struct {
	ID     string
	Name   string
	Year   string
	Author string
	Moves  template.HTML
}

func generateMisterArcadeGames() {

	generateMisterArcadeCommands()
	generateMisterArcadeNamesJSON()
	generateMisterArcadeHTML()
	CopyArcadeScripts()
	copyArcadeImages()
	// fmt.Println(arcadeSets)
	//fmt.Println(moveList["mk3"])
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

func generateMisterArcadeHTML() {

	var tmplBuffer bytes.Buffer

	// Generate index.html and jotego.html
	data := IndexPageData{
		Alphabet1: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
		Alphabet2: []string{"K", "L", "M", "N", "O", "P", "Q", "R"},
		Alphabet3: []string{"S", "T", "U", "V", "W", "X", "Y", "Z"},
		Years: []string{"1971", "1972", "1973", "1974", "1975", "1976", "1977", "1978",
			"1979", "1980", "1981", "1982", "1983", "1984", "1985", "1986", "1987", "1988",
			"1989", "1990", "1991", "1992", "1993", "1994", "1995", "1996", "1997", "1998",
			"1999", "2000", "2001"},
	}

	tmpl := template.Must(template.ParseFiles("mister/arcade/index_layout.html", "navigation.html"))
	if err := tmpl.Execute(&tmplBuffer, data); err != nil {
		fmt.Println(err)
	}
	WriteToFile("public/mister/arcade/index.html", tmplBuffer.String())
	tmplBuffer.Reset()

	// Generate Where to Buy
	tmpl = template.Must(template.ParseFiles("mister/where-to-buy_layout.html", "navigation.html"))
	if err := tmpl.Execute(&tmplBuffer, nil); err != nil {
		fmt.Println(err)
	}
	WriteToFile("public/mister/where-to-buy.html", tmplBuffer.String())
	tmplBuffer.Reset()

	// Generate Arcade Games
	for _, v := range arcadeSets {
		dataGames := ArcadeGamePageData{
			ID:     v,
			Name:   arcadeGameInfo[v].Name,
			Year:   arcadeGameInfo[v].Year,
			Author: arcadeGameInfo[v].Author,
			Moves:  template.HTML(moveList[v]),
		}
		tmpl = template.Must(template.ParseFiles("mister/arcade/arcade_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, dataGames); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/arcade/games/"+v+".html", tmplBuffer.String())
		tmplBuffer.Reset()
	}
}

func generateMisterArcadeNamesJSON() {
	entries := []Entry{}
	setnames := make(map[string]bool)

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
		}
		entry := Entry{Name: out.Name, SetName: out.SetName, Year: out.Year, Author: out.About.Author}
		// Remove duplicates
		if _, ok := setnames[out.SetName]; !ok {
			setnames[out.SetName] = true
			entries = append(entries, entry)
			arcadeSets = append(arcadeSets, entry.SetName)
			arcadeGameInfo[out.SetName] = Entry{Name: out.Name, SetName: out.SetName, Year: out.Year, Author: out.About.Author}
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	prettyJSON, err := json.MarshalIndent(entries, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	WriteToFile("mister/arcade/name.json", string(prettyJSON))
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

func CopyArcadeScripts() {
	err := CopyFile("mister/arcade/arcade.js", "public/mister/arcade/arcade.js")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = CopyFile("mister/arcade/arcade.css", "public/mister/arcade/arcade.css")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func copyArcadeImages() {

	for _, v := range arcadeSets {

		// SNAPS
		sourceSize, err := FileSize("assets/arcade/snaps/" + v + ".png")
		if err != nil {
			fmt.Println(err)
		} else {
			fi, err := os.Stat("public/mister/arcade/snaps/" + v + ".png")
			if sourceSize != fi.Size() {
				fmt.Println("Copying assets/arcade/snaps/" + v + ".png")
				if os.IsNotExist(err) {
					copyErr := CopyFile("assets/arcade/snaps/"+v+".png", "public/mister/arcade/snaps/"+v+".png")
					if copyErr != nil {
						fmt.Println(copyErr)
					}
				}
			}
		}

		// TITLES
		sourceSize, err = FileSize("assets/arcade/titles/" + v + ".png")
		if err != nil {
			fmt.Println(err)
		} else {
			fi, err := os.Stat("public/mister/arcade/titles/" + v + ".png")
			if sourceSize != fi.Size() {
				fmt.Println("Copying assets/arcade/titles/" + v + ".png")
				if os.IsNotExist(err) {
					copyErr := CopyFile("assets/arcade/titles/"+v+".png", "public/mister/arcade/titles/"+v+".png")
					if copyErr != nil {
						fmt.Println(copyErr)
					}
				}
			}
		}
	}

}
