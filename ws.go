package main

func generateMisterWonderSwanGames(generate bool) {
	wsTitleAdded := make(map[string]bool)
	wsImages := make(map[string]string)
	wsGameList := []string{}
	compileMisterConsoleData(wsTitleAdded, &wsGameList, wsImages, wsVideos, "ws")
	if generate {
		generateMisterConsoleHTML("WonderSwan", &wsGameList, wsImages, wsVideos, "ws", "wonder swan")
	}
}

var wsVideos = map[string]string{}
