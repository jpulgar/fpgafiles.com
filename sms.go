package main

func generateMisterSMSGames() {
	smsTitleAdded := make(map[string]bool)
	smsImages := make(map[string]string)
	smsGameList := []string{}
	compileMisterConsoleData(smsTitleAdded, &smsGameList, smsImages, "sms")
	generateMisterConsoleHTML("SMS", &smsGameList, smsImages, smsVideos, "sms")
}

var smsVideos = map[string]string{}
