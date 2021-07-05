function showGrid() {
    const myTarget = document.querySelector('#num');
    const aTarget = document.querySelector('#A');
    const bTarget = document.querySelector('#B');
    const cTarget = document.querySelector('#C');
    const dTarget = document.querySelector('#D');
    const eTarget = document.querySelector('#E');
    const fTarget = document.querySelector('#F');
    const gTarget = document.querySelector('#G');
    const hTarget = document.querySelector('#H');
    const iTarget = document.querySelector('#I');
    const jTarget = document.querySelector('#J');
    const kTarget = document.querySelector('#K');
    const lTarget = document.querySelector('#L');
    const mTarget = document.querySelector('#M');
    const nTarget = document.querySelector('#N');
    const oTarget = document.querySelector('#O');
    const pTarget = document.querySelector('#P');
    const qTarget = document.querySelector('#Q');
    const rTarget = document.querySelector('#R');
    const sTarget = document.querySelector('#S');
    const tTarget = document.querySelector('#T');
    const uTarget = document.querySelector('#U');
    const vTarget = document.querySelector('#V');
    const wTarget = document.querySelector('#W');
    const xTarget = document.querySelector('#X');
    const yTarget = document.querySelector('#Y');
    const zTarget = document.querySelector('#Z');

    const a1971Target = document.querySelector('#s1971');
    const a1972Target = document.querySelector('#s1972');
    const a1973Target = document.querySelector('#s1973');
    const a1974Target = document.querySelector('#s1974');
    const a1975Target = document.querySelector('#s1975');
    const a1976Target = document.querySelector('#s1976');
    const a1977Target = document.querySelector('#s1977');
    const a1978Target = document.querySelector('#s1978');
    const a1979Target = document.querySelector('#s1979');
    const a1980Target = document.querySelector('#s1980');
    const a1981Target = document.querySelector('#s1981');
    const a1982Target = document.querySelector('#s1982');
    const a1983Target = document.querySelector('#s1983');
    const a1984Target = document.querySelector('#s1984');
    const a1985Target = document.querySelector('#s1985');
    const a1986Target = document.querySelector('#s1986');
    const a1987Target = document.querySelector('#s1987');
    const a1988Target = document.querySelector('#s1988');
    const a1989Target = document.querySelector('#s1989');
    const a1990Target = document.querySelector('#s1990');
    const a1991Target = document.querySelector('#s1991');
    const a1992Target = document.querySelector('#s1992');
    const a1993Target = document.querySelector('#s1993');
    const a1994Target = document.querySelector('#s1994');
    const a1995Target = document.querySelector('#s1995');
    const a1996Target = document.querySelector('#s1996');
    const a1997Target = document.querySelector('#s1997');
    const a1998Target = document.querySelector('#s1998');
    const a1999Target = document.querySelector('#s1999');
    const a2000Target = document.querySelector('#s2000');
    const a2001Target = document.querySelector('#s2001');

    const textListUL = document.querySelector('#textListUL');

    const template = document.querySelector('#mycard');

    fetch('name.json')
    .then(response => response.json())
    .then(data => {
        for (const product of data) {
            let clone = template.content.cloneNode(true);
            let gamelink = clone.querySelector("a");
            gamelink.setAttribute("href", "/mister/arcade/games/" + product.setname + ".html")
            let span = clone.querySelector("h2");
            span.textContent = product.name.replace(/ *\([^)]*\) */g, "") + " (" + product.year + ")"
            let image = clone.querySelector("img");
            image.setAttribute("src", "/mister/arcade/snaps/" + product.setname + ".png");
            let nameString = product.name.toString();
            

            // Populate Names
            switch (nameString.charAt(0).toUpperCase()) {
                case "A":
                    aTarget.appendChild(clone);    
                    break;
                case "B":
                    bTarget.appendChild(clone);    
                    break;
                case "C":
                    cTarget.appendChild(clone);    
                    break;
                case "D":
                    dTarget.appendChild(clone);    
                    break;
                case "E":
                    eTarget.appendChild(clone);    
                    break;
                case "F":
                    fTarget.appendChild(clone);    
                    break;
                case "G":
                    gTarget.appendChild(clone);    
                    break;
                case "H":
                    hTarget.appendChild(clone);    
                    break;
                case "I":
                    iTarget.appendChild(clone);    
                    break;
                case "J":
                    jTarget.appendChild(clone);    
                    break;
                case "K":
                    kTarget.appendChild(clone);    
                    break;
                case "L":
                    lTarget.appendChild(clone);    
                    break;
                case "M":
                    mTarget.appendChild(clone);    
                    break;
                case "N":
                    nTarget.appendChild(clone);    
                    break;
                case "O":
                    oTarget.appendChild(clone);    
                    break;
                case "P":
                    pTarget.appendChild(clone);    
                    break;
                case "Q":
                    qTarget.appendChild(clone);    
                    break;
                case "R":
                    rTarget.appendChild(clone);    
                    break;
                case "S":
                    sTarget.appendChild(clone);    
                    break;
                case "T":
                    tTarget.appendChild(clone);    
                    break;
                case "U":
                    uTarget.appendChild(clone);    
                    break;
                case "V":
                    vTarget.appendChild(clone);    
                    break;
                case "W":
                    wTarget.appendChild(clone);    
                    break;
                case "X":
                    xTarget.appendChild(clone);    
                    break;
                case "Y":
                    yTarget.appendChild(clone);    
                    break;
                case "Z":
                    zTarget.appendChild(clone);    
                    break;
                default:
                    myTarget.appendChild(clone);    
            }


            // Populate Years
            let clone2 = template.content.cloneNode(true);
            let gamelink2 = clone2.querySelector("a");
            gamelink2.setAttribute("href", "/mister/arcade/games/" + product.setname + ".html")
            let span2 = clone2.querySelector("h2");
            span2.textContent = product.name.replace(/ *\([^)]*\) */g, "") + " (" + product.year + ")"
            let image2 = clone2.querySelector("img");
            image2.setAttribute("src", "/mister/arcade/snaps/" + product.setname + ".png");
            let yearString = product.year.toString();
            switch (yearString) {
                case "1971":
                    a1971Target.appendChild(clone2);    
                    break;
                case "1972":
                    a1972Target.appendChild(clone2);    
                    break;
                case "1973":
                    a1973Target.appendChild(clone2);    
                    break;
                case "1974":
                    a1974Target.appendChild(clone2);    
                    break;
                case "1975":
                    a1975Target.appendChild(clone2);    
                    break;
                case "1976":
                    a1976Target.appendChild(clone2);    
                    break;
                case "1977":
                    a1977Target.appendChild(clone2);    
                    break;
                case "1978":
                    a1978Target.appendChild(clone2);    
                    break;
                case "1979":
                    a1979Target.appendChild(clone2);    
                    break;
                case "1980":
                    a1980Target.appendChild(clone2);    
                    break;
                case "1981":
                    a1981Target.appendChild(clone2);    
                    break;
                case "1982":
                    a1982Target.appendChild(clone2);    
                    break;
                case "1983":
                    a1983Target.appendChild(clone2);    
                    break;
                case "1984":
                    a1984Target.appendChild(clone2);    
                    break;
                case "1985":
                    a1985Target.appendChild(clone2);    
                    break;
                case "1986":
                    a1986Target.appendChild(clone2);    
                    break;
                case "1987":
                    a1987Target.appendChild(clone2);    
                    break;
                case "1988":
                    a1988Target.appendChild(clone2);    
                    break;
                case "1989":
                    a1989Target.appendChild(clone2);    
                    break;
                case "1990":
                    a1990Target.appendChild(clone2);    
                    break;
                case "1991":
                    a1991Target.appendChild(clone2);    
                    break;
                case "1992":
                    a1992Target.appendChild(clone2);    
                    break;
                case "1993":
                    a1993Target.appendChild(clone2);    
                    break;
                case "1994":
                    a1994Target.appendChild(clone2);    
                    break;
                case "1995":
                    a1995Target.appendChild(clone2);    
                    break;
                case "1996":
                    a1996Target.appendChild(clone2);    
                    break;
                case "1997":
                    a1997Target.appendChild(clone2);    
                    break;                    
                case "1998":
                    a1998Target.appendChild(clone2);    
                    break;
                case "1999":
                    a1999Target.appendChild(clone2);    
                    break;
                case "2000":
                    a2000Target.appendChild(clone2);    
                    break;                    
                case "2001":
                    a2001Target.appendChild(clone2);    
                    break;

            }

            // Populate Text List
            let li = document.createElement("li");
            let a = document.createElement("a");
            a.setAttribute("href", "/mister/arcade/games/" + product.setname + ".html")
            a.appendChild(document.createTextNode(product.name.replace(/ *\([^)]*\) */g, "") + " (" + product.year + ")"));
            li.appendChild(a);
            textListUL.appendChild(li);

            // Lazy load images
            var lazyloadImages = document.querySelectorAll("img.lazy");    
            var lazyloadThrottleTimeout;
            
            function lazyload () {
                if(lazyloadThrottleTimeout) {
                clearTimeout(lazyloadThrottleTimeout);
                }    
                
                lazyloadThrottleTimeout = setTimeout(function() {
                    var scrollTop = window.pageYOffset;
                    lazyloadImages.forEach(function(img) {
                        if(img.offsetTop < (window.innerHeight + scrollTop)) {
                        img.src = img.dataset.src;
                        img.classList.remove('lazy');
                        }
                    });
                    if(lazyloadImages.length == 0) { 
                    document.removeEventListener("scroll", lazyload);
                    window.removeEventListener("resize", lazyload);
                    window.removeEventListener("orientationChange", lazyload);
                    }
                }, 20);
            }
            
            document.addEventListener("scroll", lazyload);
            window.addEventListener("resize", lazyload);
            window.addEventListener("orientationChange", lazyload);
        }
    })
    .catch(console.error);
}
