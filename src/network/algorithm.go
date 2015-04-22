package algorithm

func calculateCost(currentFloor, targetFloor, dir int) int {
	if dir > 0 {
		cost := targetFloor - currentFloor
	}
	else {
		cost := currentFloor - targetFloor
	}

	return cost
}
