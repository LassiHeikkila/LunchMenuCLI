package main

import (
	"testing"
)

var (
	daily_json_example string = `{
    "meta":{
        "generated_timestamp":1584003734,
        "ref_url":"https:\/\/www.sodexo.fi\/ravintolat\/ravintola-galaksi",
        "ref_title":"Ravintola Galaksi"
    },
    "courses":{
        "1":{
            "title_fi":"Lihapullat L,G, Kermakastike L,G, Perunasose L,G",
            "title_en":"Meat balls, cream sauce, mashed potatoes",
            "category":"Kotiruoka",
            "properties":"G, L",
            "price":"9,90 \u20ac \/ 10,50 \u20ac"
        },
        "2":{
            "title_fi":"Chilpotle chicken pita, smetana (tarjolla vaim Galaksissa)",
            "title_en":"Chilpotle chicken pita, sour cream",
            "category":"Kotiruoka",
            "properties":"L",
            "price":"9,90 \u20ac \/ 10,50 \u20ac"
        },
        "3":{
            "title_fi":"Talon kalamureke L,G, tillikastike L,G",
            "title_en":"Home made fish loaf, dill sauce",
            "category":"Kotiruoka 2",
            "properties":"G, L",
            "price":"9,90 \u20ac \/ 10,50 \u20ac"
        },
        "4":{
            "title_fi":"Papu-p\u00e4hkin\u00e4tagine L,M, riisi (tarjolla vain Galaksissa)",
            "title_en":"Beans-nuts tagine, rice",
            "category":"Kasvisruoka",
            "properties":"G, L, M",
            "price":"9,90 \u20ac \/ 10,50 \u20ac"
        },
        "5":{
            "title_fi":"lohta L,M,G,soija, Chili-kalkkunaa L,M,G, Yrtti-marinoitus nyht\u00f6kauraa L,M ",
            "title_en":"",
            "category":"Salaattibaari",
            "properties":"M",
            "price":"9,90 \u20ac \/ 10,50 \u20ac"
        },
        "6":{
            "title_fi":"Paahdettua paprikakeittoa",
            "title_en":"Roasted peppers soup",
            "category":"Kasviskeitto",
            "properties":"G, L","price":""
        },
        "7":{
            "title_fi":"Mansikka-appelsiinikiisseli\u00e4 L,M,G, kermavaahtoa L,G",
            "title_en":"Strawberry orange kissel, whipped cream",
            "category":"J\u00e4lkiruoka",
            "properties":"G, L",
            "price":""
        }
    }
}`
	weekly_json_example string = `{
    "mealdates": [
        {
            "courses": {
                "1": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L",
                    "title_en": "Pasta Campagnole and pecorino cheese ",
                    "title_fi": "Pasta Campagnole ja Pecorinojuustoa"
                },
                "2": {
                    "category": "Kotiruoka 2",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L",
                    "title_en": "Chicken schnitzel, tandoor-buttersauce, rice",
                    "title_fi": "Paneroitu broilerpihvi L,M, tandoori-voikastiketta L,G,, riisiä L,M,G"
                },
                "3": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "G, L",
                    "title_en": "Oven baked potato with smoked reindeer root filling",
                    "title_fi": "Uuniperuna savuporo-juurestäytteellä (tarjolla vain Galaksissa)"
                },
                "4": {
                    "category": "Kasvisruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "Falafel balls, lens tartar, rice",
                    "title_fi": "Falafel pyörykät L,M , linssi-tartar L,M,G riisiä L,M,G(tarjolla vain Galakissa)"
                },
                "5": {
                    "category": "Salaattibaari",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "Västerbottens cheese, tunafish, garlic-turkey",
                    "title_fi": "Västerbottens juustoa L,G, tonnikalaa L,M,G, valkosipuli marinoitua kalkkunaa L,M,G"
                },
                "6": {
                    "category": "Kasviskeitto",
                    "price": "",
                    "properties": "G, M",
                    "title_en": "Tomato chili coriander soup",
                    "title_fi": "Tomaatti-chili-korianterikeitto"
                },
                "7": {
                    "category": "Jälkiruoka",
                    "price": "",
                    "properties": "G, VL",
                    "title_en": "Passion coconut mousse",
                    "title_fi": "Passion-kookosmoussea VL,G, pyydä L, pyydä M, saattaa sisältää pieniä määriä;gluteeni, kananmuna"
                }
            },
            "date": "Maanantai"
        },
        {
            "courses": {
                "1": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L",
                    "title_en": "Home made minced meat patties a'la Lidström, cream sauce, boiled potatoes",
                    "title_fi": "Itsetehdyt Linströmin pihvit L,G, kermakastiketta L,G, perunoita"
                },
                "2": {
                    "category": "Keitto",
                    "price": "",
                    "properties": "G, L",
                    "title_en": "Salmon soup, flatbread",
                    "title_fi": "Lohikeitto L,G, rieska L,M (tarjolla vain Galaksissa)"
                },
                "3": {
                    "category": "Kotiruoka 2",
                    "price": "9,90 € / 10,50 €",
                    "properties": "G, L",
                    "title_en": "Beef in sour cream sauce , rice",
                    "title_fi": "Härkää smetanakastikkeessa, riisiä"
                },
                "4": {
                    "category": "Kasvisruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L, M",
                    "title_en": "Paella with smoked tofu",
                    "title_fi": "Paella savustetusta tofusta (tarjolla vain Galaksissa)"
                },
                "5": {
                    "category": "Salaattibaari",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "Goat cheese, lemon-shrimp, smoked reindeer tartar",
                    "title_fi": "Vuohenjuustoa VL,G, sitruunamarinoituja katkarapuja L,M,G, Savuporotartar L,G, sis.kananmuna"
                },
                "6": {
                    "category": "Kasviskeitto",
                    "price": "",
                    "properties": "G, L",
                    "title_en": "Black salsify puree soup",
                    "title_fi": "Mustajuurisosekeittoa"
                },
                "7": {
                    "category": "Jälkiruoka",
                    "price": "",
                    "properties": "G, L",
                    "title_en": "Strawberry kissel, whipped cream",
                    "title_fi": "Mansikkakiisseliä L,M,G, saattaa sisältää;gluteeni, maito, kananmuna,  Kermavaahtoa L,G"
                }
            },
            "date": "Tiistai"
        },
        {
            "courses": {
                "1": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "Frankfurt stroganoff, mashed potatoes",
                    "title_fi": "Nakkistroganoffia L,G, perunasose L,G, (tarjolla vain Galaksissa)"
                },
                "2": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L",
                    "title_en": "Tunafish lasagnette",
                    "title_fi": "Tonnikalalasagnette"
                },
                "3": {
                    "category": "Kotiruoka 2",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "Breaded chicken fillet with sesame seeds, mojosalsa, rice",
                    "title_fi": "Seesam kuorrutettua broilerinfileetä M, L, pyydä G, mojosalsaa L,M,G, riisiä L,M,G"
                },
                "4": {
                    "category": "Kasvisruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L",
                    "title_en": "Mushroom goat cheese pie, red onion compote",
                    "title_fi": "Sieni-vuohenjuustopiirakkkaa L, punasispulihilloketta L,M,G (tarjolla vain Galaksissa)"
                },
                "5": {
                    "category": "Salaattibaari",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "lemon-garlic feta cheese, smoked white fish, tofu with teriyaki",
                    "title_fi": "sitruuna-valkosipulimarinoitua fetajuustoa G, savusiikaa L,M,G, teriayaki marinoitua tofua L,M,"
                },
                "6": {
                    "category": "Kasviskeitto",
                    "price": "",
                    "properties": "G, L",
                    "title_en": "Carrot soup",
                    "title_fi": "Porkkanakeittoa L,G"
                },
                "7": {
                    "category": "Jälkiruoka",
                    "price": "",
                    "properties": "G",
                    "title_en": "Lingonberry -caramel quark",
                    "title_fi": "Puolukka-kinuskirahkaa L,G, pyydä M"
                }
            },
            "date": "Keskiviikko"
        },
        {
            "courses": {
                "1": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "G, L",
                    "title_en": "Meat balls, cream sauce, mashed potatoes",
                    "title_fi": "Lihapullat L,G, Kermakastike L,G, Perunasose L,G"
                },
                "2": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "L",
                    "title_en": "Chilpotle chicken pita, sour cream",
                    "title_fi": "Chilpotle chicken pita, smetana (tarjolla vaim Galaksissa)"
                },
                "3": {
                    "category": "Kotiruoka 2",
                    "price": "9,90 € / 10,50 €",
                    "properties": "G, L",
                    "title_en": "Home made fish loaf, dill sauce",
                    "title_fi": "Talon kalamureke L,G, tillikastike L,G"
                },
                "4": {
                    "category": "Kasvisruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "G, L, M",
                    "title_en": "Beans-nuts tagine, rice",
                    "title_fi": "Papu-pähkinätagine L,M, riisi (tarjolla vain Galaksissa)"
                },
                "5": {
                    "category": "Salaattibaari",
                    "price": "9,90 € / 10,50 €",
                    "properties": "M",
                    "title_en": "",
                    "title_fi": "lohta L,M,G,soija, Chili-kalkkunaa L,M,G, Yrtti-marinoitus nyhtökauraa L,M "
                },
                "6": {
                    "category": "Kasviskeitto",
                    "price": "",
                    "properties": "G, L",
                    "title_en": "Roasted peppers soup",
                    "title_fi": "Paahdettua paprikakeittoa"
                },
                "7": {
                    "category": "Jälkiruoka",
                    "price": "",
                    "properties": "G, L",
                    "title_en": "Strawberry orange kissel, whipped cream",
                    "title_fi": "Mansikka-appelsiinikiisseliä L,M,G, kermavaahtoa L,G"
                }
            },
            "date": "Torstai"
        },
        {
            "courses": {
                "1": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "Swiss schnitzel, cream potatoes, lemon sour cream sauce",
                    "title_fi": "Sveitsinleike, kermaperunat, sitruuna-kermaviilikastike"
                },
                "2": {
                    "category": "Kotiruoka 2",
                    "price": "9,90 € / 10,50 €",
                    "properties": "G, L",
                    "title_en": "Minced meat tortilla, salsa",
                    "title_fi": "Jauhelihatortilla, salsa (tarjolla vain Galaksissa)"
                },
                "3": {
                    "category": "Kotiruoka",
                    "price": "9,90 € / 10,50 €",
                    "properties": "M",
                    "title_en": "Chicken wings, chimichurri dip and jasmine rice",
                    "title_fi": "Chicken wings, chimichurri dippiä ja jasmiiniriisiä"
                },
                "4": {
                    "category": "Kasvisruoka",
                    "price": "",
                    "properties": "M",
                    "title_en": "Broccoli bean pasta",
                    "title_fi": "Parsakaali-papupasta"
                },
                "5": {
                    "category": "Salaattibaari",
                    "price": "9,90 € / 10,50 €",
                    "properties": "",
                    "title_en": "",
                    "title_fi": "Salad bar:"
                },
                "6": {
                    "category": "Kasviskeitto",
                    "price": "",
                    "properties": "L",
                    "title_en": "Vegetable soup with cheese",
                    "title_fi": "Juustoinen kasviskeitto (tarjolla vain Galaksissa)"
                },
                "7": {
                    "category": "Jälkiruoka",
                    "price": "",
                    "properties": "L",
                    "title_en": "Bun, raspberry jam and whipped cream",
                    "title_fi": "Pappilan hätävara"
                }
            },
            "date": "Perjantai"
        }
    ],
    "meta": {
        "generated_timestamp": 1584004524,
        "ref_title": "Ravintola Galaksi",
        "ref_url": "https://www.sodexo.fi/ravintolat/ravintola-galaksi"
    },
    "timeperiod": "9.3. - 15.3."
}`
)

func TestParsingDailyJson(t *testing.T) {
	daily_list, err := parseDailyJson([]byte(daily_json_example))

	if err != nil {
		t.Error("Error returned", err)
	}
	if daily_list == nil {
		t.Fatal("Nil result")
	}

	if daily_list.Meta.RefTitle != "Ravintola Galaksi" {
		t.Error("Wrong ref title:", daily_list.Meta.RefTitle)
	}

	if daily_list.Courses["1"].TitleFi != "Lihapullat L,G, Kermakastike L,G, Perunasose L,G" {
		t.Error("Wrong entry for course 1:", daily_list.Courses["1"].TitleFi)
	}
}

func TestParsingWeeklyJson(t *testing.T) {
	weekly_list, err := parseWeeklyJson([]byte(weekly_json_example))

	if err != nil {
		t.Error("Error returned", err)
	}
	if weekly_list == nil {
		t.Fatal("Nil result")
	}
}
