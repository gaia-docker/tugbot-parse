package parse

import (
	"encoding/xml"
)

type JunitTestSuite struct {
	Name      string          `xml:"name,attr"`
	Time      string          `xml:"time,attr"`
	Total     string          `xml:"tests,attr"`
	Failures  string          `xml:"failures,attr"`
	Errors    string          `xml:"errors,attr"`
	Skipped   string          `xml:"skipped,attr"`
	TestCases []JunitTestCase `xml:"testcase"`
}

type JunitTestCase struct {
	Name    string   `xml:"name,attr"`
	Time    string   `xml:"time,attr"`
	Failure string   `xml:"failure"`
	Skipped xml.Name `xml:"skipped"`
	Error   string   `xml:"error"`
}

func toJunitTestSuite(junit []byte) (*JunitTestSuite, error) {
	var suite JunitTestSuite
	err := xml.Unmarshal(junit, &suite)
	if err != nil {
		return nil, err
	}

	return &suite, nil
}
