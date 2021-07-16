window.addEventListener('DOMContentLoaded', (event) => {
    
    // Add Navigation
    fetch('/navigation.json')
    .then(response => response.json())
    .then(data => {
        const navTarget = document.querySelector('.navcontent');
        for (var company in data.sections) {
            navTarget.append(company + ": ")
            var companyLength = Object.keys(data.sections[company]).length;
            var i = 1;
            for (var system in data.sections[company]) {
                var shortname = data.sections[company][system].shortname;
                var longname = data.sections[company][system].longname;
                var link = data.sections[company][system].link;
                var a = document.createElement("a")
                a.append(shortname)
                a.setAttribute("href", link)
                navTarget.append(a)
                if (i != companyLength) {
                    navTarget.append(" | ");
                }
                i++;
            }
            let br = document.createElement("br")
            navTarget.append(br)
        }
    });

    
    // Credits
    var pathname = window.location.pathname;
    var directory = pathname.substring(0, pathname.lastIndexOf('/')) + "/";
    if (directory == "/mister/arcade/") {
        showCredits("Antonio Paradossi", "https://www.progettosnaps.net", "https://www.paypal.com/paypalme/progettoSNAPS", "progetto-SNAPS");
    } else if (directory == "/mister/lynx/") {
        showCredits("", "https://atarigamer.com", "https://atarigamer.com/pages/support-atari-gamer", "Atari Gamer");
    } else if (directory == "/mister/pce/") {
        showCredits("EmuMovies", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    } else if (directory == "/mister/nes/") {
        showCredits("Jardavius", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    } else if (directory == "/mister/gbc/") {
        showCredits("Jardavius", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    } else if (directory == "/mister/snes/") {
        showCredits("Jardavius", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    } else if (directory == "/mister/sms/") {
        showCredits("", "https://www.smspower.org", "https://www.smspower.org/Home/Donate", "SMS Power");
    } else if (directory == "/mister/genesis/") {
        showCredits("EmuMovies", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    } else if (directory == "/mister/segacd/") {
        showCredits("EmuMovies", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    } else if (directory == "/mister/atari2600/") {
        showCredits("EmuMovies", "https://emumovies.com", "https://emumovies.com/subscriptions/", "EmuMovies");
    }
});

// Credits Helper Function
function showCredits(author, url, donate, sitename) {
    const creditsTarget = document.querySelector('#credits');
    creditsTarget.append("Game images from ");
    if (author != "") {
        creditsTarget.append(author + " @ ");
    }
    var a1 = document.createElement("a");
    a1.setAttribute("href", url);
    a1.append(url);
    creditsTarget.append(a1);
    creditsTarget.append(".");
    if (donate != "") {
        creditsTarget.append(" Please consider ");
        var a2 = document.createElement("a");
        a2.setAttribute("href", donate);
        a2.append("donating");
        creditsTarget.append(a2);
        creditsTarget.append(" to " + sitename + ".");
    }
    creditsTarget.append(document.createElement("br"))
    creditsTarget.append("All copyright to these images are held by the companies who developed and published these games.");
}