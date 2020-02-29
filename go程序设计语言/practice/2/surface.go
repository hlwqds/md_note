package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0 //坐标轴的范围，+-xyrange
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", HttpDrawSVG)
		fmt.Println(http.ListenAndServe("localhost:8000", nil))
	} else {
		fptr, err := os.Create("test.svg")
		if err != nil {
			os.Exit(1)
		}
		WriteSVGData(fptr)
		fptr.Close()
	}
}

func HttpDrawSVG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	WriteSVGData(w)
}

func WriteSVGData(fptr io.Writer) {
	fmt.Fprintf(fptr, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(fptr, "<polygon style='stroke: #0000ff' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(fptr, "</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	//z的变化决定了是峰顶还是谷底

	//将(x, y, z)等脚投射到二维svg绘图平面上， 坐标是(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return transferFloatInf(sx), transferFloatInf(sy)
}

//r = sqrt(x^2 + y^2)
//z = sin(r)/r
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	var result float64
	result = math.Sin(r) / r

	return result
}

func transferFloatInf(x float64) float64 {
	if math.IsInf(x, 1) {
		return math.MaxFloat64
	} else if math.IsInf(x, -1) {
		return -math.MaxFloat64
	} else {
		return x
	}
}
