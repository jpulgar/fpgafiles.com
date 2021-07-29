package main

func generateMisterGenesisGames(generate bool) {
	genesisTitleAdded := make(map[string]bool)
	genesisImages := make(map[string]string)
	genesisGameList := []string{}
	compileMisterConsoleData(genesisTitleAdded, &genesisGameList, genesisImages, genesisVideos, "genesis")
	if generate {
		generateMisterConsoleHTML("Genesis Games", &genesisGameList, genesisImages, genesisVideos, "genesis", "genesis")
	}
}

var genesisVideos = map[string]string{}
