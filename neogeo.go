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

func generateMisterNeoGeoGames(generate bool) {
	compileMisterNeoGeoData()
	neogeoVideoReport()
	if generate {
		generateMisterNeoGeoHTML()
	}
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

	// Write stats.json for Homepage Use
	videosFound := 0
	for _, v := range neogeoVideos {
		if v != "" {
			videosFound++
		}
	}
	videoPercentage := 0.00
	if videosFound != 0 {
		videoPercentage = float64(videosFound) / float64(len(neogeoGameList))
	}

	stats := Stats{nameForFolder("neogeo"), len(neogeoGameList), float64(int(videoPercentage*100*100)) / 100, "neogeo"}
	prettyJSON, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	WriteToFile("public/mister/neogeo/stats.json", string(prettyJSON))

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
				temp := Game{g.SetName, g.SetName + ".png", neogeoVideos[g.SetName], g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
			// Starting with #
			if v == "num" {
				if _, err := strconv.Atoi(g.Name[0:1]); err == nil {
					temp := Game{g.SetName, g.SetName + ".png", neogeoVideos[g.SetName], g.Name + " (" + g.Year + ")"}
					tempGames = append(tempGames, temp)
				}
			}
			// Text List
			if v == "textlist" {
				temp := Game{g.SetName, g.SetName + ".png", neogeoVideos[g.SetName], g.Name + " (" + g.Year + ")"}
				tempGames = append(tempGames, temp)
			}
		}

		data := ListPageData{
			Sections: []Section{{"num", "#"}, {"a", "A"}, {"b", "B"}, {"c", "C"}, {"d", "D"}, {"e", "E"}, {"f", "F"}, {"g", "G"},
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
			Video:  neogeoVideos[v.SetName],
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

var neogeoVideos = map[string]string{
	"2020bb":   "8d2fxVEMHVM",
	"3countb":  "4K39cEB8TuE",
	"alpham2":  "sZN0b4STzwU",
	"alpham2p": "sZN0b4STzwU",
	"androdun": "iQOrXlf34es",
	"aodk":     "KQXR1f1aonA",
	"aof":      "-iQwJMIS7OU",
	"aof2":     "QBPmfdl2iDw",
	"aof3":     "SSvJnIo36YU",
	"b2b":      "3DncxpUdQEQ",
	"bakatono": "CKSc7g_JlsU",
	"bangbead": "OXfnZekue90",
	"bjourney": "68slIqrMEEo",
	"blazstar": "NLpLqgqimtI",
	"breakers": "WqtsNLfo0jA",
	"breakrev": "Z8JCOSNqX-c",
	"bstars":   "pLihYJI9Ims",
	"bstars2":  "Yo1AkhYPsME",
	"burningf": "LR2NzP9Zrts",
	"crswd2bl": "UFp8bEl5tvQ",
	"crsword":  "H485or8VOC8",
	"ct2k3sa":  "wkj08ignM4w",
	"ctomaday": "a4KXiUd3Yq4",
	"cyberlip": "11mrYej9b_g",
	"diggerma": "ZB9GJHHoGis",
	"doubledr": "jo6R6LqpNNY",
	"dragonsh": "9z994UeJ44I",
	"eightman": "hObhDF6DEvQ",
	"fatfursp": "AlZ1q2Gddsc",
	"fatfury1": "oqcYWJfLdSs",
	"fatfury2": "HbWPpkTN4tY",
	"fatfury3": "aZa05jLtFi0",
	"fbfrenzy": "lRHVlzNq2e4",
	"fightfev": "GniTJ2VdP8Y",
	"flipshot": "sFZHNw2DxYY",
	"froman2b": "uaw8xBGDod8",
	"fswords":  "Xt813iqWwdk",
	"galaxyfg": "T2_Qxys8Ezo",
	"ganryu":   "nVAplIZtbSE",
	"garou":    "yOkyX-eiofw",
	"garoubl":  "yOkyX-eiofw",
	"garouh":   "yOkyX-eiofw",
	"garoup":   "yOkyX-eiofw",
	"ghostlop": "oaUCJWiTBZ4",
	"goalx3":   "bSPQbx5He-U",
	"gowcaizr": "t8SvI-2nq-Q",
	"gpilots":  "W-eirUeg2AM",
	"gururin":  "l38qT7SjzwY",
	"ironclad": "cwru2DS8egY",
	"irrmaze":  "9E-201c6Vv8",
	"janshin":  "pPmEln8tYG0",
	"joyjoy":   "-ar3NP0hNOI",
	"kabukikl": "8W-vjrs9gbs",
	"karnovr":  "ujnRoLwiIlg",
	"kf10thep": "e3Xk9S8sdmc",
	"kf2k2mp":  "wlZKfAxjTw0",
	"kf2k2pls": "PDK7zl8yewI",
	"kf2k5uni": "UMYS8rJm-JM",
	"kizuna":   "V2aLaBlI69I",
	"kof2000":  "xrQLRCwM1Ds",
	"kof2001":  "-2TcH2ismsY",
	"kof2002":  "YNfD5ndPSjs",
	"kof2003":  "IFHvDdUAnzI",
	"kof2k4se": "WLkRhiWFd6w",
	"kof94":    "g4G0dX-w8QI",
	"kof95":    "greALx7jn4U",
	"kof96":    "yuR-XEzUHAA",
	"kof97":    "H5lc6Lle2lc",
	"kof98":    "knFWtmDkDiM",
	"kof99":    "DuAg9ILvzUo",
	"kof99p":   "7zm3rscp5uI",
	"kog":      "5FuSF7UbeRw",
	"kotm":     "fpkyLJsiHSw",
	"kotm2":    "hzcGFbUBWhY",
	"lans2004": "3pr4T0t7YnA",
	"lastblad": "pYoVkj-WCrg",
	"lastbld2": "-DQsa_nM9e4",
	"lasthope": "4FKa3HZd6Tk",
	"lastsold": "Ip_GTG8-Xp0",
	"lbowling": "UWoGPJGjnAo",
	"legendos": "PHhuFChvqd0",
	"lresort":  "y1xUT7QJ-mA",
	"magdrop2": "W0f1uiunMA8",
	"magdrop3": "X4aOHQxbJ9o",
	"maglord":  "H7ToKGzgXm0",
	"mahretsu": "Sp98HtTNatY",
	"marukodq": "oPuVFSIG6P4",
	"matrim":   "mt4NEXOTXpg",
	"miexchng": "XhnZAy_CNlA",
	"minasan":  "TtgT_anGCEE",
	"moshougi": "GYIvHFwOk84",
	"ms4plus":  "ngfD2YrTbIo",
	"mslug":    "9fzLnRvsRc4",
	"mslug2":   "6gAVh4atOSQ",
	"mslug2t":  "0GxV8wdO9pM",
	"mslug3":   "OPCHeaqZZBQ",
	"mslug4":   "ngfD2YrTbIo",
	"mslug5":   "IjM7kJ6Kj1Q",
	"mslug6":   "5-BfrFiIPt0",
	"mslugx":   "F5Q3CCmVq4E",
	"mutnat":   "16FdFl3KK0I",
	"nam1975":  "TfvBhYF-6jI",
	"ncombat":  "ShfdduWAyJg",
	"ncommand": "eoyQdxJgD0I",
	"neobombe": "It8l3CNXUyg",
	"neocup98": "miTLcvursxk",
	"neodrift": "cu9cWFqi3lk",
	"neomrdo":  "l4pmbdAtjYg",
	"ninjamas": "4auIUYQD5_E",
	"nitd":     "esTq3L5LR0U",
	"overtop":  "8fE_PMryMuM",
	"panicbom": "3XHS5n7HuRM",
	"pbobbl2n": "_MMq7yrV6fI",
	"pbobblen": "F1I0lzM_UZI",
	"pgoal":    "qcvsKWSfKgE",
	"pnyaa":    "f22CCxJ9eko",
	"popbounc": "U4FXsVClQpA",
	"preisle2": "pCUuVrxNH2A",
	"pspikes2": "dYpYqoHqWUM",
	"pulstar":  "GHkRiKPhito",
	"puzzldpr": "PRQmuMbvJb0",
	"puzzledp": "elldJ5ZZsvU",
	"quizdai2": "4zwNAxdA-X4",
	"quizdais": "Ldk7C3UU8lw",
	"quizkof":  "w0PCL2WjBWI",
	"ragnagrd": "9rEcf390X-Y",
	"rbff1":    "uxvLQbNznLQ",
	"rbff2":    "jhWomIO4EHk",
	"rbffspec": "Q4O6a9RYQ7A",
	"ridhero":  "iwsIuiK7Rjo",
	"roboarmy": "6f4U5CMLjro",
	"rotd":     "msZUQkEZKZQ",
	"s1945p":   "Zv-sd9C4ne8",
	"samsh5sp": "sGCCuZxhj_0",
	"samsho":   "zWmOMwJB8xQ",
	"samsho2":  "pJdfRC0fnEc",
	"samsho3":  "in6ojxEZH1I",
	"samsho4":  "KxtpJEvTBv4",
	"samsho5":  "P8Oeqpqlx9w",
	"savagere": "QPng_eh_864",
	"sbp":      "LfwXB-eGxEs",
	"sdodgeb":  "NKyAUOng7ew",
	"sengoku":  "SLLKQLYHQWw",
	"sengoku2": "NReHy05g26g",
	"sengoku3": "17F_NMw_tXo",
	"shocktr2": "4DF5ZFsX1MY",
	"shocktro": "RuFGB_h1_3Q",
	"socbrawl": "2_k2dnstm6w",
	"sonicwi2": "PXc0Rksr61I",
	"sonicwi3": "mKpC_lnB7FM",
	"spinmast": "VVK7dZGYgfA",
	"ssideki":  "OrP_V6lrSDI",
	"ssideki2": "4-piSx2a1zQ",
	"ssideki3": "sBUa5xJrlyo",
	"ssideki4": "eqW5SkPLh0M",
	"stakwin":  "fa9LWD7MxuE",
	"stakwin2": "Acd2QX8DSzw",
	"strhoop":  "7dRpE7FDWms",
	"superspy": "5bFFkNmhN_w",
	"svc":      "ctBMxNDGciY",
	"svcplus":  "5UbP_hZkTGM",
	"svcsplus": "kgYFyGYasi4",
	"tophuntr": "XEDONk2oaD0",
	"tpgolf":   "KYjVuVX-Vck",
	"trally":   "KjMP7PxKixA",
	"turfmast": "WDt9HYniyHs",
	"twinspri": "zEhFh_-DVD8",
	"twsoc96":  "Nf2KPfuNxCs",
	"viewpoin": "IPBQ2NhSO9M",
	"wakuwak7": "DefAwU3GcEI",
	"wh1":      "QDYhvy7uCCo",
	"wh2":      "iVBHgr6TH2w",
	"wh2j":     "1jOGp4Gtt3w",
	"whp":      "snP1uLk0s_8",
	"wjammers": "Eg93nPgR6rw",
	"zedblade": "E_YAKJSFvR4",
	"zintrckb": "lawxVuJi9VY",
	"zupapa":   "a2diSxAds2Q",
}

func neogeoVideoReport() {
	// Calculate Best Video Matches
	var distZero = 0
	var distOne = 0
	var distTwo = 0
	var distThree = 0
	var distFour = 0
	var distFive = 0
	var distMoreThanFive = 0
	for _, v := range neogeoGameList {
		// Only do this process for titles with no video
		if neogeoVideos[v.SetName] == "" {

			tempName := v.Name

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
			for _, n := range arcadeLongplays {
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
			//   0, 1, 2, 3 are done
			if lowestDistance > 0 {

				//lowestName = lowestName[strings.IndexByte(lowestName, '|')+1:] // for final
				//fmt.Println("\"" + v.Name + "\": \"" + lowestName + "\",") // for testing

				// For the rest without matches
				fmt.Println("\"" + v.SetName + "\": \"\",")
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
