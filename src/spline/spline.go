package spline

func Ts(n int) []float64 {
    ts := make([]float64, n+1)
    for i := range ts {
        if i == 0 {
            ts[i] = 0
        } else {
            ts[i] = float64(i) / float64(n)
        }
    }
    ts[n] = 1
    return ts
}

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

func copy2d(xs [][]float64) [][]float64 {
    ys := make([][]float64, len(xs))
    for i := range xs {
        ys[i] = make([]float64, len(xs[i]))
        copy(ys[i], xs[i])
    }
    return ys
}

func interpolate(
    degree,
    domain,
    d int,
    low,
    high,
    t float64,
    knots []float64,
    vs [][]float64,
) []float64 {
    if (t < low) || (t > high) {
        return []float64{}
    }
    s := 0
    for i := degree; i < domain; i++ {
        if (t >= knots[i]) && (t <= knots[i+1]) {
            s = i
            break
        }
    }
    vs = copy2d(vs)
    var alpha float64
    for l := 1; l < (degree + 1); l++ {
        for i := s; i > (s - degree - 1 + l); i-- {
            alpha = (t - knots[i]) / (knots[i+degree+1-l] - knots[i])
            for j := 0; j < (d + 1); j++ {
                vs[i][j] = ((1 - alpha) * vs[i-1][j]) + (alpha * vs[i][j])
            }
        }
    }
    y := make([]float64, d)
    for i := 0; i < d; i++ {
        y[i] = vs[s][i] / vs[s][d]
    }
    return y
}

func Spline(points [][]float64, ts []float64) [][]float64 {
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
            vs := make([][]float64, n)
            for i := 0; i < n; i++ {
                v := make([]float64, d+1)
                for j := 0; j < (d + 1); j++ {
                    if j < d {
                        v[j] = points[i][j]
                    } else {
                        v[j] = 1
                    }
                }
                vs[i] = v
            }
            ys := make([][]float64, len(ts))
            for i, t := range ts {
                ys[i] = interpolate(
                    degree,
                    domain,
                    d,
                    low,
                    high,
                    t*(high-low)+low,
                    knots,
                    vs,
                )
            }
            return ys
        }
    }
}
