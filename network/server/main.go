package main

import "./server/tcp"
import "./server/udp"
import "fmt"

func main() {
	t := make(chan string)
	u := make(chan string)
	go tcp.Server(t)
	go udp.Server(u)
	for {
		select {
		case str := <-t:
			fmt.Println("Recived on TCP:", str)
		case str := <-u:
			fmt.Println("Recived on UDP:", str)
		}
	}
}
