package main

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] >= cost[i] && canTravel(gas, cost, i) {
			return i
		}
	}
	return -1
}

func canTravel(gas []int, cost []int, start int) bool {
	tank := gas[start]
	end := ((start - 1) + len(gas)) % len(gas)
	for s := start; s != end; {
		tank -= cost[s]
		if tank < 0 {
			return false
		}
		s++
		s = s % len(gas)
		tank += gas[s]
	}
	return tank >= cost[end]
}
