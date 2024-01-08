package utils

func Map[T, U any](arr []T, f func(T) U) []U {
	modifiedArr := make([]U, len(arr))

	for i := range arr {
		modifiedArr[i] = f(arr[i])
	}
	return modifiedArr
}
