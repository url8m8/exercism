// Package raindrops checks if raindrops is a factor to certain numbers, and assigns a sound based off of those factors.
package raindrops

import "strconv"

// Convert receives the number of raindrops, checks to see if it's a factor of 3, 5, or 7.  Then returns the corresponding sound.
func Convert(rain int) string {
	var sound string
	if (rain % 3) == 0 {
		sound += "Pling"
	}
	if (rain % 5) == 0 {
		sound += "Plang"
	}
	if (rain % 7) == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = strconv.Itoa(rain)
	}
	return sound
}
