package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z, zd := 10.0, 10.0
    for {
        zd = z
        z = z - (((z * z) - x) / (2 * z))
        fmt.Println(math.Abs(z - zd))
        if math.Abs(z - zd) < 0.000001{
            break;
        }
    }
    return z

}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
