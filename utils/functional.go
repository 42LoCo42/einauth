package utils

func Map[X, Y any](xs []X, f func(x X) Y) []Y {
	ys := make([]Y, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func Filter[X any](xs []X, f func(x X) bool) []X {
	results := []X{}
	for _, x := range xs {
		if f(x) {
			results = append(results, x)
		}
	}
	return results
}

func Any[X any](xs []X, f func(x X) bool) bool {
	return len(Filter(xs, f)) > 0
}

func EmptyOrAny[X any](xs []X, f func(x X) bool) bool {
	return len(xs) == 0 || Any(xs, f)
}

func And(bs []bool) bool {
	for _, b := range bs {
		if !b {
			return b
		}
	}
	return true
}

func Or(bs []bool) bool {
	for _, b := range bs {
		if b {
			return b
		}
	}
	return false
}
