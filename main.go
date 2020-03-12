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
)

func main() {
    flag.Parse()
    fmt.Println("Checking what's for lunch")

    menuType := DailyMenu
    if *weeklyFlag {
        menuType = WeeklyMenu
    }

    restaurant := defaultRestaurant
    if *restaurantFlag != "" {
        restaurant = RestaurantId(*restaurantFlag)
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

    switch menuType {
    case DailyMenu:
        dailylist, err := parseDailyJson(data)
        if err != nil {
            fmt.Println("Error while parsing daily menu:", err)
            return
        }
        prettyPrintDailyList(dailylist)
    case WeeklyMenu:
        weeklylist, err := parseWeeklyJson(data)
        if err != nil {
            fmt.Println("Error while parsing weekly menu:", err)
            return
        }
        prettyPrintWeeklyList(weeklylist)
    }

    return
}
