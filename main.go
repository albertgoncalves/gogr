package main

// https://github.com/fogleman/gg

import (
    "fmt"
    "github.com/fogleman/gg"
    "math/rand"
    "os"
    "spline"
)

func random() float64 {
    return rand.Float64()*2 - 1
}

func drawLines(dc *gg.Context) {
    dc.SetRGBA(1, 0, 0, 0.5)
    dc.SetLineWidth(2)
    dc.Stroke()
}

func drawCurve(dc *gg.Context) {
    dc.SetRGBA(0, 0, 0, 0.1)
    dc.FillPreserve()
    dc.SetRGB(0, 0, 0)
    dc.SetLineWidth(3)
    dc.Stroke()
}

func lineToPoint(dc *gg.Context, point []float64) {
    if len(point) > 1 {
        dc.LineTo(point[0], point[1])
    }
}

func lineThruPoints(dc *gg.Context, f func(dc *gg.Context),
    points [][]float64) {
    for _, point := range points {
        lineToPoint(dc, point)
    }
    f(dc)
}

func randomPoints(n int) [][]float64 {
    points := make([][]float64, n)
    for i := range points {
        points[i] = []float64{random(), random()}
    }
    return points
}

func main() {
    const (
        S = 256
        W = 6
        H = 6
    )
    dc := gg.NewContext(S*W, S*H)
    dc.SetRGB(1, 1, 1)
    dc.Clear()
    for j := 0; j < H; j++ {
        for i := 0; i < W; i++ {
            x := float64(i)*S + S/2
            y := float64(j)*S + S/2
            dc.Push()
            dc.Translate(x, y)
            dc.Scale(S/2, S/2)
            n := rand.Intn(4) + 3
            points := randomPoints(n)
            lineThruPoints(dc, drawLines, points)
            lineThruPoints(
                dc,
                drawCurve,
                spline.Spline(points, spline.Ts(100)),
            )
            dc.Pop()
        }
    }
    dc.SavePNG(fmt.Sprintf("%s/pngs/main.png", os.Getenv("GOPATH")))
}
