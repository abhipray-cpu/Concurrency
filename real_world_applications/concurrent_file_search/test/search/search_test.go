package search

import (
	"concurrent_file_search/pkg/search"
	"io/ioutil"
	"os"
	"testing"
)

// helper function to create a temp file for testing
func createTempFile(content string) (string, func(), error) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return "", nil, err
	}

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
		return "", nil, err
	}

	if err := tmpfile.Close(); err != nil {
		os.Remove(tmpfile.Name())
		return "", nil, err
	}

	// cleanup function to delete the temp file
	cleanup := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile.Name(), cleanup, nil
}

func TestRegexpSearchFile(t *testing.T) {
	// Setup test data
	content := `First line
Second line
Third line with pattern
Fourth line`
	filePath, cleanup, err := createTempFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer cleanup()

	tests := []struct {
		name      string
		pattern   string
		wantMatch bool
		wantErr   bool
	}{
		{"Pattern exists", "Third line with pattern", true, false},
		{"Pattern does not exist", "non-existent pattern", false, false},
		{"Invalid pattern", "[", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches, ok := search.RegexpSearchFile(filePath, tt.pattern)
			if (ok != tt.wantMatch) || (len(matches) > 0 != tt.wantMatch) {
				t.Errorf("RegexpSearchFile() got = %v, want %v", ok, tt.wantMatch)
			}

		})
	}
}
