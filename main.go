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
    const (
        S = 256
        W = 6
        H = 6
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
            n := rand.Intn(4) + 3
            points := randomPoints(n)
            lineThruPoints(dc, drawLines, points)
            lineThruPoints(
                dc,
                drawCurve,
                spline.Spline(points, n, 2, spline.Ts(M)),
            )
            dc.Pop()
        }
    }
    dc.SavePNG(fmt.Sprintf("%s/pngs/main.png", os.Getenv("GOPATH")))
}
