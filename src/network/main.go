package main

import "./tcp"
import "./udp"
import . "../queue"
import "fmt"
import (
	"bufio"
	"os"
	//"runtime"
	"time"
)

var BrodcastPort, _ = strconv.Atoi(os.Getenv("BROADCASTPORT"))
var HearthbeatPort, _ = strconv.Atoi(os.Getenv("HEARTBEATPORT"))
var TCPPort, _ = strconv.Atoi(os.Getenv("TCPPORT"))

func main() {
	ti := make(chan QueItem)
	ui := make(chan Que)
	to := make(chan QueItem)
	uo := make(chan Que)

	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))

	go tcp.Server(20255, ti)
	go tcp.Client(tcp.GetLocalIP(20255).String(), to)
	go udp.Server(ui)

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
