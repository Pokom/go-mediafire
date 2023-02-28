// go_mediafire is meant to be a utility to download files from Mediafire. It is directly
// inspired by https://github.com/juvenal-yescas/mediafire-dl and meant to be a golang replacement.
// The package should be available as a CLI utility, or by importing go_mediafire and using programmatically.
package go_mediafire

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var matchDownloadURL = regexp.MustCompile("href=\"((http|https)://download[^\"]+)")
var matchFilename = regexp.MustCompile(`filename=\"(.*)\"`)

func findMatch(r io.Reader, matcher *regexp.Regexp) (string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		match := matcher.MatchString(line)
		if match {
			return matcher.FindString(line), nil
		}
	}
	return "", fmt.Errorf("download link not Found")
}

func findFileName(content string) string {
	matches := matchFilename.FindStringSubmatch(content)
	return matches[1]
}

func Download(url string, w io.Writer) error {
	// Open up connection to url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check header for 'Content-Disposition'
	contentDisposition := resp.Header.Get("Content-Disposition")
	var downloadURL string
	if contentDisposition != "" {
		downloadURL = url
	} else if contentDisposition == "" {
		// If contentDisposition is empty, update url by calling extrctDownloadLink
		downloadURL, err = findMatch(resp.Body, matchDownloadURL)
		if err != nil {
			return err
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("%s appears to be password protected", url)
	}

	resp, err = http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := findFileName(resp.Header.Get("Content-Disposition"))
	if fileName == "" {
		return fmt.Errorf("could not find filename: %s", err.Error())
	}

	_, err = io.Copy(w, resp.Body)
	return err
}
