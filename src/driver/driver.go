package driver

import (
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

func Init(t ElevatorType, internal chan string, external chan string) {

	InitElev(t)
	fmt.Println("Maxfloor=", maxFloor)
	fmt.Println("Passed string to channel")
	SetMotorDir(DIR_DOWN)
	floor := -1
	go func(msg chan string) {
		for {
			floor = GetFloorSensor()
			if floor == 0 {
				SetMotorDir(DIR_STOP)
				msg <- "Motor dir stop"
				return
			}
		}
	}(internal)

	external <- "Arrived at floor 0"
	fmt.Println("passed arrived at floor")
	go floorIndicator(internal)
	go initButtonListners(internal)
}

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

func initButtonListners(msg chan string) {
	for i := 0; i < maxFloor; i++ {
		go checkButtonPress(msg, BUTTON_COMMAND, i)
		fmt.Println("Initialized command button:", i)
		if i < maxFloor-1 {
			go checkButtonPress(msg, BUTTON_CALL_UP, i)
			fmt.Println("Initialized up button:", i)
		}
		if i > 0 {
			go checkButtonPress(msg, BUTTON_CALL_DOWN, i)
			fmt.Println("Initialized down button:", i)
		}
	}
}

func checkButtonPress(msg chan string, buttonType Elev_button_type_t, floorLevel int) {
	for {
		if GetButtonSignal(buttonType, floorLevel) == 1 {
			SetButtonLamp(buttonType, floorLevel, 1)
			msg <- "Button Call signal recived"
		}
		time.Sleep(150 * time.Millisecond)
	}
}

/*func externalLights(externalQue chan string){
	for {
		for _, lights := range externalQue{
			if lights.up {
				SetButtonLamp(BUTTON_CALL_UP, lights.floor, 1)
			} else {
				SetButtonLamp(BUTTON_CALL_DOWN, lights.floor, 1)
			}
		}
	}
}*/
