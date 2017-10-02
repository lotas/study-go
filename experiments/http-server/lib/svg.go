package lib

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.6        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

//GenSvg generate SVG with specific paint "fun" function
func GenSvg(fun string) string {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, fun)
			bx, by := corner(i, j, fun)
			cx, cy := corner(i, j+1, fun)
			dx, dy := corner(i+1, j+1, fun)
			svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += "</svg>"
	return svg
}

func corner(i, j int, fun string) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y, fun)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64, fun string) float64 {
	r := 1.0
	switch fun {
	case "hypot":
		r = math.Hypot(x, y) // distance from (0,0)
	case "copy":
		r = math.Copysign(x, y) // distance from (0,0)
	case "egg":
		r = math.Sin(y) * math.Cos(x)
	default:
		r = -math.Exp(math.Sin(x) * math.Cos(y))
	}
	return math.Sin(r) / r
}
