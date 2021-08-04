package main

func generateMisterColecoVisionGames(generate bool) {
	colecoTitleAdded := make(map[string]bool)
	colecoImages := make(map[string]string)
	colecoGameList := []string{}
	compileMisterConsoleData(colecoTitleAdded, &colecoGameList, colecoImages, colecoVideos, "coleco")
	//colecoVideoReport(&colecoGameList)
	if generate {
		generateMisterConsoleHTML("ColecoVision Games", &colecoGameList, colecoImages, colecoVideos, "coleco", "coleco")
	}
}

var colecoVideos = map[string]string{}

var colecoLongplays = []string{}
