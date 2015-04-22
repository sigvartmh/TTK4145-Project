package main

import (
	"./driver"
	"./driver/src"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	ET_COMEDI cwrapper.ElevatorType = iota
	ET_SIMULATOR
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter workstation public IP: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	message := make(chan string)
	message2 := make(chan string)
	//msg <- "Test channel"
	go func() {
		var str string = "Starting Test"
		message <- str
	}()
	go func() {
		for {
			fmt.Println("Cgo calls: ", runtime.NumCgoCall())
			time.Sleep(500 * time.Millisecond)
		}
	}()
	//fmt.Println(<-message)
	select {
	case msg := <-message:
		fmt.Println("Test message: ", msg)
	}
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	go driver.Init(ET_SIMULATOR, message, message2)

	for {
		select {
		case msg := <-message:
			fmt.Println("Recived on channel:", msg)
			fmt.Println("Cgo calls: ", runtime.NumCgoCall())
			fmt.Println("Go rutines: ", runtime.NumGoroutine())
		case msg2 := <-message2:
			fmt.Println("Recived on channel2:", msg2)
		}
	}
}
