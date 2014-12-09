package commands

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/url"
	"strings"
)

// Ask takes a question and an instance of HTTPFetcher
// it returns a byte array or an error.
func Ask(question string, fetcher HTTPFetcher) ([]byte, error) {
	return wolframAlphaQuery(fetcher, question)
}

func wolframAlphaQuery(fetcher HTTPFetcher, query string) ([]byte, error) {
	url := wolframAlphaURL(query)
	response, err := fetcher.Fetch(url)
	if err == nil {
		result, err := readPods(response)
		if err != nil {
			return nil, err
		}
		plainText := string(result)
		parts := strings.SplitAfter(plainText, " (")
		parsedSolution := strings.TrimRight(parts[0], " (")
		return []byte(parsedSolution), nil
	}
	return nil, err
}

func wolframAlphaURL(text string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://api.wolframalpha.com/v2/query?input=")
	buffer.WriteString(url.QueryEscape(text))
	buffer.WriteString("&appid=KHJ7LL-XU5JJR9HV9")
	return buffer.String()
}

func readPods(body []byte) ([]byte, error) {
	var queryResult xmlQueryResult
	if err := xml.Unmarshal(body, &queryResult); err != nil {
		return nil, err
	}
	if queryResult.Success == true {
		plainText := queryResult.Pods[1].SubPods[0].PlainText
		return []byte(plainText), nil
	}
	err := errors.New("Nothing found")
	return nil, err
}

type xmlPod struct {
	XMLName    xml.Name    `xml:"pod"`
	Title      string      `xml:"title,attr"`
	Scanner    string      `xml:"scanner,attr"`
	ID         string      `xml:"id,attr"`
	Position   string      `xml:"position,attr"`
	Error      string      `xml:"error,attr"`
	NumSubPods string      `xml:"numsubpods,attr"`
	SubPods    []xmlSubPod `xml:"subpod"`
}

type xmlSubPod struct {
	XMLName   xml.Name `xml:"subpod"`
	Title     string   `xml:"title,attr"`
	PlainText string   `xml:"plaintext"`
}

type xmlQueryResult struct {
	XMLName xml.Name `xml:"queryresult"`
	Success bool     `xml:"success,attr"`
	Error   bool     `xml:"error,attr"`
	Pods    []xmlPod `xml:"pod"`
}
