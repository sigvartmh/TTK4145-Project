Project GO elevator TODO:

Motivation:
The Go approach

Don't communicate by sharing memory, share memory by communicating.

Reasoning:
Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously.

Every node/client should know of each other and have it stored in a structure.

Graceful degradation reduces the functionality to only perform core tasks while the system waits for repairs.

##Priority:
1. Server Client structures with for select in the network module
* Make a que structure (Sort of done in tcp.go)
*  Setup Queues that are used to ligth up external and internal buttons

##Secondary:
1. WatchDog for process(use UDP, file or enviromental variables)
* Use UDP to send to broadcast so that master could get the client IP
* Fix GoToFloor Function


#Draft Program structure
24.04.2015

###elevator
	for{
		select{
			case <-internalButton press:
				Add to internalQue(1st priority)
				internalQue chan <- button number
			case <- externalButton press:
				tcp input chan <- Send tcp msg to master
				udpBroadcast input chan <- Update que for all nodes
			case
	}

###network
	select{
		case msg := <-sendTCPmessage:
			send que to master chan <- msg
		case msg := <-recivedTCPmessage:
			recived request from maaster to handle que
		case <-boradcastUDPbackup:
			update master with state of elevator
			udpChan <- state
		case <-recivedUDP
	}

Known bugs:
Go routines spawns and exectues too many cgo calls making the processor quite busy

Data
MAX go routines

go button 3+3+4 = 10
go lampListner = 1
go main	= 1

12 Base threads?
