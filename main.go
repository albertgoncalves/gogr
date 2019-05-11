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
    dc.SetRGBA(0, 0, 0, 0.25)
    dc.SetLineWidth(1)
    dc.Stroke()
}

func drawCurve(dc *gg.Context) {
    dc.SetRGBA(0, 0, 0, 0.035)
    dc.FillPreserve()
    dc.SetRGB(0, 0, 0)
    dc.SetLineWidth(2)
    dc.Stroke()
}

func lineThruPoints(
    dc *gg.Context,
    f func(dc *gg.Context),
    points []float64,
) {
    var n int = len(points)
    if (n >= 4) && (n%2 == 0) {
        var m int = n / 2
        for i := 0; i < m; i++ {
            dc.LineTo(points[i*2], points[(i*2)+1])
        }
        f(dc)
    }
}

func randomPoints(n int) []float64 {
    points := make([]float64, n*2)
    for i := range points {
        points[i] = random()
    }
    return points
}

func main() {
    rand.Seed(3)
    const (
        S = 256
        W = 8
        H = 8
        M = 100
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
            n := rand.Intn(4) + 6
            points := randomPoints(n)
            lineThruPoints(dc, drawLines, points)
            lineThruPoints(
                dc,
                drawCurve,
                spline.Spline(points, n, 2, 4, spline.Ts(M)),
            )
            dc.Pop()
        }
    }
    dc.SavePNG(fmt.Sprintf("%s/pngs/main.png", os.Getenv("GOPATH")))
}
