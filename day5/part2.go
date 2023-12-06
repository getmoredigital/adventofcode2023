package day5

import (
	"errors"
)

func createRanges(rawIds []int) ([]Span, error) {
	if len(rawIds)%2 != 0 {
		return nil, errors.New("the length of the slice is not even")
	}

	var all []Span

	for i := 0; i < len(rawIds); i += 2 {
		end := rawIds[i] + rawIds[i+1]
		all = append(all, Span{Start: rawIds[i], End: end})
	}

	return all, nil
}