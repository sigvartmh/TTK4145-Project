package main

import (
	"./driver"
	"fmt"
	"time"
)

const floorNumbers = 3
const buttonType = 2


func main() {

		
// Lag en chan med fullstedig array/struct/map med alle kommandoer og statuser fra heisen
floorCommand := make(chan int) 
	
	if driver.Init() == 1 {
		fmt.Println("Driver intialized")
	}
	
	go driver.Run()

	for (true) { // Testprogram
	// driver.GoToFloor(0)
	fmt.Println(<- floorCommand)
	time.Sleep(10*time.Millisecond)
	floorCommand <- 1
	
	}

	driver.GoToFloor(0)

	/*


		for {
			if driver.GetButtonSignal(0, 0) == 1 {
				driver.SetButtonLamp(0, 0, 1)
				driver.SetButtonLamp(2, 0, 1)
			}
			if driver.GetFloorSensor() == 2 {
				if floor != driver.GetFloorSensor() {
					fmt.Println("Arrived at floor 2")
					floor = driver.GetFloorSensor()
				}
			} else if driver.GetFloorSensor() == 3 {
			}

		}

	*/
}
