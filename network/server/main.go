package main

import "./server"
import "fmt"

func main() {
	res := make(chan string)
	go server.Server(res)
	for {
		select {
		case str := <-res:
			fmt.Println("Recived on channel1:", str)
		}
	}
}
