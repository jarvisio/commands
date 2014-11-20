package commands

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestQueryWolframAlpha(t *testing.T) {
	const expected = "392448 km"
	stub := new(stubFetcher)
	question := "some query"
	answer, err := Ask(question, stub)
	if err != nil {
		t.Errorf("%v", err)
	}
	if x := string(answer); x != expected {
		t.Errorf("Ask(\"%v\") = %v, want %v", question, x, expected)
	}
}

func TestWolframAlphaQuery(t *testing.T) {
	const expected = "392448 km"
	stub := new(stubFetcher)
	output, err := wolframAlphaQuery(stub, "some query")
	if err != nil {
		t.Errorf("%v", err)
	}
	actual := string(output)
	if actual != expected {
		t.Errorf("WolframAlphaQuery() = %v, want %v", actual, expected)
	}
}

func TestWolframAlphaQueryNoResult(t *testing.T) {
	const expected = "Nothing found"
	stub := new(stubFetcher)
	result, err := wolframAlphaQuery(stub, "produce error")
	if result != nil {
		t.Errorf("Result should be nil but is %v", result)
	}
	if err.Error() != expected {
		t.Errorf("WolframAlphaQuery() = %v, want %v!", err, expected)
	}
}

func TestWolframAlphaUrl(t *testing.T) {
	const input = "Distance Earth"
	const expected = "http://api.wolframalpha.com/v2/query?input=Distance+Earth&appid=KHJ7LL-XU5JJR9HV9"
	if x := wolframAlphaURL(input); x != expected {
		t.Errorf("WolframAlphaUrl(%v) = %v, want %v", input, x, expected)
	}
}

func TestReadPods(t *testing.T) {
	xmlFilePath, err := filepath.Abs("fixtures/wolfram_alpha.xml")
	if err != nil {
		t.Errorf("%v", err)
	}
	file, err := ioutil.ReadFile(xmlFilePath)
	if err != nil {
		t.Errorf("%v", err)
	}

	resultBytes, err := readPods(file)
	result := string(resultBytes)
	if err != nil {
		t.Errorf("%v", err)
	}
	const expected = "392448 km  (kilometers)"
	if result != expected {
		t.Errorf("ReadPods() = %v, want %v", result, expected)
	}
}
