package main

func generateMisterSNESGames() {
	snesTitleAdded := make(map[string]bool)
	snesImages := make(map[string]string)
	snesGameList := []string{}
	compileMisterConsoleData(snesTitleAdded, &snesGameList, snesImages, "snes")
	generateMisterConsoleHTML("Super Nintendo Games", &snesGameList, snesImages, snesVideos, "snes")
}

var snesVideos = map[string]string{}
