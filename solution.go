package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myInt int

func main() {
	var SidesTriangle myInt = 3
	var SidesSquare myInt = 4
	var SidesCircle myInt = 0

	fmt.Println(CalcSquare(10.0, SidesTriangle))
	fmt.Println(CalcSquare(10.0, SidesSquare))
	fmt.Println(CalcSquare(10.0, SidesCircle))
}

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	switch sidesNum {
	case 4:
		return sideLen * sideLen
	case 3:
		return math.Sqrt(3.0) / 4.0 * math.Pow(sideLen, 2)
	case 0:
		const pi float64 = 3.14
		return pi * sideLen * sideLen
	default:
		return 0
	}
}
