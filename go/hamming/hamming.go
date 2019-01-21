package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	length_a := len(a)
	length_b := len(b)

	if length_a != length_b {
		return 0, errors.New("Different sizes for Strings")
	}

	countHamming := 0

	for i := 0; i < length_a; i++ {
		if a[i] != b[i] {
			countHamming++
		}
	}

	return countHamming, nil
}
