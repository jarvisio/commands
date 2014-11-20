package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// HTTPFetcher is an interface for fetching things
// from the so called internet.
type HTTPFetcher interface {
	Fetch(url string) ([]byte, error)
}

// Fetcher is the struct you should use to fetch
// things from the internet. It conforms to HTTPFetcher interface.
type Fetcher struct{}

// Fetch performes the fetch from the given url
// it returns a byte array or an error.
func (fetcher *Fetcher) Fetch(url string) ([]byte, error) {
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// For testing
type stubFetcher struct{}

func (fetcher *stubFetcher) Fetch(url string) ([]byte, error) {
	if strings.Contains(url, "api.wolframalpha.com/v2/query?input=some+query") {
		// Return Wolfram Alpha file
		xmlFilePath, _ := filepath.Abs("fixtures/wolfram_alpha.xml")
		file, _ := ioutil.ReadFile(xmlFilePath)
		return file, nil
	} else if strings.Contains(url, "api.wolframalpha.com/v2/query?input=produce+error") {
		// Return Wolfram Alpha file
		xmlFilePath, _ := filepath.Abs("fixtures/wolfram_alpha_no_result.xml")
		file, _ := ioutil.ReadFile(xmlFilePath)
		return file, nil
	}
	return nil, errors.New("No Internet connection")
}
