package celflag

import (
	"flag"
	"fmt"

	tmpconv "github.com/my/repo/go程序设计语言/practice/7-interface/tmpcov"
)

type celsiusFlag struct{ tmpconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var uints string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &uints)
	switch uints {
	case "C":
		f.Celsius = tmpconv.Celsius(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tmpconv.Celsius, usage string) *tmpconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
