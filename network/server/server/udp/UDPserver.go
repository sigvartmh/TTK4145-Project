package udp

import "net"
import "fmt"
import "encoding/json"
import "strconv"

type QueItem struct {
	IP       string
	Floor    int
	Type     int //Up = 0, Down = 1, Command = 2
	Complete bool
}

type Que struct {
	Internal []QueItem
	External []QueItem
	Locale   []QueItem
}

const bufSize int = 1024

//TODO: Add state channel which communicates if it's master or slave
func Server(recived chan string) {
	baddr, err := net.ResolveUDPAddr("udp4", ":2055")
	if err != nil {
		//return err
		fmt.Println("Error resolving udpAddr")
	}

	lnb, err := net.ListenUDP("udp4", baddr)
	if err != nil {
		//return err
		fmt.Println("Error listening udpAddr")
		panic("Error listetning too broadcast address")
	}
	fmt.Println(lnb.LocalAddr())
	go handleReception(lnb, recived)
}

func handleReception(conn *net.UDPConn, res chan string) {
	defer conn.Close()
	var item QueItem
	buf := make([]byte, bufSize)
	l, rAddr, err := conn.ReadFromUDP(buf)

	if err != nil || l < 0 {
		fmt.Println("Error reading from UDP", err)
		return
	}

	jsonErr := json.Unmarshal(buf[:l], &item)
	if jsonErr != nil {
		fmt.Println("Error converting to JSON", err)
	}

	fmt.Printf("Received : %+v\n", item)
	//fmt.Println("recived from:", conn.RemoteAddr())
	fmt.Println("recived from:", rAddr)
	res <- item.IP
}

func GetLocalIP() string {
	baddr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(30039))

	if err != nil {
		fmt.Println("Could not resolve baddr", err)
		//return ""
		panic("Could not resolve baddr")
	}

	tempConn, err := net.DialUDP("udp4", nil, baddr)
	if err != nil {
		fmt.Println("Failed to dial baddr for laddr generation", err)
		//return ""
		panic("Failed to dial baddr for laddr generation")
	}
	tempAddr := tempConn.LocalAddr()
	laddr, err := net.ResolveTCPAddr("tcp4", tempAddr.String())
	if err != nil {
		fmt.Println("Failed to resolve laddr", err)
		//return ""
		panic("Failed to resolve laddr")

	}

	tempConn.Close()
	return string(laddr.IP)

}

/*
func Heartbeat(master chan bool) {
	for {
		select {
		case <-master == true:
			//sendUDPpacket

		}
	}
}*/
