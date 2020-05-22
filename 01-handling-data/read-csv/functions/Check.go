package functions

import "log"

// Check ... Convenience function
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
