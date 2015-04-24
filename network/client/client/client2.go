package client2

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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

	encoder := json.NewEncoder(conn)
	//item := QueItem{"localhost", 2, 0, false}
	item := QueItem{"PenisLars", 3, 0, true}
	//item := 6
	encoder.Encode(&item)
	//conn.write(&item)
	//b, err := json.Marshal(&item)
	//conn.Write(b)
	//conn.Close()
	//conn, err = net.Dial("tcp", "129.241.186.229:20013")
	item2 := QueItem{"localhost", 2, 0, false}
	//b2, err := json.Marshal(&item2)
	encoder.Encode(&item2)
	encoder.Encode(3)
	conn.Close()
	fmt.Println("done")

	fmt.Println("Start udp")
	ServerAddr, err := net.ResolveUDPAddr("udp", "129.241.186.255:20055")
	connUDP, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		fmt.Println("Error dialing server")
	}
	msg := QueItem{"UDP item", 3, 0, true}
	buf, err := json.Marshal(&msg)
	fmt.Println(buf)
	n, err := connUDP.Write(buf)
	if err != nil {
		fmt.Println("Wrote byte:", n, "to udp address")
	}
	connUDP.Close()
	fmt.Println("done")

}
