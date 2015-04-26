package tcp

import "net"
import "fmt"
import "log"
import "encoding/json"
import "sync"
import "strconv"

type QueItem struct {
	IP       string
	Floor    int
	Type     int //Up = 0, Down = 1, Command = 2
	Complete bool
}

type ConnectionList struct {
	IP []string
}
type connList struct {
	List map[string]*net.TCPConn
	Mu   sync.Mutex
}

var Connection connList

const bufSize int = 1024

func handleConnection(conn net.Conn, rec chan string, res chan int) {
	for {
		select {
		case floor := <-res:

		default:
			var item QueItem
			buf := make([]byte, bufSize)
			l, err := conn.Read(buf)
			if err != nil || l < 0 {
				fmt.Println("Error reading from conn: ", conn)
				fmt.Println("Error reading: ", err)
				Connection.Mu.Lock()
				conn.Close()
				delete(Connection.List, conn.RemoteAddr().String())
				defer Connection.Mu.Unlock()
				return
			} else {
				err = json.Unmarshal(buf[:l], &item)
				if err != nil {
					fmt.Println("Error converting to JSON", err)
				}

				fmt.Printf("Received : %+v\n", item)
				fmt.Printf("Connection map: %+v\n", Connection.List)
				fmt.Println("recived from:", conn.RemoteAddr())
				rec <- item.IP
			}
		}
	}
}

func Server(recived chan string, respond chan int) {
	Connection.List = make(map[string]*net.TCPConn)
	laddr := GetLocalIP(20013)
	ln, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		// handle error
		fmt.Println("Unable to listen to self on:20013", err)
		panic("Error listening tcp")
	}
	fmt.Println("Server listening to port:20013")
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Println("No accept", err)
			log.Println("Unable to accept connection", err)
		}
		raddr := conn.RemoteAddr()
		Connection.Mu.Lock()
		Connection.List[raddr.String()] = conn
		Connection.Mu.Unlock()
		go handleConnection(conn, recived, respond)
	}

}

func GetLocalIP(port int) *net.TCPAddr {
	baddr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(20323))

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
	laddr.Port = port

	tempConn.Close()
	return laddr

}
