//Golang Basic Lesson-6
package main

import "fmt"

func main() {

  var cinemaOpenSentence string = "Cinema opening time"
	var cinemaOpenHour string = "10"
	var cinemaOpenMinute string = "00"
	//var cinemaOpenTime string = cinemaOpenSentence + " " + "=" + " " + cinemaOpenHour + ":" + cinemaOpenMinute
	var cinemaOpenTime string = combineHourMinuteSentence(cinemaOpenSentence, cinemaOpenHour, cinemaOpenMinute)
	fmt.Println(cinemaOpenTime)
}
