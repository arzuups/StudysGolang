// Golang Basic Lesson 6-2
package main

import "fmt"

func main() {

	//FUNCTIONS

	day := "25"
	month := "08"
	year := "2008"
	hour := "06"
	minute := "30"

	conclusion := combineDate(day, month, year, hour, minute)
	fmt.Println(conclusion)
}

func combineDate(day string, month string, year string, hour string, minute string) string {

	date := combineDayMonthYear(day, month, year)
	fmt.Println(date)
	time := combineHourMinute(hour, minute)
	fmt.Println(time)
	dateTime := dateTime(date, time)
	return dateTime
}

func combineHourMinute(hour string, minute string) string {
	time := hour + ":" + minute
	return time
}

func combineDayMonthYear(day string, month string, year string) string {
	date := day + "/" + month + "/" + year
	return date
}

func dateTime(date string, time string) string {
	dateTime := date + " " + time
	return dateTime
}

