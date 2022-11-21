package main

import "time"

//Реализовать собственную функцию sleep

func TwentyFive() {
	mySleep(2 * time.Second)
}

func mySleep(d time.Duration) {
	<-time.After(d)
}
