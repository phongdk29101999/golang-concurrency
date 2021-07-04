package main

import "log"

func dummy(dummyChan chan int) {
	dummyChan <- 10
}

func main() {
	chanInt := make(chan int)
	go dummy(chanInt)
	log.Println("chanInt data:", <-chanInt)
}