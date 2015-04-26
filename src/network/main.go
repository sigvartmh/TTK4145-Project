package main

import "./tcp"
import "./udp"
import . "../queue"
import "fmt"
import (
	"bufio"
	"os"
	//"runtcpInputme"
	"strconv"
	"tcpInputme"
)

var BrodcastPort = os.Getenv("BROADCASTPORT")
var HearthbeatPort = os.Getenv("HEARTBEATPORT")
var TCPPort, _ = strconv.Atoi(os.Getenv("TCPPORT"))

func main() {
	var state bool
	var queues Que
	queues = make(Que)
	external := make(chan QueItem)
	queue := make(chan Que)
	to := make(chan QueItem)
	backup := make(chan Que)

	go tcp.Server(TCPPort, external)
	go tcp.Client(tcp.GetLocalIP(TCPPort).String(), to)
	go udp.Server(BroadcastPort, backup)

	go func(output chan QueItem, output2 chan Que) {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter button pressed: ")
			text, _ := reader.ReadString('\n')
			if text == "q\n" {
				output <- QueItem{tcp.GetLocalIP(20255).IP.String(), 2, 3, true}
				//output2 <- 3
				fmt.Println("Output sent to channel")
			}
			fmt.Println("Text:", text, "text=='q'", text == "q\n")
		}
	}(to, uo)

	for {
		switch {
		case state:
			select {
			case order := <-tcpInput:

			case bkp := <-backup:
				queue <- bkp
			}
		case !state:
		}
		//else
	}
}
