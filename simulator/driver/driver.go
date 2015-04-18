package driver

import (
	."./src"
	"os"
	"time"
	"fmt"
	"strconv"
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

var maxFloor, _ = strconv.Atoi(os.Getenv("FLOORS"))


func Init(t ElevatorType, internal chan string, external chan string) {

    InitElev(t)
	fmt.Println("Maxfloor=", maxFloor)
	//internal <- "Starting simulator"
    fmt.Println("Passed string to channel")
	SetMotorDir(DIR_DOWN)
	floor := -1
    go func(msg chan string){
        /*for(GetFloorSensor() != 0){
            SetMotorDir(DIR_DOWN)
        }
        SetMotorDir(DIR_STOP)
        msg <- "Motor dir stop"*/
        for{
        	floor = GetFloorSensor()
        	if floor == 0 {
        		SetMotorDir(DIR_STOP)
        		msg <- "Motor dir stop"
        		return
        	}
        }
    }(internal)
	//floor := GetFloorSensor()
	external <- "Arrived at floor 0"
    fmt.Println("passed arrived at floor")
    go floorIndicator(internal)
    go initButtonListners(internal)

}

/*Initializes the elevator */
/*func initElev(type_t ElevatorType) int {
	return int(C.elev_init(C.ElevatorType(C.ElevatorType(type_t)))
}*/

func GoToFloor(value int){
	var floor int
	fmt.Println("Entered GoToFloor")
	fmt.Println("Value:", value, " Floor:", floor)
	for{
		if GetFloorSensor() != -1{
			floor = GetFloorSensor()
		}
		fmt.Println("Value:", value, " Floor:", floor)
		fmt.Println("value > floor", value > floor)
		switch{
			case value > floor:
				fmt.Println("Entered GoToFloorUP")
				SetMotorDir(DIR_UP)

			case floor < value: //TODO:Find out why it cannot enter here
				fmt.Println("Entered GoToFloorDOWN")
				SetMotorDir(DIR_DOWN)

			case floor == value:
				fmt.Println("Entered GoToFloorSTOP")
				SetMotorDir(DIR_STOP)
		}
	}
}

func externalButtonPress(msg chan string) {

	var i int = 0

	for {
		if i < maxFloor-1 {
			if GetButtonSignal(BUTTON_CALL_UP, i) == 1 {
				SetButtonLamp(BUTTON_CALL_UP, i, 1)
				msg <- "Button Call up Signal recived"
				//time.Sleep(150 * time.Millisecond)
			}
		}
		if i > 0 {
			if GetButtonSignal(BUTTON_CALL_DOWN, i) == 1 {
				SetButtonLamp(BUTTON_CALL_DOWN, i, 1)
				msg <- "Button Call downSignal recived"
				//time.Sleep(150 * time.Millisecond)
			}
		}
		i++
		i = i % maxFloor
		//time.Sleep(25 * time.Millisecond)
	}
	//done <- true

}

func floorIndicator(msg chan string) {

	var lastFloor int = 0
	var floor = GetFloorSensor()

	for {
		floor = GetFloorSensor()
		if floor != -1 && floor != lastFloor {
			SetFloorIndicator(floor)
			msg <- "Floor indicator set too floor"
			lastFloor = floor
		}
        time.Sleep(500 *time.Millisecond)
	}
}

func internalButtonPress(msg chan string) {
	var i int = 0
	for {
		if GetButtonSignal(BUTTON_COMMAND, i) == 1 {
			SetButtonLamp(BUTTON_COMMAND, i, 1)
			msg <- "Button Command Signal recived"
			//time.Sleep(150* time.Millisecond)
		}
		i++
		i = i % maxFloor
		msg <- "Checked externalButtonPress"
		//time.Sleep(25 * time.Millisecond)
	}
}

func initButtonListners(msg chan string){
    for i := 0; i < maxFloor; i++ {
        go checkButtonPress(msg, BUTTON_COMMAND, i)
        fmt.Println("Initialized command button:", i)
        if(i < maxFloor-1){
        	go checkButtonPress(msg, BUTTON_CALL_UP, i)
            fmt.Println("Initialized up button:", i)
		}
        if(i > 0){
        	go checkButtonPress(msg, BUTTON_CALL_DOWN, i)
        	fmt.Println("Initialized down button:", i)
    	}
    }
}

func checkButtonPress(msg chan string, buttonType Elev_button_type_t, floorLevel int){
    for {
        if GetButtonSignal(buttonType, floorLevel) == 1 {
            SetButtonLamp(buttonType, floorLevel, 1)
            msg <- "Button Call signal recived"
        }
        //time.Sleep(150* time.Millisecond)
    }
}