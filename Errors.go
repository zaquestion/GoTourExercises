package main

import (
	"fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(f float64) (float64, error) {

    if f < 0 {
        return -1, ErrNegativeSqrt(f)
    }
	z, zd := f, f
	for {
		zd = z
		z = z - (((z * z) - f) / (2 * z))
		fmt.Println(math.Abs(z - zd))
		if math.Abs(z-zd) < 0.000001 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
