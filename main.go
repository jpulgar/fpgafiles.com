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
	generateIndex()
	generateMisterArcadeGames()
	generateMisterNESGames()
	generateMisterSNESGames()
	generateMisterSMSGames()
	generateMisterGBCGames()
	generateMisterGenesisGames()
	generateMisterPCEGames()
	generateMisterLynxGames()
}

func generateIndex() {
	var tmplBuffer bytes.Buffer

	now := time.Now()
	dataHomepage := HomePageData{
		UpdatedOn: now.Format("Monday, January 06, 2006 @ 3:04PM MST"),
	}
	tmpl := template.Must(template.ParseFiles("index_layout.html", "navigation.html"))
	if err := tmpl.Execute(&tmplBuffer, dataHomepage); err != nil {
		fmt.Println(err)
	}
	WriteToFile("public/index.html", tmplBuffer.String())
	tmplBuffer.Reset()
}
