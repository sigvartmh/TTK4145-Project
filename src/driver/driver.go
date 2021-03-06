package driver

import (
	"../network/udp"
	"../queue"
	. "./cwrapper/"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	ET_COMEDI ElevatorType = iota
	ET_SIMULATOR
)

const (
	BUTTON_CALL_UP Elev_button_type_t = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
)

const (
	DIR_DOWN Elev_motor_direction_t = -1
	DIR_STOP                        = 0
	DIR_UP                          = 1
)

type floorState struct {
	mu    sync.Mutex
	level int
	dir   int
}

//TODO: Needs fixing for the button listners, need to be
//multiplexed properly see http://talks.golang.org/2012/concurrency.slide#27
var maxFloor, _ = strconv.Atoi(os.Getenv("FLOORS"))
var floor floorState

//var queue que.InternalQue

func Init(t ElevatorType, internal, external chan QueItem, externalQue chan Que) {

	done := make(chan bool)
	InitElev(t)
	fmt.Println("Maxfloor=", maxFloor)
	SetMotorDir(DIR_DOWN)
	floor := -1
	go func(done chan bool) {
		for {
			floor = GetFloorSensor()
			if floor == 0 {
				SetMotorDir(DIR_STOP)
				done <- true
				return
			}
		}
	}(done)
	fmt.Println("passed arrived at floor")
	go floorIndicator(internal)
	<-done
	go initButtonListners(internal)
	go externalLights(externalQue)
}

//consider swapping time.sleep with
//select{ case <-time.After(time.Second * 3)}
func GoToFloor(value int) { //chan?
	//var floor int
	fmt.Println("Entered GoToFloor")
	fmt.Println("Value:", value, " Floor:", floor)
	for {
		/*if GetFloorSensor() != -1 {
			floor = GetFloorSensor()
		}*/
		fmt.Println("Value:", value, " Floor:", floor.level)
		fmt.Println("value > floor", value > floor.level)
		floor.mu.Lock()
		switch {
		case value > floor.level:
			fmt.Println("Entered GoToFloorUP")
			SetMotorDir(DIR_UP)

		case floor.level < value: //TODO:Find out why it cannot enter here
			fmt.Println("Entered GoToFloorDOWN")
			SetMotorDir(DIR_DOWN)

		case floor.level == value:
			defer floor.mu.Unlock()
			fmt.Println("Entered GoToFloorSTOP")
			SetMotorDir(DIR_STOP)
			SetFloorIndicator(1)
			time.Sleep(3 * time.Second)
			SetFloorIndicator(0)
			return
		}
		floor.mu.Unlock()
	}
}

func floorIndicator(msg chan string) {

	var lastFloor int = 0
	floor.mu.Lock()
	floor.level = -1
	defer floor.mu.Unlock()
	for {
		floor.mu.Lock()
		floor.level = GetFloorSensor()
		if floor.level == -1 {
			floor.level = lastFloor
		}
		defer floor.mu.Unlock()
		if floor.level != -1 && floor.level != lastFloor {
			SetFloorIndicator(floor.level)
			msg <- "Floor indicator set too floor"
			lastFloor = floor.level
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func initButtonListners(internalOrder, externalOrder chan QueItem) {
	for i := 0; i < maxFloor; i++ {
		go checkButtonPress(internalOrder, BUTTON_COMMAND, i)
		fmt.Println("Initialized command button:", i)
		if i < maxFloor-1 {
			go checkButtonPress(externalOrder, BUTTON_CALL_UP, i)
			fmt.Println("Initialized up button:", i)
		}
		if i > 0 {
			go checkButtonPress(externalOrder, BUTTON_CALL_DOWN, i)
			fmt.Println("Initialized down button:", i)
		}
	}
}

func checkButtonPress(order chan QueItem, buttonType Elev_button_type_t, floorLevel int) {
	for {
		if GetButtonSignal(buttonType, floorLevel) == 1 {
			order <- QueItem{udp.GetLocalIP(), buttonType, floorLevel, false}
			if buttonType == BUTTON_COMMAND {
				SetButtonLamp(BUTTON_COMMAND, floorLevel, 1)
			}
		}
		time.Sleep(150 * time.Millisecond)
	}
}

func externalLights(q chan Que) {
	var queue Que
	for {
		queue = <-q
		for _, lights := range q.External {
			if lights.Type == 0 {
				SetButtonLamp(BUTTON_CALL_UP, lights.Floor, 1)
			} else {
				SetButtonLamp(BUTTON_CALL_DOWN, lights.Floor, 1)
			}
		}
		time.Sleep(100 * time.Millisecond)
		for i := 0; i < maxFloor; i++ {
			if i < maxFloor-1 {
				SetButtonLamp(BUTTON_CALL_UP, i, 0)
			}
			if i > 0 {
				SetButtonLamp(BUTTON_CALL_DOWN, i, 0)
			}
		}
	}
}
