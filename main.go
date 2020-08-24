package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	defaultRestaurant         = SodexoGalaksi
	weeklyFlag        *bool   = flag.Bool("weekly", false, "Get weekly menu instead of daily")
	restaurantFlag    *string = flag.String("restaurant", "", "Use this restaurant id instead of default 121 (Sodexo Galaksi)")
	raw               *bool   = flag.Bool("raw", false, "Print raw menu JSON to assist debugging")
	inEnglish         *bool   = flag.Bool("english", false, "Print menu in English, if possible")
)

func main() {
	flag.Parse()

	menuType := DailyMenu
	if *weeklyFlag {
		menuType = WeeklyMenu
	}

	restaurant := defaultRestaurant
	if *restaurantFlag != "" {
		restaurant = RestaurantId(*restaurantFlag)
	}

	lang := LanguageFI
	if *inEnglish {
		lang = LanguageEN
	}

	switch lang {
	case LanguageFI:
		fmt.Println("Haetaan ruokalistaa...")
	case LanguageEN:
		fmt.Println("Checking what's for lunch")
	}

	time := time.Now()

	url := constructUrl(menuType, restaurant, time)
	if url == "" {
		fmt.Println("Failed to construct a valid URL!")
		return
	}

	data, err := fetchJson(url)
	if err != nil {
		fmt.Println("Error while fetching menu:", err)
	}

	if *raw {
		fmt.Print("Raw JSON:\n", string(data), "\n\n")
	}

	switch menuType {
	case DailyMenu:
		dailylist, err := parseDailyJson(data)
		if err != nil {
			fmt.Println("Error while parsing daily menu:", err)
			return
		}
		prettyPrintDailyList(dailylist, lang)
	case WeeklyMenu:
		weeklylist, err := parseWeeklyJson(data)
		if err != nil {
			fmt.Println("Error while parsing weekly menu:", err)
			return
		}
		prettyPrintWeeklyList(weeklylist, lang)
	}

	return
}
