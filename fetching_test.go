package main

import (
    "testing"
    "time"
)

func TestURLConstruction(t *testing.T) {
    loc, err := time.LoadLocation("UTC")
    if err != nil {
        t.Error("Failed to load UTC location")
    }
    daily_url := constructUrl(DailyMenu, SodexoGalaksi, time.Date(2020, 3, 12, 12, 0, 0, 0, loc))
    if daily_url != "https://sodexo.fi/ruokalistat/output/daily_json/121/2020-03-12" {
        t.Error("Wrong daily url produced:", daily_url)
    }

    weekly_url := constructUrl(WeeklyMenu, SodexoGalaksi, time.Date(2020, 3, 12, 12, 0, 0, 0, loc))
    if weekly_url != "https://sodexo.fi/ruokalistat/output/weekly_json/121" {
        t.Error("Wrong weekly url produced:", weekly_url)
    }
}
