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
		generateMisterConsoleHTML("Sega CD Games", &segacdGameList, segacdImages, segacdVideos, "segacd")
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
			if lowestDistance == 3 {

				//lowestName = lowestName[strings.IndexByte(lowestName, '|'):]
				fmt.Println("\"" + v + "\": \"" + lowestName + "\",")

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

var segacdVideos = map[string]string{}

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
