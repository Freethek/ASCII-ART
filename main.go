package main

import (
	"ASCII-ART/banner"
	"ASCII-ART/renderer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) >= 3 {
		fmt.Fprintf(os.Stderr, "Usage: go run . <input string> or Usage: go run . <Input String> <banner File name>")
		os.Exit(1)
	}

	//if the input is empty
	if os.Args[1] == "" {
		fmt.Println()
		os.Exit(0)
	}

	//declaring and initializing the bannerName to the file name of the banner
	bannerName := "standard"
	var input string

	//checking for input and banner file in the argument
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else if len(os.Args) == 3 {
		input = os.Args[1]
		bannerName = os.Args[2]
	}

	if input == "" {
		fmt.Println()
		os.Exit(0)
	}

	bannerMap := banner.Load("banners/" + bannerName + ".txt")

	result := renderer.Render(input, bannerMap)

	fmt.Print(result)

}
