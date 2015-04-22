package tcp

import "net"
import "fmt"
import "time"

type QueItem struct {
	IP       string
	Floor    int
	Dir      int //Up = 0, Down = 1, Command = 2
	Complete bool
}

type Que struct {
	Jobs     []QueItem
	External []QueItem
	Internal []QueItem
}

func TCPserver(externalOrders chan string) {

	tcpLocalAddr, err := net.ResolveTCPAddr("tcp", "129.241.187.104:33546")
	tcpListener, err := net.ListenTCP("tcp", tcpLocalAddr)
	if err != nil {
		fmt.Println("Error opening connection:", err.Error())
	}
	defer tcpListener.Close()
	for {
		// Listen for an incoming connection.
		conn, err := tcpListener.AcceptTCP()
		fmt.Println("accepted connection")
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleConnection(conn, externalOrders)
	} //Switch to for select loop
}

func handleConnection(conn *net.TCPConn, external chan string) {
	data := make([]byte, 1024)
	message := make([]byte, 1024)
	data = []byte("Cookies for parties\x00")
	for {
		_, err := conn.Write(data)
		if err != nil {
			fmt.Printf("Error in TCP: %s\n", err.Error())
			break
		}
		time.Sleep(100 * time.Millisecond)
		conn.Read(message)
		fmt.Println("r: ", string(message))
	}
}

//TODO:Create input output channels
func TCPclient(internalOrders chan string) {
	conn, err := Dial("tcp", serverIP+TCPPORT)
	for {
		que := <-internalOrders
		buffer, err := json.Marshal(que)
		if err != nil {
			fmt.Println("Problem encoding JSON object:", err)
		}
		conn.Write(buffer)
	}
}
