// Implements the ShareWith function, returning a string with the name variable.
package twofer

import "fmt"

// Takes a name string, returning a message.  If name string is empty, returns generic message.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
}

