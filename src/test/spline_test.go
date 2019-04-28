package spline

import (
    "math"
    "testing"
    "spline"
)

func roundArray(xs []float64) {
    for i := range xs {
        xs[i] = math.Round(xs[i]*100) / 100
    }
}

func compareArrays(xs, ys []float64) bool {
    if len(xs) != len(ys) {
        return false
    }
    for i, x := range xs {
        if x != ys[i] {
            return false
        }
    }
    return true
}

func TestTs(t *testing.T) {
    expected := []float64{
        0.0,
        0.1,
        0.2,
        0.3,
        0.4,
        0.5,
        0.6,
        0.7,
        0.8,
        0.9,
        1.0,
    }
    result := spline.Ts(10)
    if !compareArrays(result, expected) {
        t.Errorf(
            "TestTs\n"+
                "Ts(10)\nresult:\n\t%g\nexpected:\n\t%g\n\n",
            result,
            expected,
        )
    }
}

func TestSplineValid(t *testing.T) {
    ts := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
    points := []float64{
        -1.0, 0.0,
        -0.5, 0.5,
        0.5, -0.5,
        1.0, 0.0,
    }
    expected := []float64{
        -0.75, 0.25,
        -0.64, 0.32,
        -0.51, 0.33,
        -0.36, 0.28,
        -0.19, 0.17,
        0, 0,
        0.19, -0.17,
        0.36, -0.28,
        0.51, -0.33,
        0.64, -0.32,
        0.75, -0.25,
    }
    result := spline.Spline(points, 4, 2, ts)
    roundArray(result)
    if !compareArrays(result, expected) {
        t.Errorf(
            "TestSplineValid\n"+
                "Spline(points, 4, 2, ts)\nresult:\n\t%g\nexpected:\n\t%g\n\n",
            result,
            expected,
        )
    }
}

func TestSplineIncompletePoints(t *testing.T) {
    ts := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
    points := []float64{
        -1.0, 0.0,
        -0.5, 0.5,
        0.5, -0.5,
        1.0,
    }
    expected := []float64{}
    result := spline.Spline(points, 4, 2, ts)
    if !compareArrays(result, expected) {
        t.Errorf(
            "TestSplineIncompletePoints\n"+
                "Spline(points, 4, 2, ts)\nresult:\n\t%g\nexpected:\n\t%g\n\n",
            result,
            expected,
        )
    }
}

func TestSplineFewPoints(t *testing.T) {
    ts := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
    points := []float64{
        -1.0, 0.0,
        -0.5, 0.5,
    }
    expected := []float64{}
    result := spline.Spline(points, 2, 2, ts)
    if !compareArrays(result, expected) {
        t.Errorf(
            "TestSplineFewPoints\n"+
                "Spline(points, 2, 2, ts)\nresult:\n\t%g\nexpected:\n\t%g\n\n",
            result,
            expected,
        )
    }
}

func BenchmarkSpline(b *testing.B) {
    ts := spline.Ts(1000)
    points := []float64{-1.0, 0.0, -0.5, 0.5, 0.5, -0.5, 1.0, 0.0}
    for i := 0; i < b.N; i++ {
        spline.Spline(points, 4, 2, ts)
    }
}
