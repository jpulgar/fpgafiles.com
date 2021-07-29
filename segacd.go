package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateMisterSegaCDGames(generate bool) {
	segacdTitleAdded := make(map[string]bool)
	segacdImages := make(map[string]string)
	segacdGameList := []string{}
	compileMisterConsoleData(segacdTitleAdded, &segacdGameList, segacdImages, segacdVideos, "segacd")
	segacdVideoReport(&segacdGameList)
	if generate {
		generateMisterConsoleHTML("Sega CD Games", &segacdGameList, segacdImages, segacdVideos, "segacd", "sega cd")
	}
}

func segacdVideoReport(gameList *[]string) {
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
		if segacdVideos[v] == "" {

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
			for _, n := range segacdLongplays {
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
			// 0, 1, 2, 3 are correct
			if lowestDistance > 3 {

				// lowestName = lowestName[strings.IndexByte(lowestName, '|')+1:]  // for final
				// fmt.Println("\"" + v + "\": \"" + lowestName + "\",") // for testing

				// For the rest without matches
				// fmt.Println("\"" + v + "\": \"\",")
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

var segacdVideos = map[string]string{
	"3 Ninjas Kick Back":                                 "mZBmz3-JdX0",
	"A-Rank Thunder - Tanjou-hen":                        "zaY6vG2kpTA",
	"A-X-101":                                            "bSfYHEst9DE",
	"Adventures of Batman and Robin, The":                "t5ebpKFC0hI",
	"Adventures of Willy Beamish, The":                   "GoQpHg7-1fg",
	"After Armageddon Gaiden - Majuu Toushouden Eclipse": "ehZHdKqUfYA",
	"After Burner III":                                   "oSMY7w5Vnrc",
	"AH3 - Thunderstrike":                                "w0NDUj3QBdA",
	"Aisle Lord":                                         "P4aG383ALTI",
	"Alshark":                                            "NbOh_JL92j8",
	"Amazing Spider-Man vs. The Kingpin, The":            "aO_qtnmti88",
	"Android Assault - The Revenge of Bari-Arm":          "QscXy9LRzpQ",
	"Animals!, The":                                      "q3Feec8YKy8",
	"Annett Futatabi":                                    "j6mcdxHSul8",
	"Aoki Ookami to Shiroki Mejika - Genchou Hishi":      "pnXklNlhmZo",
	"Arcus I, II, III":                                   "8lCe7p2uEeM",
	"Arslan Senki":                                       "bP4NHtU7ix4",
	"AX-101":                                             "bSfYHEst9DE",
	"Bakuden - The Unbalanced Zone":                      "gUXQU8Rn73Q",
	"Bari-Arm":                                           "QscXy9LRzpQ",
	"Batman Returns":                                     "ZMc0IDbkzCs",
	"Battle Fantasy":                                     "eS8J_pTSa4I",
	"Battle Frenzy":                                      "MiBNmZAAmyI",
	"Battlecorps":                                        "mE-4gpnDhWE",
	"BattleTech - Gray Death Legion":                     "Bn1TSIAtM7o",
	"BC Racers":                                          "bMEsHnTep8o",
	"Beast II":                                           "Fek9_dHM76I",
	"Bill Walsh College Football":                        "P1CADyavGjM",
	"Black Hole Assault":                                 "KgmIV24DVdU",
	"Blackhole Assault":                                  "KgmIV24DVdU",
	"Bloodshot ~ Battle Frenzy":                          "MiBNmZAAmyI",
	"Bouncers":                                           "a1Kt68Wt8UE",
	"Bram Stoker's Dracula":                              "T7HYLydpPOg",
	"Brutal - Paws of Fury":                              "I2rUc8cH6aE",
	"Bug Blasters - The Exterminators":                   "boQu-9DdL4c",
	"Burai - Yatsudama no Yuushi Densetsu":               "tn54MjkL0pQ",
	"Burning Fists - Force Striker":                      "jlgVwHsG-P8",
	"Cadillacs and Dinosaurs - The Second Cataclysm":     "7-zU-j2HQYU",
	"Captain Tsubasa":                                    "SUoFY5QhD2I",
	"CD Sonic The Hedgehog":                              "yKpS3ja7m7w",
	"Championship Soccer '94":                            "Hggy4bOqaHc",
	"Chuck Rock II - Son of Chuck":                       "n_OSkWBL2ug",
	"Chuck Rock":                                         "lVecljI4Jao",
	"Citizen X":                                          "DLmzdsZ12QI",
	"Cliffhanger":                                        "h4W2XjGx6JE",
	"Cobra Command":                                      "0V5gboD9AIw",
	"Colors of Modern Rock, The":                         "ymrxb31N_VU",
	"Compton's Interactive Encyclopedia":                 "bcLKEjWwLHY",
	"Corpse Killer":                                      "e0JgUvRPAac",
	"Cosmic Fantasy Stories":                             "e7flypM0nH4",
	"Crime Patrol":                                       "VM8iCkWkF6U",
	"Cyborg 009":                                         "VuzRDtmiLAQ",
	"Daihoushinden":                                      "tr0qBCzxEjA",
	"Dark Wizard - Yomigaerishi Yami no Madoushi":        "RA1XLA0z_lU",
	"Dark Wizard":                                        "RA1XLA0z_lU",
	"Death Bringer":                                      "xGtr27gx4E",
	"Demolition Man":                                     "Q6sZzaUC1cI",
	"Dennin Aleste - Nobunaga and His Ninja Force":       "Qcp8M-psQmE",
	"Detonator Orgun":                                    "IZwJBoR84jA",
	"Devastator":                                         "taRwRQ-En10",
	"Double Switch":                                      "jvM3F4dyuiY",
	"Dracula Unleashed":                                  "oTNC53dSi9M",
	"Dragon's Lair":                                      "XXgAfyVYN2c",
	"Dune":                                               "MDHA7TIidUI",
	"Dungeon Explorer":                                   "Jnc9coWJKdE",
	"Dungeon Master II - Skullkeep":                      "R1e3tWGSEYE",
	"Dynamic Country Club":                               "zP_6TinX25c",
	"Earnest Evans":                                      "iE-tcM-yUZY",
	"Earthworm Jim - Special Edition":                    "m5xnzSfHi2M",
	"Ecco - The Tides of Time":                           "xJsbW4Nu9sw",
	"Ecco the Dolphin CD":                                "PCPIJy0Rjiw",
	"Ecco the Dolphin":                                   "PCPIJy0Rjiw",
	"Egawa Suguru no Super League CD":                    "FhY0SGZ8TXQ",
	"ESPN Baseball Tonight":                              "Icj4uWV-juo",
	"ESPN National Hockey Night":                         "38MTYl7ZDrY",
	"ESPN NBA Hangtime '95":                              "96m3uRw6ndc",
	"ESPN Sunday Night NFL":                              "P77-IEUpcxk",
	"Eternal Champions - Challenge from the Dark Side":   "Z4CHZccSueo",
	"Eye of the Beholder":                                "3-vyaIQ7udk",
	"F1 Circus CD":                                       "jcb9bPbSrw4",
	"Fahrenheit":                                         "OHcGK6aFij4",
	"Fatal Fury Special":                                 "bwVVICHakyU",
	"FIFA International Soccer - Championship Edition":   "hofuau0lGU8",
	"FIFA International Soccer":                          "Apg9iA5YIBs",
	"Final Fight CD":                                     "NeWkue8axG8",
	"Flashback - The Quest for Identity":                 "u5mtdJ65MrQ",
	"Flashback":                                          "u5mtdJ65MrQ",
	"Flink":                                              "zwvuUwYBgvc",
	"Formula One World Championship - Beyond the Limit":             "g-P-j51xLPE",
	"Game no Kandume Vol. 1":                                        "k7NUR_HM8MY",
	"Game no Kandume Vol. 2":                                        "z86ziWw8TGg",
	"Garou Densetsu Special":                                        "oZzV3FXjl94",
	"Gen'ei Toshi - Illusion City":                                  "3QxvsTu-2sI",
	"Ground Zero Texas":                                             "fpREyMjFoYc",
	"Gyuwambler Jiko Chuushinha 2 - Gekitou! Tokyo Mahjongland Hen": "aqE7ZfQx2XY",
	"Heart of the Alien - Out of This World Parts I and II":         "cW4RBhXPRUY",
	"Heavenly Symphony - Formula One World Championship 1993":       "HTeFfYvjqMo",
	"Heavy Nova":                                     "YcoPY-_9GYM",
	"Heimdall":                                       "_4No4sBH7wY",
	"Hook":                                           "3rwtmJKsRP4",
	"IIIrd World War, The":                           "Wil0TxdL5Nc",
	"Iron Helix":                                     "UG56ULXrLrs",
	"Ishii Hisaichi no Daiseikai":                    "HNG4RC5y8dQ",
	"Jaguar XJ220":                                   "2l5gvlXi4PU",
	"Jangou World Cup":                               "2KDcfRH6NFY",
	"Jeopardy!":                                      "auqwRQ3rdo4",
	"Joe Montana's NFL Football":                     "yASgym5uyTc",
	"Jurassic Park":                                  "685maHqYP5U",
	"Kamen Rider ZO":                                 "FO22q0s8bAM",
	"Keio Flying Squadron":                           "CJI6eoUC6o8",
	"Keiou Yuugekitai":                               "ttFV0LBA-_I",
	"Kids on Site":                                   "vNmBxJ_QnT8",
	"Lawnmower Man, The":                             "cPzx9LhJm10",
	"Lethal Enforcers II - Gun Fighters":             "o4PNEBwXAIY",
	"Lethal Enforcers II - The Western":              "o4PNEBwXAIY",
	"Lethal Enforcers":                               "uQvxzlwZkzM",
	"Links - The Challenge of Golf":                  "s1ocoSvssEA",
	"Loadstar - The Legend of Tully Bodine":          "ipXiS9mBwk0",
	"Lords of Thunder":                               "wTiuxitLTII",
	"Lunar - Eternal Blue":                           "7yAfF0ARHx0",
	"Lunar - The Silver Star":                        "xeq9fPOxz6o",
	"Mad Dog II - The Lost Gold":                     "9HWqs-kZkDA",
	"Mad Dog McCree":                                 "pWA_HIRWn_U",
	"Mahou no Shoujo - Silky Lip":                    "n9Aa54Be0RI",
	"Make My Video - INXS":                           "0Gjp-wntDhw",
	"Make My Video - Kris Kross":                     "busD3r36Z2Y",
	"Make My Video - Marky Mark and the Funky Bunch": "7P15yI2wYL8",
	"Mansion of Hidden Souls":                        "9NWxCx2gxME",
	"Marko's Magic Football":                         "ehE9eO4_l5Y",
	"Mary Shelley's Frankenstein":                    "cA0VgqxSMgU",
	"Masked Rider, The - Kamen Rider ZO":             "FO22q0s8bAM",
	"Mega Schwarzschild":                             "jTy3LeGqNO4",
	"MegaRace":                                       "XR2Y1h5LDzU",
	"Mickey Mania - The Timeless Adventures of Mickey Mouse": "EKZ72i3SkKg",
	"Mickey Mania":                         "EKZ72i3SkKg",
	"Microcosm":                            "M_av6C1UVqg",
	"Midnight Raiders":                     "o0C20hgsz9Y",
	"Might and Magic III - Isles of Terra": "EzuHa3rH7IM",
	"Mighty Mighty Missile":                "Qw0WKMvj0CA",
	"Mighty Morphin Power Rangers":         "aeeFW74X87Y",
	"Mortal Kombat Kanzenban":              "BOCS56eWljU",
	"Mortal Kombat":                        "BOCS56eWljU",
	"My Paint":                             "ISGxbmotulo",
	"NBA Jam":                              "L-_89qkuRI0",
	"NFL Football Trivia Challenge":        "nL5ai5u9ZrY",
	"NFL's Greatest - San Francisco vs. Dallas 1978-1993": "62-bEo0sOT4",
	"NHL '94":                     "VacTXWt4UiU",
	"NHL Hockey '94":              "VacTXWt4UiU",
	"Night Striker":               "9p9NfFSCWzc",
	"Night Trap":                  "plV5l_uKzDo",
	"Ninja Warriors, The":         "Mp26di7Qz8w",
	"Nobunaga no Yabou - Haouden": "Irfcp81xd9M",
	"Nostalgia 1907":              "svo_H38tHuM",
	"Note! Color Mechanica":       "CIfHXe75rSk",
	"Novastorm":                   "P6areZdswUI",
	"Panic!":                      "5ccDvvt8l_U",
	"Pier Solar and the Great Architects Enhanced Soundtrack Disc": "EhrC1iJg1TU",
	"Pitfall - The Mayan Adventure":                                "P86I5GgsLZM",
	"Popful Mail":                                                  "tIU9YHdhY1U",
	"PopfulMail":                                                   "tIU9YHdhY1U",
	"Power Factory":                                                "Y7jSwSlMuu0",
	"Power Monger":                                                 "4Zub2oMLX9s",
	"PowerMonger":                                                  "4Zub2oMLX9s",
	"Prince of Persia":                                             "9CYrG--7roM",
	"Prize Fighter":                                                "FPcnTpqjX6M",
	"Pro Yakyuu Super League CD":                                   "gMbdsYq8gdU",
	"Psychic Detective Series Vol. 3 - Aya":                        "jvRJJ-e3Fv4",
	"Psychic Detective Series Vol. 4 - Orgel":                      "bCWCVrv0gUM",
	"Puggsy":                                      "xCzEwdqjJsM",
	"Quiz Scramble Special":                       "vfzhVncK-NM",
	"Quiz Tonosama no Yabou":                      "w9nKTIQC5pE",
	"Racing Aces":                                 "O_RNCaIWT7o",
	"Radical Rex":                                 "FBmlW-uqAJ4",
	"Ranma 1-2 - Byakuranaika":                    "cz_1eYOJfmA",
	"RDF - Global Conflict":                       "G34v717juKM",
	"Record of Lodoss War":                        "kwcIs0ix0vo",
	"Revenge of the Ninja":                        "uEtWncBgrBk",
	"Revengers of Vengeance":                      "SIzOGf7Xu_w",
	"Rise of the Dragon - A Blade Hunter Mystery": "Cbmn37XngHc",
	"Rise of the Dragon":                          "Cbmn37XngHc",
	"Road Avenger":                                "5LK2KcUu6po",
	"Road Blaster FX":                             "Y1gIwmsyV6g",
	"Road Rash":                                   "Sugdc7D4zqM",
	"Robo Aleste":                                 "IOa8m2zOtBs",
	"Samurai Shodown":                             "nEqImJE2AGs",
	"Sangokushi III":                              "IXjCKHLjQuI",
	"Secret of Monkey Island, The":                "HERilq6-9JQ",
	"Sega Classic Arcade Collection - Limited Edition":  "8BbzmgskrOI",
	"Sega Classics Arcade Collection - Limited Edition": "8BbzmgskrOI",
	"Sega Classics Arcade Collection":                   "8BbzmgskrOI",
	"Seima Densetsu 3x3 Eyes":                           "VQT3EiymD6Q",
	"Seirei Shinseiki - Fhey Area":                      "HRV1utTlZRI",
	"Sengoku Denshou":                                   "tUudJpXy7fo",
	"Sensible Soccer":                                   "n9rVc0A6Utw",
	"Sewer Shark":                                       "A35QKz6cFJw",
	"Shadow of the Beast II - Juushin no Jubaku":        "Fek9_dHM76I",
	"Shadow of the Beast II":                            "Fek9_dHM76I",
	"Shadowrun":                                         "UKzt8YU24OU",
	"Sherlock Holmes - Consulting Detective Vol. II":    "2uRyxGPuYkk",
	"Sherlock Holmes - Consulting Detective":            "tMLpZROE27w",
	"Shin Megami Tensei":                                "fY0iPiHjsEY",
	"Shining Force CD":                                  "F2SSYt_ybJI",
	"Silpheed":                                          "Crjgu2aE1QE",
	"SimEarth":                                          "cElaQBRsEGU",
	"Sing!! Sega Game Music Presented by B. B. Queens":  "BLnQQ7GMlB0",
	"Slam City with Scottie Pippen":                     "WxRChVF0q9U",
	"Smurfs, The":                                       "qXkGQk0r7O0",
	"Snatcher":                                          "OsNQiaZZyLI",
	"Software Toolworks' Star Wars Chess, The":          "yF2N2kuQYPo",
	"Sol-Feace":                                         "BMIw9BcXwwo",
	"Sonic CD":                                          "oCKuWVGkdrU",
	"Sonic MegaMix":                                     "4_HtGuob62c",
	"Sonic The Hedgehog CD":                             "oCKuWVGkdrU",
	"SoulStar & Battlecorps":                            "rObi13jlNy4",
	"SoulStar":                                          "rObi13jlNy4",
	"Space Ace":                                         "6K8l2bK1s8s",
	"Space Adventure, The - Cobra - The Legendary Bandit": "8Zc0WO_1yXw",
	"Star Wars - Rebel Assault":                           "F6yOX2vmxgY",
	"Starblade":                                           "UdBd5EP-09o",
	"Stellar-Fire":                                        "WQVlp5ofrr4",
	"Super Strike Trilogy":                                "cRqIwvzHSO4",
	"Supreme Warrior":                                     "Lyak7VQhW4o",
	"Surgical Strike":                                     "eNT95--KgS0",
	"Switch":                                              "5ccDvvt8l_U",
	"Syndicate":                                           "Vj8SFATpmyk",
	"Tenbu Mega CD Special":                               "N7NuCBxJ7_g",
	"Tenkafubu - Eiyuutachi no Houkou":                    "znA-vG20_Ds",
	"Terminator, The":                                     "VZdY96eEsTU",
	"Theme Park":                                          "GHLBUR8s9kU",
	"Third World War":                                     "ROlu_XdJvs4",
	"Thunder Storm FX":                                    "K8o2iR2FRdM",
	"Thunderhawk":                                         "oPS6ObO7ZdM",
	"Time Gal":                                            "SNNKAhSYsoA",
	"TimeCop":                                             "j22sC-65Ud4",
	"Tomcat Alley":                                        "WzAvFLNVgEI",
	"Trivial Pursuit":                                     "tOawKReO1jU",
	"Ultraverse Prime":                                    "0y55dYdUWZk",
	"Urusei Yatsura - Dear My Friends":                    "Feu2bYTcXoE",
	"Vay - Ryuusei no Yoroi":                              "Yijg3dlD5xY",
	"Vay":                                                 "Yijg3dlD5xY",
	"Wakusei Woodstock - Funky Horror Band":               "o8WpwnmFLBg",
	"Warau Salesman":                                      "cbBJWLQTMMs",
	"Wheel of Fortune":                                    "WoX8dYrqLPs",
	"Who Shot Johnny Rock":                                "o8AoltzULpQ",
	"Wild Woody":                                          "EkTC335FEz0",
	"Wing Commander":                                      "7Fm3zZClHsU",
	"Winning Post":                                        "yjS6_0aqoz8",
	"WireHead":                                            "aO_KFHEWnBY",
	"Wolfchild":                                           "o7bSl7mzfMA",
	"Wonder Dog":                                          "zcAMtNMOAcA",
	"WonderMega Collection":                               "0XB4mGn-9TQ",
	"World Cup USA '94":                                   "420iByRiqxY",
	"WWF - Rage in the Cage":                              "QLKbMHBskRQ",
	"WWF Mania Tour - WWF - Rage in the Cage":             "QLKbMHBskRQ",
	"Yumemi Mystery Mansion":                              "rNJAprfXsK4",
	"Yumemi Yakata no Monogatari":                         "u_tyOlyppPY",
	"Yumimi Mix":                                          "w_qMlWjzdVQ",
}

var segacdLongplays = []string{
	"3 Ninjas kick back|mZBmz3-JdX0",
	"A/X-101|bSfYHEst9DE",
	"Adventures of Batman and Robin|t5ebpKFC0hI",
	"After Burner III|oSMY7w5Vnrc",
	"Android Assault: The Revenge of Bari-Arm|QscXy9LRzpQ",
	"Annet Futatabi|j6mcdxHSul8",
	"Batman Returns|ZMc0IDbkzCs",
	"Battle Frenzy|MiBNmZAAmyI",
	"BC Racers|bMEsHnTep8o",
	"Black Hole Assault|KgmIV24DVdU",
	"Bouncers|a1Kt68Wt8UE",
	"Bram Stokers Dracula|T7HYLydpPOg",
	"Brutal: Paws of Fury|I2rUc8cH6aE",
	"Bug Blasters: The Exterminators (Prototype)|boQu-9DdL4c",
	"Burning Fists: Force Striker (Prototype)|jlgVwHsG-P8",
	"Cadillacs and Dinosaurs: The Second Cataclysm|7-zU-j2HQYU",
	"Capcom no Quiz - Tonosama no Yabou|w9nKTIQC5pE",
	"Championship Soccer '94|Hggy4bOqaHc",
	"Chuck Rock|lVecljI4Jao",
	"Chuck Rock II - Son Of Chuck|n_OSkWBL2ug",
	"Citizen X (Prototype)|DLmzdsZ12QI",
	"Cliffhanger|h4W2XjGx6JE",
	"Cobra Command|0V5gboD9AIw",
	"Corpse Killer|e0JgUvRPAac",
	"Crime Patrol|VM8iCkWkF6U",
	"Cyborg 009|VuzRDtmiLAQ",
	"Dark Wizard|RA1XLA0z_lU",
	"Demolition Man|Q6sZzaUC1cI",
	"Devastator|taRwRQ-En10",
	"Double Switch|jvM3F4dyuiY",
	"Dracula Unleashed|oTNC53dSi9M",
	"Dragon's Lair|XXgAfyVYN2c",
	"Dune|MDHA7TIidUI",
	"Earnest Evans|iE-tcM-yUZY",
	"Earthworm Jim Special Edition|m5xnzSfHi2M",
	"Ecco the Dolphin|PCPIJy0Rjiw",
	"Ecco: The Tides of Time|xJsbW4Nu9sw",
	"Eternal Champions: Challenge from the Dark Side|Z4CHZccSueo",
	"Fahrenheit|OHcGK6aFij4",
	"Fatal Fury Special|bwVVICHakyU",
	"FIFA International Soccer|Apg9iA5YIBs",
	"Final Fight CD|NeWkue8axG8",
	"Final Fight CD (a)|T8cBEOqLf1o",
	"Flashback|u5mtdJ65MrQ",
	"Garou Densetsu Special|oZzV3FXjl94",
	"Ground Zero Texas|fpREyMjFoYc",
	"Heart of the Alien|cW4RBhXPRUY",
	"Heavy Nova|YcoPY-_9GYM",
	"Hook|3rwtmJKsRP4",
	"Jaguar XJ220|2l5gvlXi4PU",
	"Jeopardy|auqwRQ3rdo4",
	"Jurassic Park|685maHqYP5U",
	"Keio Flying Squadron|CJI6eoUC6o8",
	"Keio Yuugekitai|ttFV0LBA-_I",
	"Kids on Site|vNmBxJ_QnT8",
	"Lethal Enforcers|uQvxzlwZkzM",
	"Lethal Enforcers II - Gun Fighters|o4PNEBwXAIY",
	"Loadstar - The Legend of Tully Bodine|ipXiS9mBwk0",
	"Lords of Thunder|wTiuxitLTII",
	"Lunar - The Silver Star|xeq9fPOxz6o",
	"Lunar II: Eternal Blue|7yAfF0ARHx0",
	"Mad Dog II - The Lost Gold|9HWqs-kZkDA",
	"Mad Dog McCree|pWA_HIRWn_U",
	"Magical Popful Mail Fantasy Adventure|tIU9YHdhY1U",
	"Make My Video: Power Factory, Featuring C+C Music Factory|Y7jSwSlMuu0",
	"Mansion of Hidden Souls|9NWxCx2gxME",
	"Marko (Prototype)|ehE9eO4_l5Y",
	"Mary Shelleys Frankenstein|cA0VgqxSMgU",
	"Mega Race|XR2Y1h5LDzU",
	"Mickey Mania: The Timeless Adventures of Mickey Mouse|EKZ72i3SkKg",
	"Microcosm|M_av6C1UVqg",
	"Midnight Raiders|o0C20hgsz9Y",
	"Mighty Morphin Power Rangers|aeeFW74X87Y",
	"Mortal Kombat|BOCS56eWljU",
	"Night Striker|9p9NfFSCWzc",
	"Night Trap|plV5l_uKzDo",
	"Novastorm|P6areZdswUI",
	"Panic!|5ccDvvt8l_U",
	"Pitfall - The Mayan Adventure|P86I5GgsLZM",
	"Prince of Persia|9CYrG--7roM",
	"Prize Fighter|FPcnTpqjX6M",
	"Puggsy|xCzEwdqjJsM",
	"Radical Rex|FBmlW-uqAJ4",
	"Revenge of the Ninja|uEtWncBgrBk",
	"Revengers of Vengeance|SIzOGf7Xu_w",
	"Rise of the Dragon|Cbmn37XngHc",
	"Road Avenger|5LK2KcUu6po",
	"Road Rash|Sugdc7D4zqM",
	"Robo Aleste|IOa8m2zOtBs",
	"Samurai Shodown|nEqImJE2AGs",
	"Sega Classics Arcade Collection 5-in-1 - Columns|OtAl4Q5KazY",
	"Sega Classics Arcade Collection 5-in-1 - Golden Axe|jQp_CrHslIc",
	"Sega Classics Arcade Collection 5-in-1 - Streets of Rage|F4VQHV6hGJ0",
	"Sega Classics Arcade Collection 5-in-1 - Super Monaco GP|b_IL7z7P7b0",
	"Sega Classics Arcade Collection 5-in-1 - The Revenge of Shinobi|8BbzmgskrOI",
	"Sengoku Densyo|tUudJpXy7fo",
	"Sewer Shark|A35QKz6cFJw",
	"Shadow of the Beast II|Fek9_dHM76I",
	"Shining Force CD|F2SSYt_ybJI",
	"Silpheed|Crjgu2aE1QE",
	"Slam City with Scottie Pippen|WxRChVF0q9U",
	"Snatcher|OsNQiaZZyLI",
	"Sol-Feace|BMIw9BcXwwo",
	"Sonic CD|oCKuWVGkdrU",
	"Sonic the Hedgehog CD (a)|yKpS3ja7m7w",
	"Soulstar|rObi13jlNy4",
	"Star Strike (Prototype)|aVJoU6HoH-U",
	"Star Wars Chess|yF2N2kuQYPo",
	"Star Wars: Rebel Assault|F6yOX2vmxgY",
	"Starblade|UdBd5EP-09o",
	"Stellar Fire|WQVlp5ofrr4",
	"Supreme Warrior|Lyak7VQhW4o",
	"Surgical Strike|eNT95--KgS0",
	"The Adventures of Willy Beamish|GoQpHg7-1fg",
	"The Amazing Spider-Man vs. The Kingpin|aO_qtnmti88",
	"The Lawnmower Man|cPzx9LhJm10",
	"The Masked Rider: Kamen Rider ZO|FO22q0s8bAM",
	"The Misadventures of Flink|zwvuUwYBgvc",
	"The Ninja Warriors|Mp26di7Qz8w",
	"The Secret of Monkey Island|HERilq6-9JQ",
	"The Smurfs|qXkGQk0r7O0",
	"The Space Adventure - Cobra The Legendary Bandit|8Zc0WO_1yXw",
	"The Terminator|VZdY96eEsTU",
	"Thunderhawk|oPS6ObO7ZdM",
	"Time Cop (Prototype)|j22sC-65Ud4",
	"Time Gal|SNNKAhSYsoA",
	"Time Gal (a)|drqOpfJ2VBM",
	"Tomcat Alley|WzAvFLNVgEI",
	"Trivial Pursuit|tOawKReO1jU",
	"Ultraverse Prime|0y55dYdUWZk",
	"Urusei Yatsura: Dear My Friends|Feu2bYTcXoE",
	"Vay|Yijg3dlD5xY",
	"Warau Salesman|cbBJWLQTMMs",
	"Wheel of Fortune|WoX8dYrqLPs",
	"Who Shot Johnny Rock|o8AoltzULpQ",
	"Wild Woody|EkTC335FEz0",
	"Wing Commander|7Fm3zZClHsU",
	"Wirehead|aO_KFHEWnBY",
	"Wolfchild|o7bSl7mzfMA",
	"Wonder Dog|zcAMtNMOAcA",
	"World Cup USA '94|420iByRiqxY",
	"WWF Rage in the Cage|QLKbMHBskRQ",
	"Yumimi Mix|w_qMlWjzdVQ",
}
