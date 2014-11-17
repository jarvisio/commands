package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type HTTPFetcher interface {
	Fetch(url string) ([]byte, error)
}

type Fetcher struct{}

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

type StubFetcher struct{}

func (fetcher *StubFetcher) Fetch(url string) ([]byte, error) {
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
