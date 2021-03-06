package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Used to Generate List Pages
type ListPageData struct {
	Sections    []Section
	CurrentPage string
	Games       []Game
	ListName    string
	FolderName  string
	PageTitle   string
	Credit      template.HTML
}

type Game struct {
	Page          string
	Image         string
	Video         string
	Name          string
	YouTubeSearch string
}

type Section struct {
	Name  string
	Label string
}

// Used for Console Individual Game Pages
type ConsoleGamePageData struct {
	Name       string
	Image      string
	Video      string
	ListName   string
	FolderName string
	Credit     template.HTML
}

// Used for stats
type Stats struct {
	Name               string
	Games              int
	LongplayPercentage float64
	Link               string
}

func compileMisterConsoleData(titleAdded map[string]bool, gameList *[]string, images map[string]string, videos map[string]string, folderName string) {

	for _, f := range findAllFiles("public/mister/"+folderName+"/snaps", ".png", "") {

		original := f
		original = strings.Replace(original, "public/mister/"+folderName+"/snaps/", "", 1)
		f = strings.Replace(f, "public/mister/"+folderName+"/snaps/", "", 1)
		f = strings.Replace(f, ".png", "", 1)
		if idx := strings.IndexByte(f, '('); idx >= 0 {
			f = strings.TrimRight(f[:idx], " ")
		}
		if idx := strings.IndexByte(f, '['); idx >= 0 {
			f = strings.TrimRight(f[:idx], " ")
		}
		f = strings.TrimRight(f, " ")

		// Only add unique titles once
		if _, ok := titleAdded[f]; !ok {
			titleAdded[f] = true
			*gameList = append(*gameList, f)
			images[f] = original
		} else {
			if strings.Contains(original, "USA") {
				images[f] = original
			}
		}
	}

	// Write stats.json for Homepage Use
	videosFound := 0
	for _, v := range videos {
		if v != "" {
			videosFound++
		}
	}
	videoPercentage := 0.00
	if videosFound != 0 {
		videoPercentage = float64(videosFound) / float64(len(*gameList))
	}

	stats := Stats{nameForFolder(folderName), len(*gameList), float64(int(videoPercentage*100*100)) / 100, folderName}
	prettyJSON, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	WriteToFile("public/mister/"+folderName+"/stats.json", string(prettyJSON))
}

func nameForFolder(folder string) string {
	if folder == "arcade" {
		return "Arcade"
	} else if folder == "neogeo" {
		return "Neo Geo"
	} else if folder == "atari2600" {
		return "Atari 2600"
	} else if folder == "astrocade" {
		return "Astrocade"
	} else if folder == "odyssey2" {
		return "Odyssey2"
	} else if folder == "coleco" {
		return "ColecoVision"
	} else if folder == "lynx" {
		return "Lynx"
	} else if folder == "pce" {
		return "TurboGrafx 16 / PC Engine"
	} else if folder == "pcecd" {
		return "TurboGrafx 16 / PC Engine CD"
	} else if folder == "supergrafx" {
		return "SuperGrafx"
	} else if folder == "nes" {
		return "NES"
	} else if folder == "gb" {
		return "Game Boy"
	} else if folder == "gbc" {
		return "Game Boy Color"
	} else if folder == "gba" {
		return "Game Boy Advance"
	} else if folder == "snes" {
		return "Super Nintendo"
	} else if folder == "sms" {
		return "Sega Master System"
	} else if folder == "sg1000" {
		return "Sega SG-1000"
	} else if folder == "gamegear" {
		return "Game Gear"
	} else if folder == "genesis" {
		return "Genesis / Mega Drive"
	} else if folder == "segacd" {
		return "Sega CD / Mega-CD"
	} else if folder == "vectrex" {
		return "Vectrex"
	} else if folder == "ws" {
		return "WonderSwan"
	} else if folder == "wsc" {
		return "WonderSwan Color"
	}
	return ""
}

func generateMisterConsoleHTML(listName string, gameList *[]string, images map[string]string, videos map[string]string, folderName string, youtubesearch string) {

	var tmplBuffer bytes.Buffer

	// Generate Individual Games
	for _, v := range *gameList {

		video := videos[v]

		dataGames := ConsoleGamePageData{
			Name:       v,
			Image:      images[v],
			Video:      video,
			ListName:   listName,
			FolderName: folderName,
			Credit:     template.HTML(getCredit(folderName)),
		}
		tmpl := template.Must(template.ParseFiles("game_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, dataGames); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/"+folderName+"/games/"+urlSafe(v)+".html", tmplBuffer.String())
		tmplBuffer.Reset()

	}

	// Generate num.html, a.html, b.html, c.html, ..., z.html, textlist.html
	listFilename := [28]string{"num", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "textlist"}
	for _, v := range listFilename {

		var tempGames []Game
		for _, g := range *gameList {

			if len(g) > 0 {
				// Starting with letter
				if strings.ToLower(g[0:1]) == v {
					temp := Game{urlSafe(g), images[g], videos[g], g, youtubesearch}
					tempGames = append(tempGames, temp)
				}
				// Starting with #
				if v == "num" {
					if _, err := strconv.Atoi(g[0:1]); err == nil {
						temp := Game{urlSafe(g), images[g], videos[g], g, youtubesearch}
						tempGames = append(tempGames, temp)
					}
				}
				// Text List
				if v == "textlist" {
					temp := Game{urlSafe(g), images[g], videos[g], g, youtubesearch}
					tempGames = append(tempGames, temp)
				}
			}
		}

		data := ListPageData{
			Sections: []Section{{"num", "#"}, {"a", "A"}, {"b", "B"}, {"c", "C"}, {"d", "D"}, {"e", "E"}, {"f", "F"}, {"g", "G"},
				{"h", "H"}, {"i", "I"}, {"j", "J"}, {"k", "K"}, {"l", "L"}, {"m", "M"}, {"n", "N"}, {"o", "O"}, {"p", "P"},
				{"q", "Q"}, {"r", "R"}, {"s", "S"}, {"t", "T"}, {"u", "U"}, {"v", "V"}, {"w", "W"}, {"x", "X"}, {"y", "Y"},
				{"z", "Z"}, {"textlist", "Text List"}},
			CurrentPage: v,
			Games:       tempGames,
			ListName:    listName,
			FolderName:  folderName,
			Credit:      template.HTML(getCredit(folderName)),
		}

		tmpl := template.Must(template.ParseFiles("list_layout.html", "navigation.html"))
		if err := tmpl.Execute(&tmplBuffer, data); err != nil {
			fmt.Println(err)
		}
		WriteToFile("public/mister/"+folderName+"/"+strings.ToLower(v)+".html", tmplBuffer.String())
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

func findAllFiles(root, ext string, skip string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			if d.Name() == skip {
				return filepath.SkipDir
			}
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func findAllSubstringFiles(root, ext string, skip string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			if d.Name() == skip {
				return filepath.SkipDir
			}
		}
		if strings.Contains(d.Name(), ext) {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func CopyFile(source string, destination string) error {
	bytesRead, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile(destination, bytesRead, 0644)
	if err != nil {
		return err
	}

	return nil
}

func FileSize(filename string) (int64, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

// Sample format:
// list: map[string]string{"Game Name": "https://www.example.com/page.html",}
// destination: "temp"
// extension: ".png"
// seconds: 8
func downloadFile(list map[string]string, destination string, extension string, seconds time.Duration) {
	for k, v := range list {
		if _, err := os.Stat(destination + "/" + k + ".html"); os.IsNotExist(err) {
			response, err := http.Get(v)
			if err != nil {
				log.Fatal(err)
			}

			outFile, err := os.Create(destination + "/" + k + extension)
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(outFile, response.Body)
			if err != nil {
				log.Fatal(err)
			}
			outFile.Close()
			response.Body.Close()
			time.Sleep(seconds * time.Second)
		}
	}
}

func getCredit(system string) string {
	if system == "nes" {
		return "Game images from Jardavius @ <a href='https://emumovies.com'>https://emumovies.com</a><br/>Please consider <a href='https://emumovies.com/subscriptions/'>donating</a> to EmuMovies."
	} else if system == "snes" {
		return "Game images from Jardavius @ <a href='https://emumovies.com'>https://emumovies.com</a><br/>Please consider <a href='https://emumovies.com/subscriptions/'>donating</a> to EmuMovies."
	} else if system == "sms" {
		return "Game images from <a href='https://www.smspower.org/'>https://www.smspower.org</a><br/>Please consider <a href='https://www.smspower.org/Home/Donate'>donating</a> to SMS Power."
	} else if system == "sg1000" {
		return "Game images from <a href='https://www.smspower.org/'>https://www.smspower.org</a><br/>Please consider <a href='https://www.smspower.org/Home/Donate'>donating</a> to SMS Power."
	} else if system == "genesis" {
		return "Game images from EmuMovies @ <a href='https://emumovies.com'>https://emumovies.com</a><br/>Please consider <a href='https://emumovies.com/subscriptions/'>donating</a> to EmuMovies."
	} else if system == "pce" {
		return "Game images from EmuMovies @ <a href='https://emumovies.com'>https://emumovies.com</a><br/>Please consider <a href='https://emumovies.com/subscriptions/'>donating</a> to EmuMovies."
	} else if system == "lynx" {
		return "Game images from <a href='https://atarigamer.com'>https://atarigamer.com</a><br/>Please consider <a href='https://atarigamer.com/pages/support-atari-gamer'>donating</a> to Atari Gamer."
	} else if system == "gbc" {
		return "Game images from Jardavius @ <a href='https://emumovies.com'>https://emumovies.com</a><br/>Please consider <a href='https://emumovies.com/subscriptions/'>donating</a> to EmuMovies."
	} else if system == "arcade" {
		return "Game images from Antonio Paradossi @ <a href='https://www.progettosnaps.net'>https://www.progettosnaps.net</a><br/>Please consider <a href='https://www.paypal.com/paypalme/progettoSNAPS'>donating</a> to progetto-SNAPS."

	} else {
		return ""
	}
}
