package slices

func Map[T any, U any](slice []T, fn func(T) U) []U {
	res := make([]U, len(slice))
	for i, s := range slice {
		res[i] = fn(s)
	}

	return res
}

func Splice[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
