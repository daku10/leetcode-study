package main

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	l := len(gas)
	startPoint := 0
	currentGas := 0
	for i := 0; i < l; i++ {
		sum += (gas[i] - cost[i])
		currentGas += (gas[i] - cost[i])
		if currentGas < 0 {
			startPoint = i + 1
			currentGas = 0
		}
	}
	if sum < 0 {
		return -1
	}
	return startPoint
}
