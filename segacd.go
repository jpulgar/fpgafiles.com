package main

func generateMisterSegaCDGames() {
	segacdTitleAdded := make(map[string]bool)
	segacdImages := make(map[string]string)
	segacdGameList := []string{}
	compileMisterConsoleData(segacdTitleAdded, &segacdGameList, segacdImages, "segacd")
	generateMisterConsoleHTML("Sega CD Games", &segacdGameList, segacdImages, segacdVideos, "segacd")
}

var segacdVideos = map[string]string{}
