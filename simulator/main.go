package main

import (
	"./driver/src"
	"fmt"
)

func main() {
	message := make(chan string)
	message2 := make(chan string)
	//msg <- "Test channel"
	go func() {
		var str string = "Test"
		message <- str
	}()
	fmt.Println(<-message)
	driver.Init(message, message2)
	go func() {
		var str string = "Test"
		message <- str
	}()

	for {
		select {
		case msg := <-message:
			fmt.Println("Recived on channel:", msg)
		case msg2 := <-message2:
			fmt.Println("Recived on channel:", msg2)
		}
	}
}
