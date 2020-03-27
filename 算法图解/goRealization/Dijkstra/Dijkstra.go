package main

import "math"

func main() {
	//有序图
	directionMap := make(map[string](map[string]uint))
	//起点到各点已知的最短距离
	cost := make(map[string]uint)

	//路径，由终点反向推导
	path := make(map[string]string)

	directionMap["start"]["A"] = 6
	directionMap["start"]["B"] = 2
	directionMap["A"]["end"] = 1
	directionMap["B"]["A"] = 3
	directionMap["B"]["end"] = 5
	directionMap["end"] = map[string]uint{}

	cost["A"] = 6
	cost["B"] = 2
	cost["end"] = math.MaxUint32

	path["A"] = "start"
	path["B"] = "start"
	path["end"] = ""
}
