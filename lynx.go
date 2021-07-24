package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateMisterLynxGames(generate bool) {
	lynxTitleAdded := make(map[string]bool)
	lynxImages := make(map[string]string)
	lynxGameList := []string{}
	compileMisterConsoleData(lynxTitleAdded, &lynxGameList, lynxImages, lynxVideos, "lynx")
	//lynxVideoReport(&lynxGameList)
	if generate {
		generateMisterConsoleHTML("Lynx Games", &lynxGameList, lynxImages, lynxVideos, "lynx")
	}
}

func lynxVideoReport(gameList *[]string) {
	// Calculate Best Video Matches
	var distZero = 0
	var distOne = 0
	var distTwo = 0
	var distThree = 0
	var distFour = 0
	var distFive = 0
	var distMoreThanFive = 0
	for _, v := range *gameList {
		// Only do this process for titles with no video
		if lynxVideos[v] == "" {

			tempName := v

			if idx := strings.IndexByte(tempName, '('); idx >= 0 {
				tempName = strings.TrimRight(tempName[:idx], " ")
			}
			if strings.Contains(tempName, ", The") {
				tempName = strings.Replace(tempName, ", The", "", 1)
				tempName = "The " + tempName
			}

			var str1 = []rune(tempName)
			var lowestDistance = 99
			var lowestName = ""
			for _, n := range lynxLongplays {
				temptemp := n[:strings.IndexByte(n, '|')]
				var str2 = []rune(temptemp)
				var tempDistance = levenshtein(str1, str2)
				if tempDistance < lowestDistance {
					lowestDistance = tempDistance
					lowestName = n //temptemp
				}
			}
			if lowestDistance == 0 {
				distZero = distZero + 1
			}
			if lowestDistance == 1 {
				distOne = distOne + 1
			}
			if lowestDistance == 2 {
				distTwo = distTwo + 1
			}
			if lowestDistance == 3 {
				distThree = distThree + 1
			}
			if lowestDistance == 4 {
				distFour = distFour + 1
			}
			if lowestDistance == 5 {
				distFive = distFive + 1
			}
			if lowestDistance > 5 {
				distMoreThanFive = distMoreThanFive + 1
			}

			// Cycle through: <= 2, == 3, == 4, == 5, > 5
			//  0, 1, 2 are correct
			if lowestDistance >= 2 {

				// lowestName = lowestName[strings.IndexByte(lowestName, '|')+1:] // for final
				// fmt.Println("\"" + v + "\": \"" + lowestName + "\",")          // for testing

				// For the rest without matches
				fmt.Println("\"" + v + "\": \"\",")
				if false {
					fmt.Println(lowestName)
				}
			}
		}
	}
	// Reporting
	fmt.Println("Distance 0: " + strconv.Itoa(distZero))
	fmt.Println("Distance 1: " + strconv.Itoa(distOne))
	fmt.Println("Distance 2: " + strconv.Itoa(distTwo))
	fmt.Println("Distance 3: " + strconv.Itoa(distThree))
	fmt.Println("Distance 4: " + strconv.Itoa(distFour))
	fmt.Println("Distance 5: " + strconv.Itoa(distFive))
	fmt.Println("Distance 5+: " + strconv.Itoa(distMoreThanFive))
}

var lynxVideos = map[string]string{
	"A.P.B.":                            "APoxk2hmStY",
	"Awesome Golf":                      "f7fGulSlXBc",
	"Baseball Heroes":                   "LUF_uELADuo",
	"Basketbrawl":                       "vuZzA_Rifvg",
	"Batman Returns":                    "oj-0b1gFE3U",
	"BattleWheels":                      "rvHSEghdS3w",
	"Battlezone 2000":                   "XyisEFBRqZo",
	"Bill & Ted's Excellent Adventure":  "BYmfII3wGXw",
	"Block Out":                         "Erq8iOQlMt4",
	"Blue Lightning":                    "BiGqnUlsKP4",
	"California Games":                  "GSFYx5_bdNU",
	"Checkered Flag":                    "bOHgSxOmKeM",
	"Chip's Challenge":                  "Qw4EFHGemfQ",
	"Crystal Mines II":                  "4og9Orsz3DI",
	"Desert Strike":                     "MZIy_mjT9dI",
	"Dinolympics":                       "6xfylVuy1KY",
	"Dirty Larry - Renegade Cop":        "n-p6j5OFWHI",
	"Double Dragon":                     "fTQo6SF3M0k",
	"Dracula - The Undead":              "fH-wcU---fM",
	"Electrocop":                        "FKc_QznAIJY",
	"European Soccer Challenge":         "uv4Ffq7FJi4",
	"Fidelity Ultimate Chess Challenge": "kcXA_PWCRJM",
	"Gates of Zendocon":                 "1CA-IvToZcM",
	"Gauntlet - The Third Encounter":    "ZplHuUekqF8",
	"Gordo 106":                         "xDvjHSKXilI",
	"Hard Drivin'":                      "iJL8iA0OrQA",
	"Hockey":                            "mrvsQqAUJtw",
	"Hydra":                             "hEyLQoDAcSM",
	"Ishido - The Way of the Stones":    "o93cuoOpIf4",
	"Jimmy Connors' Tennis":             "sEVcTBVUArk",
	"Joust":                             "FgCNbrSwVJA",
	"KLAX":                              "poo4pDDtSNo",
	"Kung Food":                         "iG6XlA3ijxg",
	"Lemmings":                          "2HySrPl_lcY",
	"Lynx Casino":                       "VzaSoU9bKs0",
	"Malibu Bikini Volleyball":          "eScqNV5VaWM",
	"Ms. Pac-Man":                       "ar0DESrKquo",
	"NFL Football":                      "OIDoZDik-LE",
	"Ninja Gaiden III - The Ancient Ship of Doom": "hJxsqLw86bk",
	"Ninja Gaiden":                      "khK5RD5CFOc",
	"Pac-Land":                          "NBLGgP6phT4",
	"Paperboy":                          "HohglNPEGSI",
	"Pinball Jam":                       "etkDiZZD8Yw",
	"Pit-Fighter":                       "aI772oxFQ3s",
	"Power Factor":                      "mJgDwh85_Ws",
	"Qix":                               "T7tQo6VX7W8",
	"Rampage":                           "T_E3W_JGpzI",
	"Rampart":                           "29WA1feWIpI",
	"RoadBlasters":                      "1JBliC-dcpo",
	"Robo-Squash":                       "f4hv6FqYl1I",
	"Robotron 2084":                     "vomOrfaQJfU",
	"Rygar":                             "3nofdhw79PI",
	"S.T.U.N. Runner":                   "1LK6R3nDM7w",
	"Scrapyard Dog":                     "5Cz9J_ineY0",
	"Shadow of the Beast":               "IxGxczM78oU",
	"Shanghai":                          "phnNtxUFKB8",
	"Steel Talons":                      "6N4bpVOL7r0",
	"Super Asteroids & Missile Command": "ITQtOF7Wr_s",
	"Super Off Road":                    "hHTMpuRJu0Q",
	"Super Skweek":                      "VMh1Ft281Bs",
	"Switchblade II":                    "fG8eWllV8fY",
	"Todd's Adventures in Slime World":  "ZR47s1r7x8c",
	"Toki":                              "2BrKMKIve9U",
	"Tournament Cyberball 2072":         "PJagJithBdY",
	"Turbo Sub":                         "6bn_l3s5zDI",
	"Viking Child":                      "RXVjOXe8QYA",
	"Warbirds":                          "2z-PKq21SIE",
	"World Class Fussball - Soccer":     "y4a6SrejKqM",
	"Xenophobe":                         "Wa6EpquD5T0",
	"Xybots":                            "nZKFk72KvHU",
	"Zarlor Mercenary":                  "oBPz6sr5NNQ",
}

var lynxGameInfo = map[string]string{
	"A.P.B.":                            "https://atarigamer.com/lynx/game/APB/1513416078",
	"Awesome Golf":                      "https://atarigamer.com/lynx/game/AwesomeGolf/420322826",
	"Baseball Heroes":                   "https://atarigamer.com/lynx/game/BaseballHeroes/1999363434",
	"Basketbrawl":                       "https://atarigamer.com/lynx/game/Basketbrawl/278645002",
	"Batman Returns":                    "https://atarigamer.com/lynx/game/BatmanReturns/411617623",
	"BattleWheels":                      "https://atarigamer.com/lynx/game/BattleWheels/249599870",
	"Battlezone 2000":                   "https://atarigamer.com/lynx/game/Battlezone2000/1622914826",
	"Bill & Ted's Excellent Adventure":  "https://atarigamer.com/lynx/game/BillampTedsExcellentAdventure/1397989147",
	"Block Out":                         "https://atarigamer.com/lynx/game/BlockOut/1686572735",
	"Blue Lightning":                    "https://atarigamer.com/lynx/game/BlueLightning/1382971832",
	"California Games":                  "https://atarigamer.com/lynx/game/CaliforniaGames/625857223",
	"Checkered Flag":                    "https://atarigamer.com/lynx/game/CheckeredFlag/1034307975",
	"Chip's Challenge":                  "https://atarigamer.com/lynx/game/ChipsChallenge/2022819869",
	"Crystal Mines II":                  "https://atarigamer.com/lynx/game/CrystalMinesII/635347429",
	"Desert Strike":                     "https://atarigamer.com/lynx/game/DesertStrike/619991931",
	"Dinolympics":                       "https://atarigamer.com/lynx/game/Dinolympics/979699310",
	"Dirty Larry - Renegade Cop":        "https://atarigamer.com/lynx/game/DirtyLarryRenegadeCop/1519759294",
	"Double Dragon":                     "https://atarigamer.com/lynx/game/DoubleDragon/2041198702",
	"Dracula - The Undead":              "https://atarigamer.com/lynx/game/DraculaTheUndead/358004199",
	"Electrocop":                        "https://atarigamer.com/lynx/game/Electrocop/376582658",
	"European Soccer Challenge":         "https://atarigamer.com/lynx/game/EuropeanSoccerChallenge/1967037688",
	"Fidelity Ultimate Chess Challenge": "https://atarigamer.com/lynx/game/FidelityUltimateChessChallenge/287160750",
	"Gates of Zendocon":                 "https://atarigamer.com/lynx/game/GatesofZendocon/1628880814",
	"Gauntlet - The Third Encounter":    "https://atarigamer.com/lynx/game/GauntletTheThirdEncounter/1424091724",
	"Gordo 106":                         "https://atarigamer.com/lynx/game/Gordo106/1208819515",
	"Hard Drivin'":                      "https://atarigamer.com/lynx/game/HardDrivin/1546263734",
	"Hockey":                            "https://atarigamer.com/lynx/game/Hockey/2124334099",
	"Hydra":                             "https://atarigamer.com/lynx/game/Hydra/355953500",
	"Ishido - The Way of the Stones":    "https://atarigamer.com/lynx/game/IshidoTheWayoftheStones/428251110",
	"Jimmy Connors' Tennis":             "https://atarigamer.com/lynx/game/JimmyConnorsTennis/781957960",
	"Joust":                             "https://atarigamer.com/lynx/game/Joust/2070457013",
	"KLAX":                              "https://atarigamer.com/lynx/game/KLAX/1284965295",
	"Kung Food":                         "https://atarigamer.com/lynx/game/KungFood/519135118",
	"Lemmings":                          "https://atarigamer.com/lynx/game/Lemmings/654071298",
	"Lynx Casino":                       "https://atarigamer.com/lynx/game/LynxCasino/592741719",
	"Malibu Bikini Volleyball":          "https://atarigamer.com/lynx/game/MalibuBikiniVolleyball/604444289",
	"Ms. Pac-Man":                       "https://atarigamer.com/lynx/game/MsPacMan/1016758045",
	"NFL Football":                      "https://atarigamer.com/lynx/game/NFLFootball/2168813906",
	"Ninja Gaiden":                      "https://atarigamer.com/lynx/game/NinjaGaiden/827588241",
	"Ninja Gaiden III - The Ancient Ship of Doom": "https://atarigamer.com/lynx/game/NinjaGaidenIIITheAncientShipofDoom/1587437967",
	"Pac-Land":                          "https://atarigamer.com/lynx/game/PacLand/849005384",
	"Paperboy":                          "https://atarigamer.com/lynx/game/Paperboy/1049850191",
	"Pinball Jam":                       "https://atarigamer.com/lynx/game/PinballJam/1337434009",
	"Pit-Fighter":                       "https://atarigamer.com/lynx/game/PitFighter/1689876186",
	"Power Factor":                      "https://atarigamer.com/lynx/game/PowerFactor/1929084752",
	"Qix":                               "https://atarigamer.com/lynx/game/Qix/358975471",
	"Rampage":                           "https://atarigamer.com/lynx/game/Rampage/780287746",
	"Rampart":                           "https://atarigamer.com/lynx/game/Rampart/1158997951",
	"RoadBlasters":                      "https://atarigamer.com/lynx/game/RoadBlasters/1758123820",
	"Robo-Squash":                       "https://atarigamer.com/lynx/game/RoboSquash/31958652",
	"Robotron 2084":                     "https://atarigamer.com/lynx/game/Robotron2084/550563919",
	"Rygar":                             "https://atarigamer.com/lynx/game/Rygar/1204550012",
	"S.T.U.N. Runner":                   "https://atarigamer.com/lynx/game/STUNRunner/486488701",
	"Scrapyard Dog":                     "https://atarigamer.com/lynx/game/ScrapyardDog/257159884",
	"Shadow of the Beast":               "https://atarigamer.com/lynx/game/ShadowoftheBeast/1261224969",
	"Shanghai":                          "https://atarigamer.com/lynx/game/Shanghai/1902366290",
	"Steel Talons":                      "https://atarigamer.com/lynx/game/SteelTalons/1174986717",
	"Super Asteroids & Missile Command": "https://atarigamer.com/lynx/game/SuperAsteroidsMissileCommand/1968158458",
	"Super Off Road":                    "https://atarigamer.com/lynx/game/SuperOffRoad/215317696",
	"Super Skweek":                      "https://atarigamer.com/lynx/game/SuperSkweek/1621976337",
	"Switchblade II":                    "https://atarigamer.com/lynx/game/SwitchbladeII/2076886345",
	"Todd's Adventures in Slime World":  "https://atarigamer.com/lynx/game/ToddsAdventuresinSlimeWorld/1301381345",
	"Toki":                              "https://atarigamer.com/lynx/game/Toki/1313984342",
	"Tournament Cyberball 2072":         "https://atarigamer.com/lynx/game/TournamentCyberball2072/1667194341",
	"Turbo Sub":                         "https://atarigamer.com/lynx/game/TurboSub/752912127",
	"Viking Child":                      "https://atarigamer.com/lynx/game/VikingChild/596449372",
	"Warbirds":                          "https://atarigamer.com/lynx/game/Warbirds/586329012",
	"World Class Fussball - Soccer":     "https://atarigamer.com/lynx/game/WorldClassFussballSoccer/990173250",
	"Xenophobe":                         "https://atarigamer.com/lynx/game/Xenophobe/1367643974",
	"Xybots":                            "https://atarigamer.com/lynx/game/Xybots/329677926",
	"Zarlor Mercenary":                  "https://atarigamer.com/lynx/game/ZarlorMercenary/1243783599",
}

var lynxLongplays = []string{
	"A.P.B. All Points Bulletin|APoxk2hmStY",
	"Awesome Golf|f7fGulSlXBc",
	"Baseball Heroes|LUF_uELADuo",
	"Basketbrawl|vuZzA_Rifvg",
	"Batman Returns|oj-0b1gFE3U",
	"Battle Wheels|rvHSEghdS3w",
	"Bill and Ted's Excellent Adventure|BYmfII3wGXw",
	"Blockout|Erq8iOQlMt4",
	"Blue Lightning|BiGqnUlsKP4",
	"Bubble Trouble|0zDUl1DIuE8",
	"California Games|GSFYx5_bdNU",
	"Checkered Flag|bOHgSxOmKeM",
	"Chip's Challenge|Qw4EFHGemfQ",
	"Desert Strike|MZIy_mjT9dI",
	"Dinolympics|6xfylVuy1KY",
	"Dirty Larry - Renegade Cop|n-p6j5OFWHI",
	"Double Dragon|fTQo6SF3M0k",
	"Dracula: The Undead|fH-wcU---fM",
	"Electrocop|FKc_QznAIJY",
	"European Soccer Challenge|uv4Ffq7FJi4",
	"Gauntlet: The Third Encounter|ZplHuUekqF8",
	"Gordo 106: The Mutated Lab Monkey|xDvjHSKXilI",
	"Hard Drivin|iJL8iA0OrQA",
	"Hockey|mrvsQqAUJtw",
	"Hydra|hEyLQoDAcSM",
	"Ishido: The Way of the Stones|o93cuoOpIf4",
	"Jimmy Connors Tennis|sEVcTBVUArk",
	"Joust|FgCNbrSwVJA",
	"Krazy Ace Miniature Golf|w9O1L7M9HHA",
	"Kung Food|iG6XlA3ijxg",
	"Lemmings|2HySrPl_lcY",
	"Lynx Casino|VzaSoU9bKs0",
	"Malibu Bikini Volleyball|eScqNV5VaWM",
	"Ms. Pac-Man|ar0DESrKquo",
	"NFL Football|OIDoZDik-LE",
	"Ninja Gaiden|khK5RD5CFOc",
	"Ninja Gaiden III - The Ancient Ship of Doom|hJxsqLw86bk",
	"Pac-Land|NBLGgP6phT4",
	"Paperboy|HohglNPEGSI",
	"Pinball Jam|etkDiZZD8Yw",
	"Pit Fighter - The Ultimate Competition|aI772oxFQ3s",
	"Power Factor|mJgDwh85_Ws",
	"Prophecy I: The Viking Child|RXVjOXe8QYA",
	"QIX|T7tQo6VX7W8",
	"Raiden (Unlicensed)|HkXA8N_n0qA",
	"Rampage|T_E3W_JGpzI",
	"Rampart|29WA1feWIpI",
	"Road Riot 4WD (Prototype)|94lnOoChAg0",
	"Roadblasters|1JBliC-dcpo",
	"Robo-Squash|f4hv6FqYl1I",
	"Robotron 2048|vomOrfaQJfU",
	"Rygar - Legendary Warrior|3nofdhw79PI",
	"S.T.U.N. Runner|1LK6R3nDM7w",
	"Scrapyard Dog|5Cz9J_ineY0",
	"Shadow of the Beast|IxGxczM78oU",
	"Shanghai|phnNtxUFKB8",
	"Super Asteroids & Missile Command|ITQtOF7Wr_s",
	"Super Off-Road|hHTMpuRJu0Q",
	"Super Skweek|VMh1Ft281Bs",
	"Switchblade II|fG8eWllV8fY",
	"The Fidelity Ultimate Chess Challenge|kcXA_PWCRJM",
	"The Gates of Zendocon|1CA-IvToZcM",
	"Toki|2BrKMKIve9U",
	"Tournament Cyberball|PJagJithBdY",
	"Turbo Sub|6bn_l3s5zDI",
	"Warbirds|2z-PKq21SIE",
	"World Class Soccer|y4a6SrejKqM",
	"Xenophobe|Wa6EpquD5T0",
	"Xybots|nZKFk72KvHU",
	"Zarlor Mercenary|oBPz6sr5NNQ",
}
