package main

func generateMisterGBGames() {
	gbTitleAdded := make(map[string]bool)
	gbImages := make(map[string]string)
	gbGameList := []string{}
	compileMisterConsoleData(gbTitleAdded, &gbGameList, gbImages, "gb")
	generateMisterConsoleHTML("Game Boy", &gbGameList, gbImages, gbVideos, "gb")
}

var gbVideos = map[string]string{}
