package driver

/*
#cgo CFLAGS: -std=c99 -Ilib
#include "io.h"
#include "channels.h"
#include "elev.h"
*/
import "C"
import (
	//"time"
	"fmt"
	"os"
	"strconv"
)

var maxFloor, _ = strconv.Atoi(os.Getenv("FLOORS"))
var floor int

type elev_button_type_t int
type elev_motor_direction_t int
type ElevatorType int

const (
	BUTTON_CALL_UP elev_button_type_t = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
)

const (
	DIR_DOWN elev_motor_direction_t = -1
	DIR_STOP                        = 0
	DIR_UP                          = 1
)

const (
	ET_COMEDI ElevatorType = iota
	ET_SIMULATOR
)

func Init(internal chan string, external chan string) {

    initElev(ET_SIMULATOR)
	fmt.Println("Maxfloor=", maxFloor)
	//internal <- "Starting simulator"
    fmt.Println("Passed string to channel")
	setMotorDir(DIR_DOWN)
    go func(msg chan string){
        for(GetFloorSensor() != 0){
            setMotorDir(DIR_DOWN)
        }
        setMotorDir(DIR_STOP)
        msg <- "Motor dir stop"
        return
    }(internal)
	floor = GetFloorSensor()
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
	fmt.Println("Entered GoToFloor")
	fmt.Println("Value:", value, " Floor:", floor)
	for{
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
		//time.Sleep(25 * time.Millisecond)
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

	for {
		floor = GetFloorSensor()
		if floor != -1 && floor != lastFloor {
			SetFloorIndicator(floor+1)
			msg <- "Floor indicator set too floor"
			lastFloor = floor
	}
        //time.Sleep(150 *time.Millisecond)
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
        if(i < maxFloor-1){ go checkButtonPress(msg, BUTTON_CALL_UP, i) }
        if(i > 0){ go checkButtonPress(msg, BUTTON_CALL_DOWN, i) }
    }
}

func checkButtonPress(msg chan string, buttonType elev_button_type_t, floorLevel int){
    for {
        if GetButtonSignal(buttonType, floorLevel) == 1 {
            SetButtonLamp(buttonType, floorLevel, 1)
            msg <- "Button Call signal recived"
        }
    }
}

func initElev(elevType ElevatorType) int {
	return int(C.elev_init(C.ElevatorType(elevType)))
}

func setMotorDir(dirn elev_motor_direction_t) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(dirn))
}

func GetFloorSensor() int {
	return int(C.elev_get_floor_sensor_signal())
}

func GetButtonSignal(button elev_button_type_t, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func GetStopSignal() int {
	return int(C.elev_get_stop_signal())
}

func GetObstructionSignal() int {
	return int(C.elev_get_stop_signal())
}

func SetFloorIndicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func SetButtonLamp(button elev_button_type_t, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func SetStopLamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func SetDoorOpenLamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}