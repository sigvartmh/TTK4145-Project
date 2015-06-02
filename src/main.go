package main

import (
	. "../queue"
	"./driver"
	"./driver/src"
	"./network/tcp"
	"./network/udp"
	"fmt"
	"runtime"
	"time"
)

const (
	ET_COMEDI cwrapper.ElevatorType = iota
	ET_SIMULATOR
)

var BrodcastPort = os.Getenv("BROADCASTPORT")
var HearthbeatPort = os.Getenv("HEARTBEATPORT")
var TCPPort, _ = strconv.Atoi(os.Getenv("TCPPORT"))

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var state bool
	var queues Que
	b := make([]byte, 1024)
	queues = make(Que)
	external := make(chan QueItem)
	queue := make(chan Que)
	orders := make(chan QueItem)
	sendTCP := make(chan QueItem)
	reciveTCP := make(chan QueItem)
	backup := make(chan Que)
	master := make(chan bool)
	nomaster := make(chan bool)
	newmaster := make(chan string)

	_, _, err := ln.ReadFromUDP(b)
	saddr, _ := ResolveUDPAddr("udp", HearthbeatPort)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))
	_, rAddr, err := ln.ReadFromUDP(b)
	if err != nil {
		state = true
	}
	ln.Close()

	if state {
		go tcp.Server(TCPPort, reciveTCP)
		go tcp.Client(tcp.GetLocalIP(TCPPort).String(), sendTCP)
		go udp.Heartbeat(HearthbeatPort, master)
	} else {
		go tcp.Client(rAddr, to)
	}
	go udp.Server(BroadcastPort, backup)
	go driver.Init(ET_COMEDI, orders, external, queue)

	for {
		select {
		case order := <-external:
			sendTCP <- order
		case order := <-reciveTCP:
			queue <- order
		case bkp := <-backup:
			queue <- bkp
		case <-nomaster:
			alive := checkLowestIP(queue)
			if alive.state {
				newmaster <- alive.IP
			}
		default:
			if state {
				master <- state
			}
		}
		//else
	}
}
