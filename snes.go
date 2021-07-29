package main

func generateMisterSNESGames(generate bool) {
	snesTitleAdded := make(map[string]bool)
	snesImages := make(map[string]string)
	snesGameList := []string{}
	compileMisterConsoleData(snesTitleAdded, &snesGameList, snesImages, snesVideos, "snes")
	if generate {
		generateMisterConsoleHTML("Super Nintendo Games", &snesGameList, snesImages, snesVideos, "snes", "super nintendo")
	}
}

var snesVideos = map[string]string{}
