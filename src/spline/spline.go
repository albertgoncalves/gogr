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

func floatRange(a, b int) []float64 {
    xs := make([]float64, b-a)
    for i := range xs {
        xs[i] = float64(a + i)
    }
    return xs
}

func Spline(points []float64, n, m int, ts []float64) []float64 {
    const degree = 2
    if ((n * m) != len(points)) || (n <= degree) {
        return []float64{}
    }
    knots := floatRange(0, n+degree+1)
    domain := len(knots) - 1 - degree
    low := knots[degree]
    high := knots[domain]
    vs := make([]float64, n*(m+1))
    for i := 0; i < n; i++ {
        for j := 0; j < (m + 1); j++ {
            if j < m {
                vs[(i*(m+1))+j] = points[(i*m)+j]
            } else {
                vs[(i*(m+1))+j] = 1
            }
        }
    }
    ys := make([]float64, len(ts)*m)
    for k, t := range ts {
        t = (t * (high - low)) + low
        if (t < low) || (t > high) {
            continue
        }
        s := 0
        for i := degree; i < domain; i++ {
            if (t >= knots[i]) && (t <= knots[i+1]) {
                s = i
                break
            }
        }
        xs := make([]float64, len(vs))
        copy(xs, vs)
        for l := 1; l < (degree + 1); l++ {
            for i := s; i > (s - degree - 1 + l); i-- {
                alpha := (t - knots[i]) / (knots[i+degree+1-l] - knots[i])
                for j := 0; j < (m + 1); j++ {
                    xs[(i*(m+1))+j] = ((1 - alpha) * xs[((i-1)*(m+1))+j]) +
                        (alpha * xs[(i*(m+1))+j])
                }
            }
        }
        for j := 0; j < m; j++ {
            ys[(k*m)+j] = xs[(s*(m+1))+j] / xs[(s*(m+1))+m]
        }
    }
    return ys
}
