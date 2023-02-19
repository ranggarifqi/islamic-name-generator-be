package helper

func RemoveSliceDuplicate[T comparable](input []T) []T {

	mapping := make(map[T]bool)

	for _, val := range input {
		if !mapping[val] {
			mapping[val] = true
		}
	}

	var result []T

	for key := range mapping {
		result = append(result, key)
	}

	return result
}
