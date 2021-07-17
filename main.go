package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type HomePageData struct {
	UpdatedOn string
}

func main() {

	generateIndex() // always run
	generateAll := false
	// generateMisterArcadeGames(generateAll)
	// generateMisterNeoGeoGames(generateAll)
	generateMisterNESGames(generateAll)
	generateMisterSNESGames(generateAll)
	generateMisterSMSGames(generateAll)
	generateMisterGBCGames(generateAll)
	generateMisterGBAGames(generateAll)
	generateMisterGenesisGames(generateAll)
	generateMisterSegaCDGames(generateAll)
	generateMisterAtari2600Games(generateAll)
	generateMisterPCEGames(generateAll)
	generateMisterLynxGames(generateAll)
}

func generateIndex() {
	var tmplBuffer bytes.Buffer

	now := time.Now()
	dataHomepage := HomePageData{
		UpdatedOn: now.Format("January 06, 2006"),
	}
	tmpl := template.Must(template.ParseFiles("index_layout.html", "navigation.html"))
	if err := tmpl.Execute(&tmplBuffer, dataHomepage); err != nil {
		fmt.Println(err)
	}
	WriteToFile("public/index.html", tmplBuffer.String())
	tmplBuffer.Reset()

	generateMisterArcadeCommands()
}
