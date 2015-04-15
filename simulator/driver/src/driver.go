package driver

/*
#cgo CFLAGS: -std=c99 -Ilib
#include "io.h"
#include "channels.h"
#include "elev.h"
*/
import "C"
import (
	// f		 "time"
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
	ET_COMIDI ElevatorType = iota
	ET_SIMULATOR
)

func initElev(elevType ElevatorType) int {
	return int(C.elev_init(C.ElevatorType(elevType)))
}

func Init(internal chan string, external chan string) {
	//success = init(1)
	initElev(ET_SIMULATOR)
	fmt.Println("Maxfloor= ", maxFloor)
	//internal <- "Starting simulator"
	//fmt.Println("message sent to main")
	setMotorDir(DIR_DOWN)
	go func() {
		for {
			floor = GetFloorSensor()
			if floor == 0 {
				setMotorDir(DIR_STOP)
				internal <- "Arrived at floor 0"
				return
			}
		}
	}()

	go floorIndicator(internal)
	go internalButtonPress(internal)
	go externalButtonPress(external)
}

/*Initializes the elevator */
/*func initElev(type_t ElevatorType) int {
	return int(C.elev_init(C.ElevatorType(C.ElevatorType(type_t)))
}*/

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

//func GoToFloor(value int){}

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
		msg <- "Checked externalButtonPress"
		//time.Sleep(25 * time.Millisecond)
	}
	//done <- true

}

func floorIndicator(msg chan string) {

	var lastFloor int = 0

	for {
		floor = GetFloorSensor()
		if floor != -1 && floor != lastFloor {
			SetFloorIndicator(floor)
			msg <- "Floor indicator set too floor"
			lastFloor = floor
		}
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
