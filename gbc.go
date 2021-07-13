package main

func generateMisterGBCGames() {
	gbcTitleAdded := make(map[string]bool)
	gbcImages := make(map[string]string)
	gbcGameList := []string{}
	compileMisterConsoleData(gbcTitleAdded, &gbcGameList, gbcImages, "gbc")
	generateMisterConsoleHTML("Game Boy Color", &gbcGameList, gbcImages, gbVideos, "gbc")
}

var gbcVideos = map[string]string{}
