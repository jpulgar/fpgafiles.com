package main

func generateMisterSuperGrafxGames(generate bool) {
	supergrafxTitleAdded := make(map[string]bool)
	supergrafxImages := make(map[string]string)
	supergrafxGameList := []string{}
	compileMisterConsoleData(supergrafxTitleAdded, &supergrafxGameList, supergrafxImages, supergrafxVideos, "supergrafx")
	//pceVideoReport(&supergrafxGameList)
	if generate {
		generateMisterConsoleHTML("SuperGrafx Games", &supergrafxGameList, supergrafxImages, supergrafxVideos, "supergrafx", "supergrafx")
	}
}

var supergrafxVideos = map[string]string{
	"1941 - Counter Attack": "DwzDaDNUxio",
	"Aldynes":               "tJyylrgm07c",
	"Battle Ace":            "WTaAzy5BLGU",
	"Daimakaimura":          "DwBKqABy7jg",
	"Mado King Granzort":    "RXN2JOXep5M",
}
var supergrafxLongplays = []string{}
