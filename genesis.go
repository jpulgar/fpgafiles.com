package main

func generateMisterGenesisGames() {
	genesisTitleAdded := make(map[string]bool)
	genesisImages := make(map[string]string)
	genesisGameList := []string{}
	compileMisterConsoleData(genesisTitleAdded, &genesisGameList, genesisImages, "genesis")
	generateMisterConsoleHTML("Genesis", &genesisGameList, genesisImages, genesisVideos, "genesis")
}

var genesisVideos = map[string]string{}
