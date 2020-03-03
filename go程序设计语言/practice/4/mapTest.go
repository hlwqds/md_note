package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages1 := map[string]int{
		"huanglin":  24,
		"huangmian": 25,
	}
	ages["huanglin"] = 24
	ages["huangmian"] = 25

	fmt.Println(ages)
	fmt.Println(ages1)

	delete(ages, "huanglin")

	age, ok := ages["caixukun"]
	if !ok {
		fmt.Println(age)
	}

	if age2, ok := ages["caixukun"]; !ok {
		fmt.Println(age2)
	}
}
