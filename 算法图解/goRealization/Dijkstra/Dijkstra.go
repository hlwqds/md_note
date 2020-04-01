package main

import (
	"fmt"
	"math"
)

type NeighborMap map[string]uint
type DirectionMap map[string](map[string]uint)

func findLowestCostNode(costs map[string]uint, handled map[string]bool) string {
	lowestCostNode := ""
	var lowestCost uint = math.MaxUint32

	for node, cost := range costs {
		if cost < lowestCost && handled[node] == false {
			lowestCostNode = node
			lowestCost = cost
		}
	}
	return lowestCostNode
}

func main() {
	//有序图
	//起点到各点已知的最短距离
	cost := make(map[string]uint)

	//路径，由终点反向推导
	path := make(map[string]string)

	handled := make(map[string]bool)
	directionMap := DirectionMap{
		"start": NeighborMap{
			"A": 6,
			"B": 2,
		},
		"A": NeighborMap{
			"end": 1,
		},
		"B": NeighborMap{
			"A":   3,
			"end": 5,
		},
		"end": nil,
	}

	cost["A"] = 6
	cost["B"] = 2
	cost["end"] = math.MaxUint32

	path["A"] = "start"
	path["B"] = "start"
	path["end"] = ""

	node := findLowestCostNode(cost, handled)
	for node != "" {
		fmt.Println(node)
		costValue := cost[node]
		neighbors := directionMap[node]
		for neighbor, neighborCost := range neighbors {
			newCost := costValue + neighborCost
			if newCost < cost[neighbor] {
				cost[neighbor] = newCost
				path[neighbor] = node
			}
		}
		handled[node] = true

		node = findLowestCostNode(cost, handled)
	}

	fmt.Println(path)
}
