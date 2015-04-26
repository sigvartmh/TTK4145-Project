package client2

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
    "time"
)

type QueItem struct {
	IP       string
	Floor    int
	Type     int //Up = 0, Down = 1, Command = 2
	Complete bool
}

func Start() {
	fmt.Println("start client")
	conn, err := net.Dial("tcp", "129.241.186.210:20013")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	//b := make([]byte, 1024)
	//buf := make([]byte, 1024)

	//item := QueItem{"localhost", 2, 0, false}
	item := QueItem{"PenisLars", 3, 0, true}
	//item := 6
	b, err := json.Marshal(&item)
	fmt.Println("Json buff:", b)
	conn.Write(b)
	//conn.Write(b)
    for{
        conn.Write(b)
        time.Sleep(3* time.Second)
    }
	conn.Close()

	fmt.Println("Start udp")
	ServerAddr, err := net.ResolveUDPAddr("udp", "129.241.186.255:2055")
	connUDP, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		fmt.Println("Error dialing server")
	}
	msg := QueItem{"UDP item", 3, 0, true}
	buf, err := json.Marshal(&msg)
	fmt.Println(buf)
	n, err2 := connUDP.Write(buf)
	if err2 != nil {
		fmt.Println("Wrote byte:", n, "to udp address")
	}
    
	connUDP.Close()
	fmt.Println("done")

}

func emulateButtonPress() {

}
