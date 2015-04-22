package client

import "fmt"
import "net"
import "container/list"
import "bytes"

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

type Client struct {
	IP       string
	Input	 chan QueItem
	Output	 chan QueItem
	Conn	 net.Conn
	Quit	 chan bool
	ClientList *list.List
}

func (c *Client) Close() {
    c.Alive <- false
    c.Conn.Close()
    c.RemoveMe()
}

func Log(v ...interface{}) {
    fmt.Println(v...)
}

func ClientSender(client *Client) {
    for {
        select {
        case buff := client.Input:
        	Log("Client Sending: ", string(buffer), "to" client.IP)
        case <-client.Quit:
        	Log("Client:", client.IP, "Disconnected")
        	client.Conn.Close()
        	break
        }
    }
}

func ClientHandler(conn net.Conn, item chan QueItem, clientList *list.List) {
    buffer := make([]byte, 1024)
    bytesRead, error := conn.Read(buffer)
    if error != nil {
        Log("Client connection error: ", error)
    }
 
    name := string(buffer[0:bytesRead])
    newClient := &Client{name, make(chan string), ch, conn, make(chan bool), clientList}
 
    go ClientSender(newClient)
    go ClientReader(newClient)
    clientList.PushBack(*newClient)
    ch <-string(name + " has joined the chat")
}

func IOHandler(Input <-chan string, clientList *list.List) {
    for {
        Log("IOHandler: Waiting for input")
        input := <-Input
        Log("IOHandler: Handling ", input)
        for e := clientList.Front(); e != nil; e = e.Next() {
            client := e.Value.(Client)
            client.Incoming <-input
        }
    }
}