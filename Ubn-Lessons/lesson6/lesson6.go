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

  var movieStartSentence string = "Movie start time"
	var movieStartHour string = "12"
	var movieStartMinute string = "30"
	//var movieStartTime string = movieStartSentence + " " + "=" + " " + MovieStartHour + ":" + movieStartMinute
	var movieStartTime string = combineHourMinuteSentence(movieStartSentence, movieStartHour, movieStartMinute)
	fmt.Println(movieStartTime)
}
