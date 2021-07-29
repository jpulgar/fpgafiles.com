package main

func generateMisterAtari2600Games(generate bool) {
	atari2600TitleAdded := make(map[string]bool)
	atari2600Images := make(map[string]string)
	atari2600GameList := []string{}
	compileMisterConsoleData(atari2600TitleAdded, &atari2600GameList, atari2600Images, atari2600Videos, "atari2600")
	if generate {
		generateMisterConsoleHTML("Atari 2600 Games", &atari2600GameList, atari2600Images, atari2600Videos, "atari2600", "atari 2600")
	}
}

var atari2600Videos = map[string]string{}
