// Package hamming calculates the hamming distance between two DNA strands.
package hamming

import "fmt"

// Distance checks to see if strands are equal, if so, a for loop compares the two and keeps account of difference in index variable.
func Distance(a, b string) (int, error) {
    var index int

    if len(a) != len(b) {
        return 0, fmt.Errorf("Strands must be of equal length.")
    }
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            index += 1
        }
    }
    return index, nil
}
