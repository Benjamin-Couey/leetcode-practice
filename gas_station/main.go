/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/gas-station/description/
*/
package gas_station

/*
CanCompleteCircuit returns the starting gas station's index from which you
can travel around the circuit once in a clockwise direction (increasing the
index) defined by gas and cost. If there is no such gas station, returns -1.
There are n gas stations along a circular route. The ith station has gas[i] gas.
It costs cost[i] to travel from the ith station to the (i+1)th station. You
begin the journey with no gas at one gas station.
CanCompleteCircuit assumes that:
if there exists a solution, it is guaranteed to be unique,
n == gas.length == cost.length,
1 <= n <= 10^5,
and 0 <= gas[i], cost[i] <= 10^4.
*/
func CanCompleteCircuit(gas []int, cost []int) int {

	net_sum := 0
	net_gas := make([]int, len(gas))

	for index, gas_num := range gas {
		net_sum += gas_num - cost[index]
		net_gas[index] = gas_num - cost[index]
	}

	if net_sum < 0 {
		return -1
	}

	left_sum := 0
	return_sum := left_sum
	return_index := 0
	for index, _ := range net_gas {
		if index > 0 {
			left_sum += net_gas[index-1]
			if left_sum < return_sum {
				return_sum = left_sum
				return_index = index
			}
		}
	}

	return return_index
}
