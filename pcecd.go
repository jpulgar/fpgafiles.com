package main

func generateMisterPCECDGames(generate bool) {
	pcecdTitleAdded := make(map[string]bool)
	pcecdImages := make(map[string]string)
	pcecdGameList := []string{}
	compileMisterConsoleData(pcecdTitleAdded, &pcecdGameList, pcecdImages, pcecdVideos, "pcecd")
	//pceVideoReport(&pcecdGameList)
	if generate {
		generateMisterConsoleHTML("TurboGrafx 16 / PC Engine CD Games", &pcecdGameList, pcecdImages, pcecdVideos, "pcecd", "pc engine")
	}
}

var pcecdVideos = map[string]string{}
var pcecdLongplays = []string{}
