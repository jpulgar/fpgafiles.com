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
	CopyHomepageFiles()
	generateHomepageHTML()
	generateMisterArcadeGames()
}

func CopyHomepageFiles() {
	err := CopyFile("assets/fpgafiles.png", "public/fpgafiles.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = CopyFile("assets/favicon.ico", "public/favicon.ico")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func generateHomepageHTML() {

	var tmplBuffer bytes.Buffer

	now := time.Now()
	dataHomepage := HomePageData{
		UpdatedOn: now.Format(time.RFC1123),
	}
	tmpl := template.Must(template.ParseFiles("index_layout.html", "navigation.html"))
	if err := tmpl.Execute(&tmplBuffer, dataHomepage); err != nil {
		fmt.Println(err)
	}
	WriteToFile("public/index.html", tmplBuffer.String())
	tmplBuffer.Reset()

}
