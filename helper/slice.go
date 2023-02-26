package helper

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func RemoveSliceDuplicate[T constraints.Ordered](input []T) []T {

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

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
