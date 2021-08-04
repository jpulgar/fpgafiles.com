package main

func generateMisterVectrexGames(generate bool) {
	vectrexTitleAdded := make(map[string]bool)
	vectrexImages := make(map[string]string)
	vectrexGameList := []string{}
	compileMisterConsoleData(vectrexTitleAdded, &vectrexGameList, vectrexImages, vectrexVideos, "vectrex")
	//vectrexVideoReport(&vectrexGameList)
	if generate {
		generateMisterConsoleHTML("Vectrex Games", &vectrexGameList, vectrexImages, vectrexVideos, "vectrex", "vectrex")
	}
}

var vectrexVideos = map[string]string{}

var vectrexLongplays = []string{}
