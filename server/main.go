package main

import (
	"html/template"
	"net/http"
	"fmt"
	"asciiartweb/server/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ascii Art Web</title>
</head>
<body>
    <h1>Ascii Art Web</h1>
    <hr/>
    <form action="/ascii-art" method="post">
        <label for="user_input">Input Text: </label>
        <input type="text" name="user_input" id="user_input">
        <div>
            <h4>Change Font Color </h4>
        <select id="color-select" name="colors">
        <option value="aquamarine">Aquamarine</option>
        <option value="alice-blue">Alice Blue</option>
        <option value="alizarin-crimson">Alizarin Crimson</option>
        <option value="almond">Almond</option>
        <option value="amber">Amber</option>
        <option value="amaranth">Amaranth</option>
        <option value="amaranth-pink">Amaranth Pink</option>
        <option value="amazon">Amazon</option>
        <option value="amber-mist">Amber Mist</option>
        <option value="american-rose">American Rose</option>
        <option value="aqua">Aqua</option>
        <option value="azure">Azure</option>
        <option value="amethyst">Amethyst</option>
        <option value="apricot">Apricot</option>
        <option value="apple-green">Apple Green</option>
        <option value="asparagus">Asparagus</option>
        <option value="atomic-tangerine">Atomic Tangerine</option>
        <option value="aliceblue">AliceBlue</option>
        <option value="antiquewhite">AntiqueWhite</option>
        <option value="beige">Beige</option>
        <option value="bisque">Bisque</option>
        <option value="black">Black</option>
        <option value="blanched-almond">Blanched Almond</option>
        <option value="blue">Blue</option>
        <option value="blue-violet">Blue Violet</option>
        <option value="brick-red">Brick Red</option>
        <option value="brown">Brown</option>
        <option value="burlywood">Burlywood</option>
        <option value="byzantine">Byzantine</option>
        <option value="baby-blue">Baby Blue</option>
        <option value="baby-pink">Baby Pink</option>
        <option value="banana">Banana</option>
        <option value="barn-red">Barn Red</option>
        <option value="battleship-grey">Battleship Grey</option>
        <option value="beaver">Beaver</option>
        <option value="bistre">Bistre</option>
        <option value="bittersweet">Bittersweet</option>
        <option value="blond">Blond</option>
        <option value="blue-green">Blue Green</option>
        <option value="blue-grey">Blue Grey</option>
        <option value="blue-purple">Blue Purple</option>
        <option value="blush">Blush</option>
        <option value="bottle-green">Bottle Green</option>
        <option value="boysenberry">Boysenberry</option>
        <option value="brass">Brass</option>
        <option value="brick">Brick</option>
        <option value="bright-green">Bright Green</option>
        <option value="bright-pink">Bright Pink</option>
        <option value="bright-turquoise">Bright Turquoise</option>
        <option value="bright-ube">Bright Ube</option>
        <option value="brilliant-lavender">Brilliant Lavender</option>
        <option value="brilliant-rose">Brilliant Rose</option>
        <option value="brink-pink">Brink Pink</option>
        <option value="british-racing-green">British Racing Green</option>
        <option value="bronze">Bronze</option>
        <option value="brown-sugar">Brown Sugar</option>
        <option value="bubble-gum">Bubble Gum</option>
        <option value="bubbles">Bubbles</option>
        <option value="buff">Buff</option>
        <option value="burgundy">Burgundy</option>
        <option value="burple">Burple</option>
        <option value="byzantium">Byzantium</option>
        <option value="baby-blue-eyes">Baby Blue Eyes</option>
        <option value="blanchedalmond">BlanchedAlmond</option>
        <option value="blueviolet">BlueViolet</option>
        <option value="cadet-blue">Cadet Blue</option>
        <option value="cadmium-green">Cadmium Green</option>
        <option value="cadmium-orange">Cadmium Orange</option>
        <option value="cadmium-red">Cadmium Red</option>
        <option value="cadmium-yellow">Cadmium Yellow</option>
        <option value="caf">Café</option>
        <option value="cal-poly-green">Cal Poly Green</option>
        <option value="cambridge-blue">Cambridge Blue</option>
        <option value="camel">Camel</option>
        <option value="cameo-pink">Cameo Pink</option>
        <option value="camouflage-green">Camouflage Green</option>
        <option value="canary-yellow">Canary Yellow</option>
        <option value="candy-apple-red">Candy Apple Red</option>
        <option value="candy-pink">Candy Pink</option>
        <option value="capri">Capri</option>
        <option value="caput-mortuum">Caput Mortuum</option>
        <option value="caramel">Caramel</option>
        <option value="cardinal">Cardinal</option>
        <option value="caribbean-green">Caribbean Green</option>
        <option value="carmine">Carmine</option>
        <option value="carnation-pink">Carnation Pink</option>
        <option value="carnelian">Carnelian</option>
        <option value="carolina-blue">Carolina Blue</option>
        <option value="carrot-orange">Carrot Orange</option>
        <option value="celadon">Celadon</option>
        <option value="celeste">Celeste</option>
        <option value="celtic-blue">Celtic Blue</option>
        <option value="cerise">Cerise</option>
        <option value="cerulean">Cerulean</option>
        <option value="cerulean-blue">Cerulean Blue</option>
        <option value="cerulean-frost">Cerulean Frost</option>
        <option value="cg-blue">CG Blue</option>
        <option value="cg-red">CG Red</option>
        <option value="champagne">Champagne</option>
        <option value="charcoal">Charcoal</option>
        <option value="chartreuse">Chartreuse</option>
        <option value="cherry">Cherry</option>
        <option value="chestnut">Chestnut</option>
        <option value="chili-red">Chili Red</option>
        <option value="china-pink">China Pink</option>
        <option value="china-rose">China Rose</option>
        <option value="chinese-red">Chinese Red</option>
        <option value="chinese-violet">Chinese Violet</option>
        <option value="chocolate">Chocolate</option>
        <option value="chrome-yellow">Chrome Yellow</option>
        <option value="cinnamon">Cinnamon</option>
        <option value="cadetblue">CadetBlue</option>
        <option value="coral">Coral</option>
        <option value="cornflowerblue">CornflowerBlue</option>
        <option value="cornsilk">Cornsilk</option>
        <option value="crimson">Crimson</option>
        <option value="cyan">Cyan</option>
        <option value="daffodil">Daffodil</option>
        <option value="dandelion">Dandelion</option>
        <option value="dark-blue">Dark Blue</option>
        <option value="dark-brown">Dark Brown</option>
        <option value="dark-byzantium">Dark Byzantium</option>
        <option value="dark-candy-apple-red">Dark Candy Apple Red</option>
        <option value="dark-cerulean">Dark Cerulean</option>
        <option value="dark-chestnut">Dark Chestnut</option>
        <option value="dark-coral">Dark Coral</option>
        <option value="dark-cyan">Dark Cyan</option>
        <option value="dark-electric-blue">Dark Electric Blue</option>
        <option value="dark-goldenrod">Dark Goldenrod</option>
        <option value="dark-gray">Dark Gray</option>
        <option value="dark-green">Dark Green</option>
        <option value="dark-imperial-blue">Dark Imperial Blue</option>
        <option value="dark-jungle-green">Dark Jungle Green</option>
        <option value="dark-khaki">Dark Khaki</option>
        <option value="dark-lava">Dark Lava</option>
        <option value="dark-lavender">Dark Lavender</option>
        <option value="dark-liver">Dark Liver</option>
        <option value="dark-magenta">Dark Magenta</option>
        <option value="dark-midnight-blue">Dark Midnight Blue</option>
        <option value="dark-moss-green">Dark Moss Green</option>
        <option value="dark-olive-green">Dark Olive Green</option>
        <option value="dark-orange">Dark Orange</option>
        <option value="dark-orchid">Dark Orchid</option>
        <option value="dark-pastel-blue">Dark Pastel Blue</option>
        <option value="dark-pastel-green">Dark Pastel Green</option>
        <option value="dark-pastel-purple">Dark Pastel Purple</option>
        <option value="dark-pastel-red">Dark Pastel Red</option>
        <option value="dark-pink">Dark Pink</option>
        <option value="dark-powder-blue">Dark Powder Blue</option>
        <option value="dark-raspberry">Dark Raspberry</option>
        <option value="dark-red">Dark Red</option>
        <option value="dark-salmon">Dark Salmon</option>
        <option value="dark-scarlet">Dark Scarlet</option>
        <option value="dark-sea-green">Dark Sea Green</option>
        <option value="dark-sienna">Dark Sienna</option>
        <option value="dark-sky-blue">Dark Sky Blue</option>
        <option value="dark-slate-blue">Dark Slate Blue</option>
        <option value="dark-slate-gray">Dark Slate Gray</option>
        <option value="dark-spring-green">Dark Spring Green</option>
        <option value="dark-tan">Dark Tan</option>
        <option value="dark-tangerine">Dark Tangerine</option>
        <option value="dark-taupe">Dark Taupe</option>
        <option value="dark-terra-cotta">Dark Terra Cotta</option>
        <option value="dark-turquoise">Dark Turquoise</option>
        <option value="dark-violet">Dark Violet</option>
        <option value="dark-yellow">Dark Yellow</option>
        <option value="dartmouth-green">Dartmouth Green</option>
        <option value="davys-grey">Davys Grey</option>
        <option value="debian-red">Debian Red</option>
        <option value="deep-blue">Deep Blue</option>
        <option value="deep-carmine">Deep Carmine</option>
        <option value="deep-champagne">Deep Champagne</option>
        <option value="deep-chestnut">Deep Chestnut</option>
        <option value="deep-coffee">Deep Coffee</option>
        <option value="deep-fuchsia">Deep Fuchsia</option>
        <option value="deep-green">Deep Green</option>
        <option value="deep-magenta">Deep Magenta</option>
        <option value="deep-pink">Deep Pink</option>
        <option value="deep-purple">Deep Purple</option>
        <option value="deep-red">Deep Red</option>
        <option value="deep-saffron">Deep Saffron</option>
        <option value="deep-sky-blue">Deep Sky Blue</option>
        <option value="deep-space-sparkle">Deep Space Sparkle</option>
        <option value="deep-tuscan-red">Deep Tuscan Red</option>
        <option value="denim">Denim</option>
        <option value="denim-blue">Denim Blue</option>
        <option value="desert">Desert</option>
        <option value="desert-sand">Desert Sand</option>
        <option value="dim-gray">Dim Gray</option>
        <option value="dodger-blue">Dodger Blue</option>
        <option value="dogwood-rose">Dogwood Rose</option>
        <option value="drab">Drab</option>
        <option value="duke-blue">Duke Blue</option>
        <option value="dust-storm">Dust Storm</option>
        <option value="dutch-white">Dutch White</option>
        <option value="earth-yellow">Earth Yellow</option>
        <option value="ebony">Ebony</option>
        <option value="ecru">Ecru</option>
        <option value="eggplant">Eggplant</option>
        <option value="eggshell">Eggshell</option>
        <option value="egyptian-blue">Egyptian Blue</option>
        <option value="electric-blue">Electric Blue</option>
        <option value="electric-crimson">Electric Crimson</option>
        <option value="electric-cyan">Electric Cyan</option>
        <option value="electric-green">Electric Green</option>
        <option value="electric-indigo">Electric Indigo</option>
        <option value="electric-lavender">Electric Lavender</option>
        <option value="electric-lime">Electric Lime</option>
        <option value="electric-purple">Electric Purple</option>
        <option value="electric-ultramarine">Electric Ultramarine</option>
        <option value="electric-violet">Electric Violet</option>
        <option value="electric-yellow">Electric Yellow</option>
        <option value="emerald">Emerald</option>
        <option value="eminence">Eminence</option>
        <option value="english-green">English Green</option>
        <option value="english-lavender">English Lavender</option>
        <option value="english-red">English Red</option>
        <option value="english-violet">English Violet</option>
        <option value="eton-blue">Eton Blue</option>
        <option value="ebony">Ebony</option>
        <option value="eggshell">Eggshell</option>
        <option value="floralwhite">FloralWhite</option>
        <option value="forestgreen">ForestGreen</option>
        <option value="fuchsia">Fuchsia</option>
        <option value="gold">Gold</option>
        <option value="gainsboro">Gainsboro</option>
        <option value="ghostwhite">GhostWhite</option>
        <option value="goldenrod">GoldenRod</option>
        <option value="greenyellow">GreenYellow</option>
        <option value="honeydew">HoneyDew</option>
        <option value="hotpink">HotPink</option>
        <option value="indianred">IndianRed</option>
        <option value="indigo">Indigo</option>
        <option value="ivory">Ivory</option>
        <option value="khaki">Khaki</option>
        <option value="lavender">Lavender</option>
        <option value="lavenderblush">LavenderBlush</option>
        <option value="lawngreen">LawnGreen</option>
        <option value="lemonchiffon">LemonChiffon</option>
        <option value="lightblue">LightBlue</option>
        <option value="lightcoral">LightCoral</option>
        <option value="lightcyan">LightCyan</option>
        <option value="lightgoldenrodyellow">LightGoldenRodYellow</option>
        <option value="lightgray">LightGray</option>
        <option value="lightgreen">LightGreen</option>
        <option value="lightpink">LightPink</option>
        <option value="lightsalmon">LightSalmon</option>
        <option value="lightseagreen">LightSeaGreen</option>
        <option value="lightskyblue">LightSkyBlue</option>
        <option value="lightslategray">LightSlateGray</option>
        <option value="lightsteelblue">LightSteelBlue</option>
        <option value="lightyellow">LightYellow</option>
        <option value="lime">Lime</option>
        <option value="limegreen">LimeGreen</option>
        <option value="linen">Linen</option>
        <option value="magenta">Magenta</option>
        <option value="maroon">Maroon</option>
        <option value="mediumaquamarine">MediumAquaMarine</option>
        <option value="mediumblue">MediumBlue</option>
        <option value="mediumorchid">MediumOrchid</option>
        <option value="mediumpurple">MediumPurple</option>
        <option value="mediumseagreen">MediumSeaGreen</option>
        <option value="mediumslateblue">MediumSlateBlue</option>
        <option value="mediumspringgreen">MediumSpringGreen</option>
        <option value="mediumturquoise">MediumTurquoise</option>
        <option value="mediumvioletred">MediumVioletRed</option>
        <option value="midnightblue">MidnightBlue</option>
        <option value="mintcream">MintCream</option>
        <option value="mistyrose">MistyRose</option>
        <option value="moccasin">Moccasin</option>
        <option value="navajowhite">NavajoWhite</option>
        <option value="navy">Navy</option>
        <option value="oldlace">OldLace</option>
        <option value="olive">Olive</option>
        <option value="olivedrab">OliveDrab</option>
        <option value="orange">Orange</option>
        <option value="orangered">OrangeRed</option>
        <option value="orchid">Orchid</option>
        <option value="palegoldenrod">PaleGoldenRod</option>
        <option value="palegreen">PaleGreen</option>
        <option value="paleturquoise">PaleTurquoise</option>
        <option value="palevioletred">PaleVioletRed</option>
        <option value="papayawhip">PapayaWhip</option>
        <option value="peachpuff">PeachPuff</option>
        <option value="peru">Peru</option>
        <option value="pink">Pink</option>
        <option value="plum">Plum</option>
        <option value="powderblue">PowderBlue</option>
        <option value="purple">Purple</option>
        <option value="rebeccapurple">RebeccaPurple</option>
        <option value="red">Red</option>
        <option value="rosybrown">RosyBrown</option>
        <option value="royalblue">RoyalBlue</option>
        <option value="saddlebrown">SaddleBrown</option>
        <option value="salmon">Salmon</option>
        <option value="sandybrown">SandyBrown</option>
        <option value="seagreen">SeaGreen</option>
        <option value="seashell">SeaShell</option>
        <option value="sienna">Sienna</option>
        <option value="silver">Silver</option>
        <option value="skyblue">SkyBlue</option>
        <option value="slateblue">SlateBlue</option>
        <option value="slategray">SlateGray</option>
        <option value="snow">Snow</option>
        <option value="springgreen">SpringGreen</option>
        <option value="steelblue">SteelBlue</option>
        <option value="tan">Tan</option>
        <option value="teal">Teal</option>
        <option value="thistle">Thistle</option>
        <option value="tomato">Tomato</option>
        <option value="turquoise">Turquoise</option>
        <option value="violet">Violet</option>
        <option value="wheat">Wheat</option>
        <option value="white">White</option>
        <option value="whitesmoke">WhiteSmoke</option>
        <option value="yellow">Yellow</option>
        <option value="yellowgreen">YellowGreen</option>
        </select>
        </div>
        <h4>Choose A Banner File</h4>
        <fieldset>
            <legend>Banner Options</legend>
            <input type="radio" name="banner" id="standard" value="standard" checked>
            <label for="standard1">Standard</label><br>
            <input type="radio" name="banner" id="shadow" value="shadow">
            <label for="shadow">Shadow</label><br>
            <input type="radio" name="banner" id="thinkertoy" value="thinkertoy">
            <label for="thinkertoy">Thinkertoy</label><br>
        </fieldset>
        <input type="submit" value="Create Art">
    </form>
    <div id="response"></div>
</body>
</html>
`
			t, _ := template.New("index").Parse(tmpl)
			t.Execute(w, nil)
		} else {
			api.PostHandler(w, r)
		}
	})

	fmt.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
