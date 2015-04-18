package udp

import "fmt"
import "time"
import "net"

//TODO: every slave listen to the broadcast channel for backup que
//TODO: every slave sends a package to broadcast channel to register their IP
//TODO: Master sends a package to broadcast channel to notify the slaves


func udpSend(done chan bool, port , saddr *UDPAddr){
	conn, err := net.DialUDP("udp", nil, saddr)
	if err != nil {
		fmt.Println("Error connecting to" + saddr)
	}

	for {
		time.Sleep(1000*time.Millisecond)
		conn.Write([]byte("The cake is a lie"))
		fmt.Println("Msg sent on udp")
	} //Switch to for select loop
	done <- true

}

func udpRecive(done chan bool, port , saddr * net.UDPAddr) {
	buff := make([]byte, 1024)

	l, err := net.ListenUDP("udp4", saddr)
	if err != nil {
		fmt.Println("Error listening to" + saddr)
	}

	_,_, err = l.ReadFromUDP(buff)

	if err != nil {
			fmt.Println(err)
	}
	fmt.Println(string(buff[:]))

}