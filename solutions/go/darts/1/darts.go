package darts
import "math"

func Score(x, y float64) int {
		pointr := math.Sqrt(x*x + y*y)  
	if pointr <= 1 {
		return 10
	} else if pointr <= 5 {
		return 5
	} else if pointr <= 10 {
		return 1
	} else {          
		return 0
	}
}
