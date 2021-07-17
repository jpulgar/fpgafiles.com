package main

func generateMisterGBGames(generate bool) {
	gbTitleAdded := make(map[string]bool)
	gbImages := make(map[string]string)
	gbGameList := []string{}
	compileMisterConsoleData(gbTitleAdded, &gbGameList, gbImages, gbVideos, "gb")
	if generate {
		generateMisterConsoleHTML("Game Boy", &gbGameList, gbImages, gbVideos, "gb")
	}
}

var gbVideos = map[string]string{}
