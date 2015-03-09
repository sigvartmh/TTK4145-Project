package main

import (
	"./driver"
	"fmt"
)

const floorNumbers = 3
const buttonType = 2

func main() {
	if driver.Init() == 1 {
		fmt.Println("Driver intialized")
	}

	driver.SetMotorDir(0)
	floor := 0
	driver.SetButtonLamp(1, 1, 0)
	driver.SetButtonLamp(2, 1, 1)

	for {
		if driver.GetButtonSignal(0, 0) == 1 {
			driver.SetButtonLamp(0, 0, 1)
			driver.SetButtonLamp(2, 0, 1)
		}
		if driver.GetFloorSensor() == 2 {
			if floor != driver.GetFloorSensor() {
				fmt.Println("Arrived at floor 2")
				floor = driver.GetFloorSensor()
				driver.SetMotorDir(0)
			}
			driver.SetMotorDir(1)
		} else if driver.GetFloorSensor() == 3 {
			driver.SetMotorDir(-1)
		}

	}

}
