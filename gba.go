package main

func generateMisterGBAGames(generate bool) {
	gbaTitleAdded := make(map[string]bool)
	gbaImages := make(map[string]string)
	gbaGameList := []string{}
	compileMisterConsoleData(gbaTitleAdded, &gbaGameList, gbaImages, gbaVideos, "gba")
	if generate {
		generateMisterConsoleHTML("Game Boy Advance Games", &gbaGameList, gbaImages, gbaVideos, "gba")
	}
}

var gbaVideos = map[string]string{}
