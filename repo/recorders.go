package repo

import "fmt"

type FileRecorder struct{}

func (FileRecorder) Record(speech string) {
	// Write our speech to a file.
	// Assume it worked.
	fmt.Println("Recorded speech to file:", speech)
}
