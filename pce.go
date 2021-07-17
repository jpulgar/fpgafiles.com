package main

func generateMisterPCEGames(generate bool) {
	pceTitleAdded := make(map[string]bool)
	pceImages := make(map[string]string)
	pceGameList := []string{}
	compileMisterConsoleData(pceTitleAdded, &pceGameList, pceImages, pceVideos, "pce")
	if generate {
		generateMisterConsoleHTML("TurboGrafx 16 / PC Engine Games", &pceGameList, pceImages, pceVideos, "pce")
	}
}

var pceVideos = map[string]string{}
