package main

import (
    "fmt"
    S "spline"
)

func main() {
    points := [][]float64{
        {-0.5, 5},
        {-2, 0},
        {0.5, -0.5},
        {0, 2},
        {3.5, 0},
        {-0.5, 0.5},
        {-2, 0},
        {-3, -1},
        {2, -0.5},
        {0, -2.75},
        {5, -5},
    }
    fmt.Println(S.Spline(points, []float64{0.0, 0.5, 1.0}))
}
