package main

import (
	"encoding/json"
)

type Language int8

const (
	LanguageFI Language = 0
	LanguageEN Language = 1
)

type WeeklyList struct {
	Meta       MetaInfo
	TimePeriod string `json:"timeperiod"`
	MealDates  []struct {
		Courses map[string]FoodEntry `json:"courses"`
		Date    string               `json:"date"`
	} `json:"mealdates"`
}

type DailyList struct {
	Meta    MetaInfo
	Courses map[string]FoodEntry `json:"courses"`
}

type FoodEntry struct {
	TitleFi    string `json:"title_fi"`
	TitleEn    string `json:"title_en"`
	Category   string `json:"category"`
	Properties string `json:"properties"`
	Price      string `json:"price"`
}

type MetaInfo struct {
	GeneratedTimestamp int    `json:"generated_timestamp"`
	RefURL             string `json:"ref_url"`
	RefTitle           string `json:"ref_title"`
}

func parseDailyJson(data []byte) (*DailyList, error) {
	var l DailyList

	err := json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func parseWeeklyJson(data []byte) (*WeeklyList, error) {
	var l WeeklyList

	err := json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
