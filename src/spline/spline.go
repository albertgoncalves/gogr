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

// func sameLens(d int, xs [][]float64) bool {
//     for _, x := range xs {
//         if len(x) != d {
//             return false
//         }
//     }
//     return true
// }

func floatRange(min, max int) []float64 {
    a := make([]float64, max-min)
    for i := range a {
        a[i] = float64(min + i)
    }
    return a
}

// func copy2d(xs [][]float64) [][]float64 {
//     ys := make([][]float64, len(xs))
//     for i := range xs {
//         ys[i] = make([]float64, len(xs[i]))
//         copy(ys[i], xs[i])
//     }
//     return ys
// }

// func interpolate(
//     degree int,
//     domain,
//     d *int,
//     low,
//     high *float64,
//     t float64,
//     knots []float64,
//     vs [][]float64,
// ) []float64 {
// }

func Spline(points []float64, n, m int, ts []float64) []float64 {
    const degree = 2
    if ((n * m) != len(points)) || (n <= degree) {
        return []float64{}
    }
    knots := floatRange(0, n+degree+1)
    domain := len(knots) - 1 - degree
    low := knots[degree]
    high := knots[domain]
    vs := make([]float64, n * (m + 1))
    for i := 0; i < n; i++ {
        for j := 0; j < (m + 1); j++ {
            if j < m {
                vs[(i*(m + 1))+j] = points[(i*m)+j]
            } else {
                vs[(i*(m + 1))+j] = 1
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
                    xs[(i*(m + 1))+j] = ((1 - alpha) * xs[((i-1)*(m + 1))+j]) +
                        (alpha * xs[(i*(m + 1))+j])
                }
            }
        }
        for j := 0; j < m; j++ {
            ys[(k*m)+j] = xs[(s*(m + 1))+j] / xs[(s*(m + 1))+m]
        }
    }
    return ys
}
