package commands

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/url"
	"strings"
)

func Ask(question string, fetcher HTTPFetcher) ([]byte, error) {
	return WolframAlphaQuery(fetcher, question)
}

func WolframAlphaQuery(fetcher HTTPFetcher, query string) ([]byte, error) {
	url := WolframAlphaUrl(query)
	response, err := fetcher.Fetch(url)
	if err == nil {
		result, err := ReadPods(response)
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

func WolframAlphaUrl(text string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://api.wolframalpha.com/v2/query?input=")
	buffer.WriteString(url.QueryEscape(text))
	buffer.WriteString("&appid=KHJ7LL-XU5JJR9HV9")
	return buffer.String()
}

func ReadPods(body []byte) ([]byte, error) {
	var queryResult XMLQueryResult
	if err := xml.Unmarshal(body, &queryResult); err != nil {
		return nil, err
	}
	if queryResult.Success == true {
		plainText := queryResult.Pods[1].SubPods[0].PlainText
		return []byte(plainText), nil
	} else {
		err := errors.New("Nothing found")
		return nil, err
	}
}

type XMLPod struct {
	XMLName    xml.Name    `xml:"pod"`
	Title      string      `xml:"title,attr"`
	Scanner    string      `xml:"scanner,attr"`
	ID         string      `xml:"id,attr"`
	Position   string      `xml:"position,attr"`
	Error      string      `xml:"error,attr"`
	NumSubPods string      `xml:"numsubpods,attr"`
	SubPods    []XMLSubPod `xml:"subpod"`
}

type XMLSubPod struct {
	XMLName   xml.Name `xml:"subpod"`
	Title     string   `xml:"title,attr"`
	PlainText string   `xml:"plaintext"`
}

type XMLQueryResult struct {
	XMLName xml.Name `xml:"queryresult"`
	Success bool     `xml:"success,attr"`
	Error   bool     `xml:"error,attr"`
	Pods    []XMLPod `xml:"pod"`
}
