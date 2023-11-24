package file

import (
	"os"
	"testing"
)

func Test_Validate(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.txt")
	if err != nil {
		panic(err) // This should never happen
	}

	// Once all the tests are done. Delete tmp file
	defer os.Remove(tmpFile.Name())

	tests := []struct {
		name     string
		filePath string
		wErr     bool
	}{
		{"File exist", tmpFile.Name(), false},
		{"File not exist", "notadir/file.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hErr := validate(tt.filePath)
			if (hErr != nil) != tt.wErr {
				t.Errorf("Validate():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}
		})
	}

}
