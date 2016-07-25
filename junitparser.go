package parse

import (
	"encoding/xml"
)

type JunitTestSuite struct {
	Name      string          `xml:"name,attr"`
	Time      string          `xml:"time,attr"`
	TestCases []JunitTestCase `xml:"testcase"`
}

type JunitTestCase struct {
	Name    string `xml:"name,attr"`
	Time    string `xml:"time,attr"`
	Failure string `xml:"failure"`
}

func toJunitTestSuite(junit []byte) (*JunitTestSuite, error) {
	var suite JunitTestSuite
	err := xml.Unmarshal(junit, &suite)
	if err != nil {
		return nil, err
	}

	return &suite, nil
}
