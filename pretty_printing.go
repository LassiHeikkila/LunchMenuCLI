package main

import (
	"fmt"
	"sort"
)

func prettyPrintDailyList(l *DailyList, lan Language) {
	if l == nil {
		return
	}

	var intro_str string

	switch lan {
	case LanguageEN:
		intro_str = "Today's menu for"
	case LanguageFI:
		intro_str = "Ruokalista tänään ravintolassa"
	}


	fmt.Printf("%s: %s\n", intro_str, l.Meta.RefTitle)

	var keys []string
	for n := range l.Courses {
		keys = append(keys, n)
	}
	sort.Strings(keys)
	for _, key := range keys {
		course := l.Courses[key]
		courseTitle := course.TitleFi
		if lan == LanguageEN && course.TitleEn != "" {
			courseTitle = course.TitleEn
		}
		fmt.Printf("[%s]: %s\n", categoryToLang(course.Category, lan), courseTitle)
	}
}

func prettyPrintWeeklyList(l *WeeklyList, lan Language) {
	if l == nil {
		return
	}
	var intro_str string
	var option_str string

	switch lan {
	case LanguageEN:
		intro_str = "This week's menu for"
		option_str = "Options on"
	case LanguageFI:
		intro_str = "Tämän viikon ruokalista ravintolassa"
		option_str = "Vaihtoehdot"
	}

	fmt.Printf("%s: %s\n", intro_str, l.Meta.RefTitle)
	for _, m := range l.MealDates {
		fmt.Printf("%s %s:\n", option_str, dateToLang(m.Date, lan))
		var keys []string
		for n := range m.Courses {
			keys = append(keys, n)
		}
		sort.Strings(keys)
		for _, key := range keys {
			course := m.Courses[key]
			courseTitle := course.TitleFi
			if lan == LanguageEN && course.TitleEn != "" {
				courseTitle = course.TitleEn
			}
			fmt.Printf("\t[%s] %s: %s\n", categoryToLang(course.Category, lan), key, courseTitle)
		}
	}
}

func dateToLang(date string, lan Language) string {
	switch lan {
	case LanguageFI:
		return date
	case LanguageEN:
		dateMap := map[string]string{
			"Maanantai":   "Monday",
			"Tiistai":     "Tuesday",
			"Keskiviikko": "Wednesday",
			"Torstai":     "Thursday",
			"Perjantai":   "Friday",
			"Lauantai":    "Saturday",
			"Sunnuntai":   "Sunday",
		}
		if entry, ok := dateMap[date]; ok {
			return entry
		}
	}
	return date
}

func categoryToLang(cat string, lan Language) string {
	switch lan {
	case LanguageFI:
		return cat
	case LanguageEN:
		catMap := map[string]string{
			"Kotiruoka":     "Home-made dish",
			"Kotiruoka 2":   "Home-made dish 2",
			"Jälkiruoka":    "Dessert",
			"Kasviskeitto":  "Vegetable soup",
			"Salaattibaari": "Salad bar",
		}
		if entry, ok := catMap[cat]; ok {
			return entry
		}
	}
	return cat
}
