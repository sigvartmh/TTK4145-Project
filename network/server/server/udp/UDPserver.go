package udp

import "net"
import "fmt"
import "encoding/json"

type QueItem struct {
	IP       string
	Floor    int
	Type     int //Up = 0, Down = 1, Command = 2
	Complete bool
}

type Que struct {
	Internal []QueItem
	External []QueItem
}

const bufSize int = 1024

func Server(recived chan string) {
	baddr, err := net.ResolveUDPAddr("udp", ":20055")
	if err != nil {
		//return err
		fmt.Println("Error resolving udpAddr")
	}

	lnb, err := net.ListenUDP("udp", baddr)
	if err != nil {
		//return err
		fmt.Println("Error listening udpAddr")
		panic("Error listetning too broadcast address")
	}
	fmt.Println(lnb.LocalAddr())
	go handleReception(lnb, recived)
}

func handleReception(conn *net.UDPConn, res chan string) {
	buf := make([]byte, bufSize)
	var item QueItem
	for {
		//for {
		l, rAddr, err := conn.ReadFromUDP(buf)
		if err != nil || l < 0 {
			fmt.Println("Error reading from UDP")
		}
		jerr := json.Unmarshal(buf[:l], &item)
		fmt.Printf("Received : %+v\n", item)
		fmt.Println("recived from:", conn.RemoteAddr())
		fmt.Println("recived from:", rAddr)
		fmt.Println("Json err:", jerr)
		res <- item.IP
	}

	//}
}
