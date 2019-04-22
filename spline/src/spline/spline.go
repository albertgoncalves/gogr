package spline

func sameLens(d int, xs [][]float64) bool {
    for _, x := range xs {
        if len(x) != d {
            return false
        }
    }
    return true
}

func floatRange(min, max int) []float64 {
    a := make([]float64, max-min)
    for i := range a {
        a[i] = float64(min + i)
    }
    return a
}

func Interpolate(points [][]float64, ts []float64) [][]float64 {
    const degree int = 2
    n := len(points)
    if n <= degree {
        return [][]float64{}
    } else {
        d := len(points[0])
        if !sameLens(d, points) {
            return [][]float64{}
        } else {
            knots := floatRange(0, n+degree+1)
            domain := len(knots) - 1 - degree
            low := knots[degree]
            high := knots[domain]
            ys := make([][]float64, len(ts))
            for k, t := range ts {
                t = t*(high-low) + low
                if (t < low) || (t > high) {
                    return [][]float64{}
                }
                var s int
                for i := degree; i < domain; i++ {
                    if (t >= knots[i]) && (t <= knots[i+1]) {
                        s = i
                        break
                    }
                }
                vs := make([][]float64, n)
                for i := 0; i < n; i++ {
                    var v = make([]float64, d+1)
                    for j := 0; j < (d + 1); j++ {
                        if j < d {
                            v[j] = points[i][j]
                        } else {
                            v[j] = 1
                        }
                    }
                    vs[i] = v
                }
                var alpha float64
                for l := 1; l < (degree + 1); l++ {
                    for i := s; i > (s - degree - 1 + l); i-- {
                        alpha = (t - knots[i]) /
                            (knots[i+degree+1-l] - knots[i])
                        for j := 0; j < (d + 1); j++ {
                            vs[i][j] = ((1 - alpha) * vs[i-1][j]) +
                                (alpha * vs[i][j])
                        }
                    }
                }
                y := make([]float64, d)
                for i := 0; i < d; i++ {
                    y[i] = vs[s][i] / vs[s][d]
                }
                ys[k] = y
            }
            return ys
        }
    }
}
