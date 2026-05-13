package main

import ("fmt")

func main() {
 var city string // explict type declaration

 city = "london"

 var channel = "web dev"  // auto inference

 subcribers := 5000  //short way of declaring a variable wiith implict type defnition

 subcribers = subcribers+100

 likes, comments := 100, 300


 fmt.Println(city, channel, subcribers, likes, comments)

}