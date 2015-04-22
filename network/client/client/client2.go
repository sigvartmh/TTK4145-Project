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
	conn, err := net.Dial("tcp", "129.241.186.229:20013")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	//b := make([]byte, 1024)

	encoder := json.NewEncoder(conn)
	//item := QueItem{"localhost", 2, 0, false}
	item := QueItem{"PenisLars", 3, 0, true}
	encoder.Encode(&item)
	//conn.write(&item)
	//b, err := json.Marshal(&item)
	//conn.Write(b)
	//conn.Close()
	//conn, err = net.Dial("tcp", "129.241.186.229:20013")
	item2 := QueItem{"localhost", 2, 0, false}
	//b2, err := json.Marshal(&item2)
	encoder.Encode(&item2)
	conn.Close()
	fmt.Println("done")
}
