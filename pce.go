package main

func generateMisterPCEGames() {
	pceTitleAdded := make(map[string]bool)
	pceImages := make(map[string]string)
	pceGameList := []string{}
	compileMisterConsoleData(pceTitleAdded, &pceGameList, pceImages, "pce")
	generateMisterConsoleHTML("TurboGrafx 16 / PC Engine", &pceGameList, pceImages, pceVideos, "pce")
}

var pceVideos = map[string]string{}
