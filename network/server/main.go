package main

import "./server/tcp"
import "./server/udp"
import "fmt"
import (
	"bufio"
	"runtime"
	"time"
)

func main() {
	ti := make(chan string)
	ui := make(chan string)
	to := make(chan int)
	uo := make(chan int)
	go tcp.Server(t)
	go udp.Server(u)

	go func(output, output2 chan int) {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter button pressed: ")
			text, _ := reader.ReadString('\n')
			if text == "q" {
				output <- 2
				output2 <- 3
			}
		}
	}(to, uo)

	for {
		select {
		case str := <-ti:
			fmt.Println("Recived on TCP:", str)
		case str := <-ui:
			fmt.Println("Recived on UDP:", str)
		case i := <-uo:
			fmt.Println("Sending data on UDP:", i)
		case i := <-io:
			fmt.Println("Sending data on TCP:", i)
		default:
			fmt.Println("Go rutines running: ", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}
}
