package driver

/*
#cgo CFLAGS: -std=c99 -g -Wall
#cgo LDFLAGS:  "/usr/lib/libcomedi.a" -lm
#include "io.h"
#include "channels.h"
#include "elev.h"
*/
import "C"
import "fmt"
import "time"

type elev_button_type_t int
type elev_motor_direction_t int

const (
	BUTTON_CALL_UP elev_button_type_t = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
)

const (
	DIRN_DOWN elev_motor_direction_t = -1
	DIRN_STOP                        = 0
	DIRN_UP                          = 1
)

// elevatorStatus := make(map[string]string)

// Eller struct?
//type elevatorStatus struct {
//    Current floor 1

//}

func Init() int {
	initOK := int(C.elev_init())
	for getFloorSensor() != 0 {
		setMotorDir(DIRN_DOWN)
		// Fail and report after 10 seconds
	        // For testing: fmt.Println(getFloorSensor())
		fmt.Println("kvakk")
		time.Sleep(1*time.Second) // In order to stop nicely on the floor


	}
	setMotorDir(DIRN_STOP)
	 // Start the go routine that will keep sending status updaates
	 // go senseElevatorStatus()	
	return initOK
}

func setMotorDir(dir elev_motor_direction_t) { // made private
	C.elev_set_motor_direction(C.elev_motor_direction_t(dir))
}

func getFloorSensor() int { // made private
	return int(C.elev_get_floor_sensor_signal())
}

func getButtonSignal(button elev_button_type_t, floor int) int { // made private
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func getStopSignal() int { // made private
	return int(C.elev_get_stop_signal())
}

func getObstructionSignal() int { // made private
	return int(C.elev_get_stop_signal())
}

func setFloorIndicator(floor int) { // made private
	C.elev_set_floor_indicator(C.int(floor))
}

func setButtonLamp(button elev_button_type_t, floor int, value int) { // made private
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func setStopLamp(value int) { // made private
	C.elev_set_stop_lamp(C.int(value))
}

func setDoorOpenLamp(value int) { // made private
	C.elev_set_door_open_lamp(C.int(value))
}

func GoToFloor(desiredFloor int) {
    currentFloor = GetFloorSensor()
    if desiredFloor == currentFloor {
        return
        } else {
            setMotorDir(desiredFloor - currentFloor)
            for desiredFloor != currentFloor {
                // Wait for how long?
        }
        return
    }
}

func updateFloorLights() { // Run as go routine from Init()
	// currentFloor := GetFloorSensor()
	// if currentFloor
	// setFloorIndicator(currentFloor)
}

/*
Public functions:
GoToFloor(floor int)
GetFloorSensor
ButtonPushedOnFloor() chan


buttonpress channel:

const (
floorButton3down type_t = iota
floorButton2up
floorButton2down
floorButton1up
floorButton1down
floorButton0up
elevatorButtonCommand3
elevatorButtonCommand2
elevatorButtonCommand1
elevatorButtonCommand0
elevatorButtonStop

go routine som sjekker for knapper "hele tiden" og putter dem ut på en channel som et map
*/
