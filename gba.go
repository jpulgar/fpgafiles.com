package main

func generateMisterGBAGames() {
	gbaTitleAdded := make(map[string]bool)
	gbaImages := make(map[string]string)
	gbaGameList := []string{}
	compileMisterConsoleData(gbaTitleAdded, &gbaGameList, gbaImages, "gba")
	generateMisterConsoleHTML("Game Boy Advance Games", &gbaGameList, gbaImages, gbaVideos, "gba")
}

var gbaVideos = map[string]string{}
