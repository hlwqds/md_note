package main

import (
	"fmt"
	"time"

)

//Rocket speed
type Rocket struct {
	Speed float64
}

//Launch speed
func (r *Rocket) Launch() {
	fmt.Println(r.Speed)
}

func main() {
	r := Rocket{
		Speed: 0.1,
	}
	fmt.Println(r.Speed)
	//注意AfterFunc并不会造成阻塞
	timer := time.AfterFunc(6*time.Second, r.Launch)

	timer.Stop()
	time.Sleep(7 * time.Second)
}
