package udp

import "net"
import "fmt"
import (
	"time"
)
import "encoding/json"
import "strconv"
import . "../../queue"

const bufSize int = 1024

//TODO: Add state channel which communicates if it's master or slave
func Server(port string, backup chan Que) {
	baddr, err := net.ResolveUDPAddr("udp4", port)
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
	go handleReception(lnb, backup)
}

func handleReception(conn *net.UDPConn, backup chan Que) {

	var item Que
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
	defer conn.Close()
	fmt.Printf("Received : %+v\n", item)
	fmt.Println("recived from:", rAddr)
	backup <- item
}

func Client(port int, queue chan Que) {
	bAddr := GetLocalIP()[:12] + "255" + strconv.Itoa(port)
	broadcast, err := net.ResolveUDPAddr("udp", bAddr)
	conn, err := net.DialUDP("udp", nil, broadcast)
	if err != nil {
		fmt.Println("Error dialing server")
	}

	for {
		select {
		case q := <-queue:
			buf, err := json.Marshal(&q)
			fmt.Println(buf)
			l, err := conn.Write(buf)
			if err != nil {
				fmt.Println("Error writing byte:", l, "to udp address:", broadcast)
			}
		}
	}
	defer conn.Close()
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
	return laddr.IP.String()

}

func Heartbeat(int port, master chan bool, slave chan bool) {
	select {
		b := make([]byte, bufSize)
		bAddr := GetLocalIP()[:12] + "255" + strconv.Itoa(port)
		broadcast, err := net.ResolveUDPAddr("udp", bAddr)
		conn, err := net.DialUDP("udp", nil, broadcast)
		if err != nil {
			fmt.Println("Failed to dial baddr", err)
			panic("Failed to dial baddr for laddr generation")
		}
	case <-master:
		for {
			select {
			case <-time.After(time.Millisecond * 150):
				buf := "Alive"
				conn.Write(buf)
			}
		}
		defer conn.Close()

	}

}
