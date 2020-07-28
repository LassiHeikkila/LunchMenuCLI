package main

import (
	"fmt"
	"sort"
)

func prettyPrintDailyList(l *DailyList) {
	if l == nil {
		return
	}
	fmt.Printf("Daily Menu for %s\n", l.Meta.RefTitle)

	var keys []string
	for n := range l.Courses {
		keys = append(keys, n)
	}
	sort.Strings(keys)
	for _, key := range keys {
		course := l.Courses[key]
		fmt.Printf("[%s] option %s: %s\n", course.Category, key, course.TitleFi)
	}
}

func prettyPrintWeeklyList(l *WeeklyList) {
	if l == nil {
		return
	}
	fmt.Printf("Weekly Menu for: %s\n", l.Meta.RefTitle)
	for _, m := range l.MealDates {
		fmt.Printf("Options on %s:\n", m.Date)
		var keys []string
		for n := range m.Courses {
			keys = append(keys, n)
		}
		sort.Strings(keys)
		for _, key := range keys {
			course := m.Courses[key]
			fmt.Printf("\t[%s] %s: %s\n", course.Category, key, course.TitleFi)
		}
	}
}
