package main

import (
    "fmt"
)

func check_dimensions(d int, xs [][]float32) bool {
    for _, x := range xs {
        if len(x) != d {
            return false
        }
    }
    return true
}

func interpolate(points [][]float32) []float32 {
    const degree int = 2
    var n int = len(points)
    if n <= degree {
        return []float32{}
    } else {
        // var d int = len(points[0])
        return points[degree - 1]
    }
}

func main() {
    points := [][]float32{
        {0.0, 1.0},
        {1.0, 0.5},
        {0.5, 0.75},
    }
    fmt.Println(interpolate(points))
}
