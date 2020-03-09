package main

import (
	"fmt"
	"go/types"

)

//Point xy
type Point struct {
	X, Y float64
}

//Add add
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

//Sub sub
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func main() {
	var op func(p, q Point) Point

	p := Point{1, 2}
	q := Point{4, 7}
	op = Point.Add
	fmt.Println(op(p, q))
}

func ivrDetailHandle() {
	new_callpush := getCallPush
	old_call_detail := getCallDetail()
	if call_detail != nil {
		if (type == 呼入 || type == 查询类型) && old_call_detail.status == 0{
			//包括called
			updateDetail(new_callpush)
		} else {
			不更新
		}
	} else {
		insertDetail(new_callpush)
	}
	if 是交互式查询信息 {

	}
}
