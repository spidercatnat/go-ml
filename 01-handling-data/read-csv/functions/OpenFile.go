package functions

import "os"

// OpenFile ... Really, just open any file from a given path
func OpenFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	Check(err)
	return f
}
