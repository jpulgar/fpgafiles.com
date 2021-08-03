package main

func generateMisterSG1000Games(generate bool) {
	sg1000TitleAdded := make(map[string]bool)
	sg1000Images := make(map[string]string)
	sg1000GameList := []string{}
	compileMisterConsoleData(sg1000TitleAdded, &sg1000GameList, sg1000Images, sg1000Videos, "sg1000")
	// sg1000VideoReport(&sg1000GameList)
	if generate {
		generateMisterConsoleHTML("Sega SG-1000", &sg1000GameList, sg1000Images, sg1000Videos, "sg1000", "sg-1000")
	}
	// for k := range sg1000GameInfo {
	// 	data, err := ioutil.ReadFile("public/rips/smspower.org-sg1000/" + k + ".html")
	// 	if err != nil {
	// 		fmt.Println("File reading error", err)
	// 		return
	// 	}
	// 	re := regexp.MustCompile(`(?m)^<div><span class='shot'>.*</span></div>$`)
	// 	results := re.FindAllString(string(data), 2)
	// 	for kk, vv := range results {
	// 		iStart := strings.Index(vv, "src='")
	// 		iEnd := strings.Index(vv, "' alt=''")
	// 		fmt.Println("Saving: https:" + vv[iStart+5:iEnd])
	// 		url := "https:" + vv[iStart+5:iEnd]
	// 		if kk == 0 {
	// 			downloadFile(map[string]string{k: url}, "public/mister/sg1000/titles", ".png", 4)
	// 		} else {
	// 			downloadFile(map[string]string{k: url}, "public/mister/sg1000/snaps", ".png", 4)
	// 		}
	// 	}
	// }
}

var sg1000Videos = map[string]string{}

var sg1000GameInfo = map[string]string{
	"Bank Panic":                    "https://www.smspower.org/Games/BankPanic-SG",
	"The Black Onyx":                "https://www.smspower.org/Games/BlackOnyx-SG",
	"Bomberman Special":             "https://www.smspower.org/Games/BombermanSpecial-SG",
	"Bomb Jack":                     "https://www.smspower.org/Games/BombJack-SG",
	"Cabbage Patch Kids":            "https://www.smspower.org/Games/CabbagePatchKids-SG",
	"The Castle":                    "https://www.smspower.org/Games/Castle-SG",
	"The Castle [MSX]":              "https://www.smspower.org/Games/CastleMsx-SG",
	"Chack'n Pop":                   "https://www.smspower.org/Games/ChacknPop-SG",
	"Champion Baseball":             "https://www.smspower.org/Games/ChampionBaseball-SG",
	"Champion Billiards":            "https://www.smspower.org/Games/ChampionBilliards-SG",
	"Champion Boxing":               "https://www.smspower.org/Games/ChampionBoxing-SG",
	"Champion Golf":                 "https://www.smspower.org/Games/ChampionGolf-SG",
	"Champion Ice Hockey":           "https://www.smspower.org/Games/ChampionIceHockey-SG",
	"Champion Kendou":               "https://www.smspower.org/Games/ChampionKendou-SG",
	"Champion Pro Wrestling":        "https://www.smspower.org/Games/ChampionProWrestling-SG",
	"Championship Lode Runner":      "https://www.smspower.org/Games/ChampionshipLodeRunner-SG",
	"Champion Soccer":               "https://www.smspower.org/Games/ChampionSoccer-SG",
	"Champion Tennis":               "https://www.smspower.org/Games/ChampionTennis-SG",
	"Choplifter":                    "https://www.smspower.org/Games/Choplifter-SG",
	"Circus Charlie":                "https://www.smspower.org/Games/CircusCharlie-SG",
	"Congo Bongo":                   "https://www.smspower.org/Games/CongoBongo-SG",
	"C_So!":                         "https://www.smspower.org/Games/CSo-SG",
	"Doki Doki Penguin Land":        "https://www.smspower.org/Games/DokiDokiPenguinLand-SG",
	"Dragon Wang":                   "https://www.smspower.org/Games/DragonWang-SG",
	"Drol":                          "https://www.smspower.org/Games/Drol-SG",
	"Elevator Action":               "https://www.smspower.org/Games/ElevatorAction-SG",
	"Exerion":                       "https://www.smspower.org/Games/Exerion-SG",
	"Flicky":                        "https://www.smspower.org/Games/Flicky-SG",
	"Girl's Garden":                 "https://www.smspower.org/Games/GirlsGarden-SG",
	"Golgo 13":                      "https://www.smspower.org/Games/Golgo13-SG",
	"GP World":                      "https://www.smspower.org/Games/GPWorld-SG",
	"Gulkave":                       "https://www.smspower.org/Games/Gulkave-SG",
	"Hang On II":                    "https://www.smspower.org/Games/HangOnII-SG",
	"H.E.R.O.":                      "https://www.smspower.org/Games/HERO-SG",
	"Home Mahjong":                  "https://www.smspower.org/Games/HomeMahjong-SG",
	"Hustle Chumy":                  "https://www.smspower.org/Games/HustleChumy-SG",
	"Hyper Sports":                  "https://www.smspower.org/Games/HyperSports-SG",
	"Hyper Sports 2":                "https://www.smspower.org/Games/HyperSports2-SG",
	"King's Valley":                 "https://www.smspower.org/Games/KingsValley-SG",
	"Knightmare":                    "https://www.smspower.org/Games/Knightmare-SG",
	"The Legend of Kage":            "https://www.smspower.org/Games/LegendOfKage-SG",
	"Lode Runner":                   "https://www.smspower.org/Games/LodeRunner-SG",
	"Magical Kid Wiz":               "https://www.smspower.org/Games/MagicalKidWiz-SG",
	"Magical Tree":                  "https://www.smspower.org/Games/MagicalTree-SG",
	"Mahjong":                       "https://www.smspower.org/Games/Mahjong-SG",
	"Monaco GP":                     "https://www.smspower.org/Games/MonacoGP-SG",
	"Ninja Princess":                "https://www.smspower.org/Games/NinjaPrincess-SG",
	"N-Sub":                         "https://www.smspower.org/Games/NSub-SG",
	"Othello":                       "https://www.smspower.org/Games/Othello-SG",
	"Pacar":                         "https://www.smspower.org/Games/Pacar-SG",
	"Pachinko":                      "https://www.smspower.org/Games/Pachinko-SG",
	"Pachinko II":                   "https://www.smspower.org/Games/PachinkoII-SG",
	"Konami's Ping Pong":            "https://www.smspower.org/Games/PingPong-SG",
	"Pippols":                       "https://www.smspower.org/Games/Pippols-SG",
	"Pitfall II ~The Lost Caverns~": "https://www.smspower.org/Games/PitfallII-SG",
	"Pop Flamer":                    "https://www.smspower.org/Games/PopFlamer-SG",
	"Rally-X":                       "https://www.smspower.org/Games/RallyX-SG",
	"Road Fighter":                  "https://www.smspower.org/Games/RoadFighter-SG",
	"Rock n' Bolt":                  "https://www.smspower.org/Games/RocknBolt-SG",
	"Safari Hunting":                "https://www.smspower.org/Games/SafariHunting-SG",
	"Safari Race":                   "https://www.smspower.org/Games/SafariRace-SG",
	"Sega Flipper":                  "https://www.smspower.org/Games/SegaFlipper-SG",
	"Sega-Galaga":                   "https://www.smspower.org/Games/SegaGalaga-SG",
	"SG-1000 M2 Check Program":      "https://www.smspower.org/Games/SG1000M2CheckProgram-SG",
	"Shinnyuushain Tooru-Kun":       "https://www.smspower.org/Games/ShinnyuushainTooruKun-SG",
	"Sindbad Mystery":               "https://www.smspower.org/Games/SindbadMystery-SG",
	"Soukoban":                      "https://www.smspower.org/Games/Soukoban-SG",
	"Space Invaders":                "https://www.smspower.org/Games/SpaceInvaders-SG",
	"Space Slalom":                  "https://www.smspower.org/Games/SpaceSlalom-SG",
	"Star Force":                    "https://www.smspower.org/Games/StarForce-SG",
	"Star Jacker":                   "https://www.smspower.org/Games/StarJacker-SG",
	"Star Soldier":                  "https://www.smspower.org/Games/StarSoldier-SG",
	"Super Tank":                    "https://www.smspower.org/Games/SuperTank-SG",
	"Tank Battalion":                "https://www.smspower.org/Games/TankBattalion-SG",
	"Terebi Oekaki":                 "https://www.smspower.org/Games/TerebiOekaki-SG",
	"TwinBee":                       "https://www.smspower.org/Games/TwinBee-SG",
	"Wonder Boy":                    "https://www.smspower.org/Games/WonderBoy-SG",
	"Yamato":                        "https://www.smspower.org/Games/Yamato-SG",
	"Yie Ar Kung-Fu":                "https://www.smspower.org/Games/YieArKungFu-SG",
	"Yie Ar Kung-Fu II":             "https://www.smspower.org/Games/YieArKungFuII-SG",
	"Zaxxon":                        "https://www.smspower.org/Games/Zaxxon-SG",
	"Zippy Race":                    "https://www.smspower.org/Games/ZippyRace-SG",
	"Zoom 909":                      "https://www.smspower.org/Games/Zoom909-SG",
}
