package main

// https://github.com/fogleman/gg

import (
    "fmt"
    "github.com/fogleman/gg"
    "math/rand"
    "os"
)

func random() float64 {
    return rand.Float64()*2 - 1
}

func point() (x, y float64) {
    return random(), random()
}

func drawCurve(dc *gg.Context) {
    dc.SetRGBA(0, 0, 0, 0.1)
    dc.FillPreserve()
    dc.SetRGB(0, 0, 0)
    dc.SetLineWidth(8)
    dc.Stroke()
}

func drawPoints(dc *gg.Context) {
    dc.SetRGBA(1, 0, 0, 0.5)
    dc.SetLineWidth(2)
    dc.Stroke()
}

func randomQuadratic(dc *gg.Context) {
    x0, y0 := point()
    x1, y1 := point()
    x2, y2 := point()
    dc.MoveTo(x0, y0)
    dc.QuadraticTo(x1, y1, x2, y2)
    drawCurve(dc)
    dc.MoveTo(x0, y0)
    dc.LineTo(x1, y1)
    dc.LineTo(x2, y2)
    drawPoints(dc)
}

func randomCubic(dc *gg.Context) {
    x0, y0 := point()
    x1, y1 := point()
    x2, y2 := point()
    x3, y3 := point()
    dc.MoveTo(x0, y0)
    dc.CubicTo(x1, y1, x2, y2, x3, y3)
    drawCurve(dc)
    dc.MoveTo(x0, y0)
    dc.LineTo(x1, y1)
    dc.LineTo(x2, y2)
    dc.LineTo(x3, y3)
    drawPoints(dc)
}

func lineToPoint(dc *gg.Context, point []float64) {
    dc.LineTo(point[0], point[1])
}

func lineThruPoints(dc *gg.Context, points [][]float64) {
    for _, point := range points {
        lineToPoint(dc, point)
    }
    drawPoints(dc)
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
        W = 1
        H = 1
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
            points := randomPoints(5)
            fmt.Println(points)
            lineThruPoints(dc, points)
            // if j%2 == 0 {
            //     randomCubic(dc)
            // } else {
            //     randomQuadratic(dc)
            // }
            dc.Pop()
        }
    }
    dc.SavePNG(fmt.Sprintf("%s/pngs/out.png", os.Getenv("GOPATH")))
}
