Project GO elevator TODO:

Motivation:
The Go approach

Don't communicate by sharing memory, share memory by communicating.

Reasoning:
Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously.

Every node/client should know of each other and have it stored in a structure. 

Graceful degradation reduces the functionality to only perform core tasks while the system waits for repairs.

Priority:
#1 Server Client structures with for select in the network module
#2 Make a que structure (Sort of done in tcp.go)
#3 Setup Queues that are used to ligth up external and internal buttons
#4 
Secondary:
#1 WatchDog for process(use UDP, file or enviromental variables)
#2 Use UDP to send to broadcast so that master could get the client IP
#3 Fix GoToFloor Function





select{
	case TCPmessage:
	case UDPmessage:
	case UDPbackup:
}

MAX go routines

go button 3+3+4 = 10
go lampListner = 1
go main	= 1

12 Base threads?