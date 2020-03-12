package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type MenuType int

const (
    WeeklyMenu MenuType = 0
    DailyMenu  MenuType = 1
)

func constructUrl(menuType MenuType, restaurant RestaurantId, time time.Time) string {
    prefix := "https://sodexo.fi/ruokalistat/output/"
    switch menuType {
    case WeeklyMenu:
        return fmt.Sprintf("%sweekly_json/%s", prefix, restaurant)
    case DailyMenu:
        year, month, day := time.Date()
        date_str := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
        return fmt.Sprintf("%sdaily_json/%s/%s", prefix, restaurant, date_str)
    }
    return ""
}

func fetchJson(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}
