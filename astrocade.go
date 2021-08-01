package main

func generateMisterAstrocadeGames(generate bool) {
	astrocadeTitleAdded := make(map[string]bool)
	astrocadeImages := make(map[string]string)
	astrocadeGameList := []string{}
	compileMisterConsoleData(astrocadeTitleAdded, &astrocadeGameList, astrocadeImages, astrocadeVideos, "astrocade")
	//astrocadeVideoReport(&astrocadeGameList)
	if generate {
		generateMisterConsoleHTML("Astrocade Games", &astrocadeGameList, astrocadeImages, astrocadeVideos, "astrocade", "astrocade")
	}
}

var astrocadeVideos = map[string]string{}

var astrocadeLongplays = []string{}
