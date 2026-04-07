package renderer

import "strings"

func Render(input string, bannerMap map[rune][]string) string {

	splitedForNewLine := strings.Split(input, "\\n")

	//checking if the last slice is empty
	if len(splitedForNewLine) > 0 && splitedForNewLine[len(splitedForNewLine)-1] == "" {
		//remove trailing space
		splitedForNewLine = splitedForNewLine[:len(splitedForNewLine)-1]
	}
	result := ""
	//print only the first segment if a \n is prescent in the input string
	for _, segment := range splitedForNewLine {
		//checking for a segment that hold an empty string
		if segment == "" {
			result += "\n"
		} else {

			//the loop loops through the slice of string
			for row := 0; row <= 7; row++ {
				line := ""

				for _, ch := range segment {
					//getting the first to the eigh row of the whole letters in the segment
					line += bannerMap[ch][row]
				}
				//adding "\n" to the accumulatd row, to continue the printing on the next row till the next row
				result += line + "\n"
			}

		}

	}

	return result
}
