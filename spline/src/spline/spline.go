package spline

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

func matrix(n, m int) [][]float64 {
    a := make([][]float64, n)
    for i := 0; i < n; i++ {
        a[i] = make([]float64, m)
    }
    return a
}

func Interpolate(points [][]float64, t float64) []float64 {
    const degree int = 2
    n := len(points)
    if n <= degree {
        return []float64{}
    } else {
        d := len(points[0])
        if !sameLens(d, points) {
            return []float64{}
        } else {
            knots := floatRange(0, n+degree+1)
            domain := len(knots) - 1 - degree
            low := knots[degree]
            high := knots[domain]
            t = t*(high-low) + low
            if (t < low) || (t > high) {
                return []float64{}
            }
            var s int
            for i := degree; i < domain; i++ {
                if (t >= knots[i]) && (t <= knots[i+1]) {
                    s = i
                    break
                }
            }
            v := matrix(n, d+1)
            for i := 0; i < n; i++ {
                for j := 0; j < (d + 1); j++ {
                    if j < d {
                        v[i][j] = points[i][j]
                    } else {
                        v[i][j] = 1
                    }
                }
            }
            var alpha float64
            for l := 1; l < (degree + 1); l++ {
                for i := s; i > (s - degree - 1 + l); i-- {
                    alpha = (t - knots[i]) /
                        (knots[i+degree+1-l] - knots[i])
                    for j := 0; j < (d + 1); j++ {
                        v[i][j] = ((1 - alpha) * v[i-1][j]) +
                            (alpha * v[i][j])
                    }
                }
            }
            y := make([]float64, d)
            for i := 0; i < d; i++ {
                y[i] = v[s][i] / v[s][d]
            }
            return y
        }
    }
}
