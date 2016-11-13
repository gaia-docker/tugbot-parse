package parse

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAnalyticsTestsFailedTest(t *testing.T) {
	log.Info("TestToAnalyticsTestsFailedTest")
	analyticsTests, err := ToAnalyticsTests([]byte(getMockMochaXunit2()))
	assert.NoError(t, err)
	assert.Equal(t, 5, len(analyticsTests))
	analyticsTestSet := analyticsTests[0].TestSet
	assert.Equal(t, FloatNumber(0.149), analyticsTestSet.Time)
	assert.Equal(t, 5, analyticsTestSet.Total, "Total")
	assert.Equal(t, 1, analyticsTestSet.Failures, "Failures")
	assert.Equal(t, 1, analyticsTestSet.Errors, "Errors")
	assert.Equal(t, 0, analyticsTestSet.Skipped, "Skipped")
	failureAnalyticsTests := getFailureAnalyticsTests(analyticsTests)
	assert.Equal(t, 1, len(failureAnalyticsTests))
	assert.Equal(t, 1, failureAnalyticsTests[0].NumericStatus)
	assert.Equal(t, 0, getPassedAnalyticsTests(analyticsTests)[0].NumericStatus)
}

func TestToAnalyticsTestsErrorTest(t *testing.T) {
	log.Info("TestToAnalyticsTestsErrorTest")
	analyticsTests, err := ToAnalyticsTests([]byte(getMockMochaXunitWithErrorTest()))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(analyticsTests))
	analyticsTestSet := analyticsTests[0].TestSet
	assert.Equal(t, FloatNumber(0.003), analyticsTestSet.Time)
	assert.Equal(t, 1, analyticsTestSet.Total, "Total")
	assert.Equal(t, 0, analyticsTestSet.Failures, "Failures")
	assert.Equal(t, 1, analyticsTestSet.Errors, "Errors")
	assert.Equal(t, 0, analyticsTestSet.Skipped, "Skipped")
	errorAnalyticsTests := getErrorAnalyticsTests(analyticsTests)
	assert.Equal(t, 1, len(errorAnalyticsTests))
	assert.Equal(t, 1, errorAnalyticsTests[0].NumericStatus)
	assert.Equal(t, 0, len(getPassedAnalyticsTests(analyticsTests)))
}

func getFailureAnalyticsTests(tests []AnalyticsTest) []AnalyticsTest {
	var ret []AnalyticsTest
	for _, currTest := range tests {
		if currTest.Failure != "" && currTest.Status == Failed {
			ret = append(ret, currTest)
		}
	}

	return ret
}

func getErrorAnalyticsTests(tests []AnalyticsTest) []AnalyticsTest {
	var ret []AnalyticsTest
	for _, currTest := range tests {
		if currTest.Error != "" && currTest.Status == Error {
			ret = append(ret, currTest)
		}
	}

	return ret
}

func getPassedAnalyticsTests(tests []AnalyticsTest) []AnalyticsTest {
	var ret []AnalyticsTest
	for _, currTest := range tests {
		if currTest.Status == Passed {
			ret = append(ret, currTest)
		}
	}

	return ret
}
