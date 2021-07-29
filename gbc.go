package main

func generateMisterGBCGames(generate bool) {
	gbcTitleAdded := make(map[string]bool)
	gbcImages := make(map[string]string)
	gbcGameList := []string{}
	compileMisterConsoleData(gbcTitleAdded, &gbcGameList, gbcImages, gbcVideos, "gbc")
	if generate {
		generateMisterConsoleHTML("Game Boy Color Games", &gbcGameList, gbcImages, gbcVideos, "gbc", "game boy color")
	}
}

var gbcVideos = map[string]string{}
