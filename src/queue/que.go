package que

import "sync"

type QueItem struct {
	IP       string
	Floor    int
	Type     int //Up = 0, Down = 1, Command = 2
	Complete bool
}

type Que struct {
	External []QueItem
}

type InternalQue struct {
	Internal []QueItem
	Ordered  []QueItem
	Mu       sync.Mutex
}

func sortQue(q *internalQue) {

}
