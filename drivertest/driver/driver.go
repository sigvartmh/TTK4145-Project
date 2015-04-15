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
		fmt.Println("Kjører heisen ned til første ... vent litt")
		time.Sleep(500*time.Millisecond) // In order to stop nicely on the floor
	}
	setMotorDir(DIRN_STOP)
	return initOK
}

func Run() {
	var currentFloor int
	for (true) {
		fmt.Println("polling")
		if currentFloor != getFloorSensor() {
			currentFloor = getFloorSensor()
			fmt.Println( currentFloor ) // replace with update of status array
			if currentFloor != -1 {setFloorIndicator(currentFloor)}
		}
		// replace the lines below with update of status array or something similar put on a channel
		fmt.Println(getButtonSignal(BUTTON_COMMAND, 0)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_COMMAND, 1)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_COMMAND, 2)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_COMMAND, 3)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_CALL_UP, 0)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_CALL_UP, 1)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_CALL_UP, 2)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_CALL_DOWN, 1)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_CALL_DOWN, 2)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getButtonSignal(BUTTON_CALL_DOWN, 3)) // Make consts: FLOOR-0 etc.?
		fmt.Println(getStopSignal())
		//fmt.Println(int(C.elev_get_stop_signal()))
		//fmt.Println(C.elev_get_stop_signal())
	
	
		/*
		if getButtonSignal(BUTTON_COMMAND, 0) != 0 {
			floorCommand <- 0
		}
	
		if getButtonSignal(BUTTON_COMMAND, 1) != 0 {
			floorCommand <- 1
		}
	
		if getButtonSignal(BUTTON_COMMAND, 2) != 0 {
			GoToFloor(2)
		}
	
		if getButtonSignal(BUTTON_COMMAND, 3) != 0 {
			GoToFloor(3)
		}
		*/
	
 		time.Sleep(100*time.Millisecond)
	}
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

func SetButtonLamp(button elev_button_type_t, floor int, value int) {
	// These lights must only be turned on from main(), as they act as a confirmation that the button press was logged
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func SetStopLamp(value int) { 
	// These lights must only be turned on from main(), as they act as a confirmation that the button press was logged
	C.elev_set_stop_lamp(C.int(value))
}

func setDoorOpenLamp(value int) { // made private
	C.elev_set_door_open_lamp(C.int(value))
}

func GoToFloor(desiredFloor int) {
    if desiredFloor == getFloorSensor() {
        return
    } else if desiredFloor > getFloorSensor(){
            setMotorDir(DIRN_UP)
    } else {
            setMotorDir(DIRN_DOWN)
    }
    for desiredFloor != getFloorSensor()  {
	// Lag timeout som returnerer en feilmelding om heisen ikke når frem
    }

	setMotorDir(DIRN_STOP)
        return
}



/*
buttonpress channel:

const (
Button3down type_t = iota
Button2up
Button2down
Button1up
Button1down
Button0up
elevatorButtonCommand3
elevatorButtonCommand2
elevatorButtonCommand1
elevatorButtonCommand0
elevatorButtonStop

*/
