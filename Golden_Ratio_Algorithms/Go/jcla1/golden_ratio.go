package phi

import "math"

func RootIteration(phi, epsilon float64) float64 {
	tmp := 0.0

	for math.Abs(tmp-phi) > epsilon {
		tmp = phi
		phi = math.Sqrt(1 + phi)
	}

	return phi
}
