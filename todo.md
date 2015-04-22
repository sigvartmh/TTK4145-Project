Project GO elevator TODO:
Priority:
#1 Server Client structures with for select in the network module
#2 Make a que structure (Sort of done in tcp.go)
#3 Setup Queues that are used to ligth up external and internal buttons
#4 
Secondary:
#1 WatchDog for process(use UDP, file or enviromental variables)
#2 Use UDP to send to broadcast so that master could get the client IP
#3 Fix GoToFloor Function


Every node/client should know of each other and have it stored in a structure. 

Graceful degradation reduces the functionality to only perform core tasks while the system waits for repairs.
