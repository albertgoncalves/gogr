package spline

import (
    "math"
    "testing"
)

func compareArrays(xs, ys []float64) bool {
    if len(xs) != len(ys) {
        return false
    }
    for i := range xs {
        if xs[i] != ys[i] {
            return false
        }
    }
    return true
}

func TestSameLensTrue(t *testing.T) {
    var expected bool = true
    points := [][]float64{
        {0.0, 1.0},
        {1.0, 0.5},
        {0.5, 0.75},
    }
    result := sameLens(len(points[0]), points)
    if !expected {
        t.Errorf(
            "sameLens was incorrect, got: %t, wanted: %t",
            result,
            expected,
        )
    }
}

func TestSameLensFalse(t *testing.T) {
    var expected bool = false
    points := [][]float64{
        {0.0, 1.0},
        {1.0, 0.5},
        {0.5},
    }
    result := sameLens(len(points[0]), points)
    if expected {
        t.Errorf(
            "sameLens was incorrect, got: %t, wanted: %t",
            result,
            expected,
        )
    }
}

func TestFloatRange(t *testing.T) {
    expected := []float64{0, 1, 2, 3, 4, 5}
    result := floatRange(0, 6)
    if !compareArrays(result, expected) {
        t.Errorf(
            "floatRange was incorrect, got: %v, wanted: %v",
            result,
            expected,
        )
    }
}

func roundArray(xs []float64) []float64 {
    for i := range xs {
        xs[i] = math.Round(xs[i]*100) / 100
    }
    return xs
}

func TestInterpolateValid(t *testing.T) {
    ts := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
    points := [][]float64{{-1.0, 0.0}, {-0.5, 0.5}, {0.5, -0.5}, {1.0, 0.0}}
    expected := [][]float64{
        {-0.75, 0.25},
        {-0.64, 0.32},
        {-0.51, 0.33},
        {-0.36, 0.28},
        {-0.19, 0.17},
        {0, 0},
        {0.19, -0.17},
        {0.36, -0.28},
        {0.51, -0.33},
        {0.64, -0.32},
        {0.75, -0.25},
    }
    result := Interpolate(points, ts)
    for i := range ts {
        result[i] = roundArray(result[i])
    }
    for i := range result {
        if !compareArrays(result[i], expected[i]) {
            t.Errorf(
                "Interpolate was incorrect, got: %v, wanted: %v",
                result,
                expected,
            )
            break
        }
    }
}
