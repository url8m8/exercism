// Package hamming calculates the hamming distance between two DNA strands.
package hamming

import "errors"

// Distance checks to see if strands are equal, if so, a for loop compares the two and keeps account of difference in distance variable.
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("strands must be of equal length")
	}
	var distance int
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
