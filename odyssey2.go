package main

func generateMisterOdyssey2Games(generate bool) {
	odyssey2TitleAdded := make(map[string]bool)
	odyssey2Images := make(map[string]string)
	odyssey2GameList := []string{}
	compileMisterConsoleData(odyssey2TitleAdded, &odyssey2GameList, odyssey2Images, odyssey2Videos, "odyssey2")
	//odyssey2VideoReport(&odyssey2GameList)
	if generate {
		generateMisterConsoleHTML("Odyssey2 Games", &odyssey2GameList, odyssey2Images, odyssey2Videos, "odyssey2", "odyssey2")
	}
}

var odyssey2Videos = map[string]string{}

var odyssey2Longplays = []string{}
