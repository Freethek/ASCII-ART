package banner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Load(bannerName string) map[rune][]string {
	//reading the banner file and storing it
	fileData, err := os.Open(bannerName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an Error reading file %v", err)
		os.Exit(1)
	}

	//closing the file
	defer fileData.Close()

	//declaring the variable that will load the content of the loaded file
	var lines []string

	//going into the content
	scanner := bufio.NewScanner(fileData)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimRight(line, "\r")

		lines = append(lines, line)
	}

	bannerMap := make(map[rune][]string)

	for i := 32; i <= 126; i++ {
		start := (i-32)*9 + 1
		grab8lines := lines[start : start+8]
		bannerMap[rune(i)] = grab8lines
	}
	return bannerMap
}
