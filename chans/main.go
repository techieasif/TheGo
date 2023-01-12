package main

import (
	"github.com/techieasif/TheGo/chans/chans_helper"
	"log"
)

func main() {
	intChan := make(chan int)

	go getValThroughChannel(intChan)
	log.Print(<-intChan)
}

func getValThroughChannel(intChan chan int) {
	intChan <- chans_helper.RandomNumberGenerator(100)
}
