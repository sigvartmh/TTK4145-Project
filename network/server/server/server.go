package server

import "net"
import "fmt"
import "log"
import "encoding/json"

type QueItem struct {
	IP       string
	Floor    int
	Type     int //Up = 0, Down = 1, Command = 2
	Complete bool
}

/*func something() {
	jsonObject := json.Marshal(data)
	var res QueItem
	json.Unmarshal(data, &res)
}*/

func handleConnection(conn net.Conn, rec chan string) {
	dec := json.NewDecoder(conn)
	res := QueItem{}
	dec.Decode(&res)
	fmt.Printf("Received : %+v\n", res)
	fmt.Println("recived from:", conn.RemoteAddr())
	rec <- res.IP

	dec.Decode(&res)
	fmt.Printf("Received : %+v\n", res)
	fmt.Println("recived from:", conn.RemoteAddr())
	rec <- res.IP
}

func Server(recived chan string) {
	ln, err := net.Listen("tcp", ":20013")
	if err != nil {
		// handle error
		fmt.Println("Unable to listen to self on:20013")
	}
	fmt.Println("Server listening to port:20013")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("No accept", err)
			log.Println("Unable to accept connection", err)
		}
		go handleConnection(conn, recived)
	}

}
