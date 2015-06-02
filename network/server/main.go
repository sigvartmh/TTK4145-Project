package main

import "./server/tcp"
import "./server/udp"
import "fmt"
import (
	"bufio"
	"os"
	//"runtime"
	"time"
)

func main() {
	ti := make(chan tcp.QueItem)
	ui := make(chan udp.Que)
	to := make(chan tcp.QueItem)
	uo := make(chan udp.Que)
	go tcp.Server(20255, ti)
	go tcp.Client(tcp.GetLocalIP(20255).String(), to)
	go udp.Server(ui)

	go func(output chan tcp.QueItem, output2 chan udp.Que) {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter button pressed: ")
			text, _ := reader.ReadString('\n')
			if text == "q\n" {
				output <- tcp.QueItem{tcp.GetLocalIP(20255).IP.String(), 2, 3, true}
				//output2 <- 3
				fmt.Println("Output sent to channel")
			}
			fmt.Println("Text:", text, "text=='q'", text == "q\n")
		}
	}(to, uo)

	for {
		select {
		case str := <-ti:
			fmt.Println("Recived on TCP:", str)
		case str := <-ui:
			fmt.Println("Recived on UDP:", str)
		//case i := <-uo:
		//	fmt.Println("Sending data on UDP:", i)
		//case i := <-to:
		//	fmt.Println("Sending data on TCP:", i)
		default:
			//fmt.Println("Go rutines running: ", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}
}
