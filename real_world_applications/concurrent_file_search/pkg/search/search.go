package search

import (
	"bufio"
	"concurrent_file_search/pkg/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// SearchFile searches for a given pattern within a file.
// It returns a slice of strings containing the file path if the pattern is found,
// and a boolean indicating whether any matches were found.
func SearchFile(filePath, searchPattern string) ([]string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		utils.LogMessage(utils.ERROR, fmt.Sprintf("Error opening file %s: %v", filePath, err))
		return nil, false
	}

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchPattern) {
			found = true
			break // found the pattern, no need to continue scanning
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file: ", err)
		utils.LogMessage(utils.ERROR, fmt.Sprintf("Error scanning file %s: %v", filePath, err))
		return nil, false
	}
	if found {
		return []string{filePath}, true
	} else {
		return []string{}, false
	}
}

// RegexpSearchFile searches for a given regular expression pattern in a file.
// It returns matches with context: the matching line along with the line number.
func RegexpSearchFile(filePath, pattern string) ([]string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		utils.LogMessage(utils.ERROR, fmt.Sprintf("Error opening file %s: %v", filePath, err))
		return []string{}, false
	}
	defer file.Close()

	var matches []string
	scanner := bufio.NewScanner(file)
	lineNum := 0
	regex, err := regexp.Compile(pattern)
	if err != nil {
		utils.LogMessage(utils.ERROR, fmt.Sprintf("Error compiling regular expression pattern %s: %v", pattern, err))
		return []string{}, false // Invalid regular expression pattern
	}

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if regex.MatchString(line) {
			match := "Line " + string(lineNum) + ": " + line
			matches = append(matches, match)
		}
	}

	return matches, len(matches) > 0
}
