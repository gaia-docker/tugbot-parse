package parse

import (
	log "github.com/Sirupsen/logrus"
	"strconv"
)

// Test Status
const (
	Passed  = "Passed"
	Skipped = "Skipped"
	Failed  = "Failed"
	Error = "Error"
)

type TestSet struct {
	Name     string
	Time     float64
	Total    int
	Failures int
	Errors   int
	Skipped  int
	Tests    []Test
}

type Test struct {
	Name    string
	Status  string
	Time    float64
	Failure string
	Error   string
}

func ToTestSet(junit []byte) (*TestSet, error) {
	suite, err := toJunitTestSuite(junit)
	if err != nil {
		log.Error("Failed to parse junit xml", err, junit)
		return nil, err
	}

	var ret TestSet
	ret.Name = suite.Name
	ret.Total = toInt(suite.Total)
	ret.Failures = toInt(suite.Failures)
	ret.Errors = toInt(suite.Errors)
	ret.Skipped = toInt(suite.Skipped)
	ret.Time = toFloat(suite.Time)
	if len(suite.TestCases) > 0 {
		for _, currTestCase := range suite.TestCases {
			test := Test{
				Name:    currTestCase.Name,
				Time:    toFloat(currTestCase.Time),
				Failure: currTestCase.Failure,
				Error: currTestCase.Error,
			}
			test.Status = getStatus(currTestCase)
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

func toInt(val string) int {
	//property not exist - we return 0
	if val == "" {
		return 0
	}
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Error(err)
		return -1
	}

	return ret
}

func getStatus(test JunitTestCase) string {
	ret := Passed
	if test.Skipped.Local != "" {
		log.Info("local:", test.Skipped.Local)
		ret = Skipped
	} else if test.Failure != "" {
		ret = Failed
	} else if test.Error != "" {
		ret = Error
	}

	return ret
}
