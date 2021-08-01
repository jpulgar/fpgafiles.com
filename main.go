package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type HomePageData struct {
	UpdatedOn string
	Stats     []Stats
}

func main() {

	generateAll := true
	// generateMisterArcadeGames(generateAll)
	// generateMisterNeoGeoGames(generateAll)
	// generateMisterNESGames(generateAll)
	// generateMisterSNESGames(generateAll)
	// generateMisterSMSGames(generateAll)
	//generateMisterSG1000Games(generateAll)
	generateMisterGameGearGames(generateAll)
	// generateMisterGBGames(generateAll)
	// generateMisterGBCGames(generateAll)
	// generateMisterGBAGames(generateAll)
	// generateMisterGenesisGames(generateAll)
	// generateMisterSegaCDGames(generateAll)
	// generateMisterAtari2600Games(generateAll)
	// generateMisterAstrocadeGames(generateAll)
	// generateMisterPCEGames(generateAll)
	// generateMisterLynxGames(generateAll)
	// generateMisterWonderSwanGames(generateAll)
	// generateMisterWonderSwanColorGames(generateAll)

	generateIndex() // run last

}

func generateIndex() {
	var tmplBuffer bytes.Buffer
	var jsonStatFiles = []string{"arcade", "neogeo", "atari2600", "astrocade", "lynx", "pce", "nes", "gb", "gbc", "gba",
		"snes", "sms", "genesis", "segacd", "ws", "wsc"}
	var tempStats = []Stats{}
	for _, v := range jsonStatFiles {
		out := Stats{}
		if err := readStatFile("public/mister/"+v+"/stats.json", &out); err != nil {
			fmt.Println(err)
			return
		}
		temp := Stats{out.Name, out.Games, out.LongplayPercentage, v}
		tempStats = append(tempStats, temp)
	}

	now := time.Now()
	//fmt.Println(now.Format(time.RFC1123))
	dataHomepage := HomePageData{
		UpdatedOn: now.Format(time.RFC1123),
		Stats:     tempStats,
	}
	tmpl := template.Must(template.ParseFiles("index_layout.html", "navigation.html"))
	if err := tmpl.Execute(&tmplBuffer, dataHomepage); err != nil {
		fmt.Println(err)
	}
	WriteToFile("public/index.html", tmplBuffer.String())
	tmplBuffer.Reset()

	generateMisterArcadeCommands()
}
