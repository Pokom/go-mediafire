package main

import (
	"fmt"
	go_mediafire "github.com/pokom/go-mediafire"
	"os"
)

func run(urls []string, outputDir string)  error {
	for _, url := range urls {
		// TODO: Open up a file to write out to
		err := go_mediafire.Download(url, outputDir)
		if err != nil {
			return fmt.Errorf("could not process url(%s): %s", url, err.Error())
		}
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Incorrect usage. Need to pass in list of urls to download.")
		os.Exit(1)
	}

	// TODO: Convert this into an optional flag
	outputDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := run(os.Args[1:], outputDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}