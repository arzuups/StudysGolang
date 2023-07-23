// Golang Basic Lesson-6
package main

import (
    "fmt"
)

func main() {

    //FUNCTIONS

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

    var goSentence string = "My time to go"
    var goHour string = "12"
    var goMinute string = "00"
    //var goTime string = goSentence + " " + "=" + " " + goHour + ":" + goMinute
    var goTime string = combineHourMinuteSentence(goSentence, goHour, goMinute)
    fmt.Println(goTime)

}

func combineHourMinuteSentence(sentence string, hour string, minute string) string {
    time := sentence + " " + "=" + " " + hour + ":" + minute
    //fmt.Println(time)
    return time
}
