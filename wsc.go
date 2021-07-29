package main

func generateMisterWonderSwanColorGames(generate bool) {
	wscTitleAdded := make(map[string]bool)
	wscImages := make(map[string]string)
	wscGameList := []string{}
	compileMisterConsoleData(wscTitleAdded, &wscGameList, wscImages, wscVideos, "wsc")
	if generate {
		generateMisterConsoleHTML("WonderSwan Color", &wscGameList, wscImages, wscVideos, "wsc", "wonder swan color")
	}
}

var wscVideos = map[string]string{}
