package parse

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAnalyticsTests(t *testing.T) {
	log.Info("TestToAnalyticsTests")
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

func getFailureAnalyticsTests(tests []AnalyticsTest) []AnalyticsTest {
	var ret []AnalyticsTest
	for _, currTest := range tests {
		if currTest.Failure != "" {
			ret = append(ret, currTest)
		}
	}

	return ret
}

func getPassedAnalyticsTests(tests []AnalyticsTest) []AnalyticsTest {
	var ret []AnalyticsTest
	for _, currTest := range tests {
		if currTest.Failure == "" {
			ret = append(ret, currTest)
		}
	}

	return ret
}
