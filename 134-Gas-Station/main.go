package main

func canCompleteCircuit(gas []int, cost []int) int {
	candidates, exists := findCandidates(gas, cost)
	if len(gas) == 1 {
		if gas[0] >= cost[0] {
			return 0
		}
		return -1
	}
	if !exists {
		return -1
	}
	for _, c := range candidates {
		if canTravel(gas, cost, c) {
			return c
		}
	}
	return -1
}

func findCandidates(gas []int, cost []int) ([]int, bool) {
	var candidates []int
	sum := 0
	for i := 0; i < len(gas); i++ {
		ci := gas[i] - cost[i]
		sum += ci
		if ci < 0 {
			j := ((i + 1) % len(gas))
			if gas[j]-cost[j] > 0 {
				candidates = append(candidates, j)
			}
		}
	}
	if sum < 0 {
		return nil, false
	}
	return candidates, true
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
		s %= len(gas)
		tank += gas[s]
	}
	return tank >= cost[end]
}
