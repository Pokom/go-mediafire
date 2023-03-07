package main

import (
	"fmt"
	"os"
	"strings"

	go_mediafire "github.com/pokom/go-mediafire"
)

func createFile(fileName string, outputDir string) (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("%s/%s.rar", outputDir, fileName))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getFileName(url string) string {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}

func run(urls []string, outputDir string) error {
	for _, url := range urls {
		fileName := getFileName(url)
		file, err := createFile(fileName, outputDir)
		if err != nil {
			return fmt.Errorf("could not create file: %s", err.Error())
		}
		// TODO: Open up a file to write out to
		_, err = go_mediafire.Download(url, file)
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
