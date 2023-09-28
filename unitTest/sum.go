package uTest

func Addition[T int | float64 | string](val1 T, val2 T) T {
	return val1 + val2
}
