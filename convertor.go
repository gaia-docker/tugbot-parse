package parse

import (
	log "github.com/Sirupsen/logrus"
	"strconv"
)

type TestSet struct {
	Name  string
	Time  float64
	Tests []Test
}

type Test struct {
	Name    string
	Time    float64
	Failure string
}

func ToTestSet(junit []byte) (*TestSet, error) {
	suite, err := toJunitTestSuite(junit)
	if err != nil {
		log.Error("Failed to parse junit xml", err, junit)
		return nil, err
	}

	var ret TestSet
	ret.Name = suite.Name
	ret.Time = toFloat(suite.Time)
	if len(suite.TestCases) > 0 {
		for _, currTestCase := range suite.TestCases {
			test := Test{
				Name:    currTestCase.Name,
				Time:    toFloat(currTestCase.Time),
				Failure: currTestCase.Failure,
			}
			ret.Tests = append(ret.Tests, test)
		}
	}

	return &ret, nil
}

func toFloat(val string) float64 {
	ret, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Error(err)
		return -1
	}

	return ret
}
