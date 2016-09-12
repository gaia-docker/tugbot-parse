package parse

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJunitToTestSet(t *testing.T) {
	log.Info("TestJunitToTestSet")
	testSet, err := ToTestSet([]byte(getMockMochaXunit2()))
	assert.NoError(t, err)
	assert.Equal(t, 0.149, testSet.Time)
	assert.Equal(t, 5, len(testSet.Tests))
	assert.Equal(t, 5, testSet.Total, "Total")
	assert.Equal(t, 1, testSet.Failures, "Failures")
	assert.Equal(t, 1, testSet.Errors, "Errors")
	assert.Equal(t, 0, testSet.Skipped, "Skipped")
	failureTests := getFailureTests(testSet.Tests)
	assert.Equal(t, 1, len(failureTests))
	assert.Equal(t, Failed, failureTests[0].Status)
	assert.Equal(t, "", failureTests[0].Error, "Check Test's Error is empty - Failure should be with data")
	assert.NotEqual(t, "", failureTests[0].Failure, "Check that Test's Failure contains something (the exception)")
}

func TestJunitToTestSet_SkippedTest(t *testing.T) {
	log.Info("TestJunitToTestSet_SkippedTest")
	testSet, err := ToTestSet([]byte(getMockMochaXunitWithSkippedTest()))
	assert.NoError(t, err)
	assert.Equal(t, 4.045, testSet.Time)
	assert.Equal(t, 5, len(testSet.Tests))
	assert.Equal(t, 5, testSet.Total, "Total")
	assert.Equal(t, 4, testSet.Failures, "Failures")
	assert.Equal(t, 4, testSet.Errors, "Errors")
	assert.Equal(t, 1, testSet.Skipped, "Skipped")
	assert.Equal(t, 1, len(getSkippedTests(testSet.Tests)))
}

func TestJunitToTestSet_ErrorTest(t *testing.T) {
	log.Info("TestJunitToTestSet_ErrorTest")
	testSet, err := ToTestSet([]byte(getMockMochaXunitWithErrorTest()))
	assert.NoError(t, err)
	assert.Equal(t, 0.003, testSet.Time)
	assert.Equal(t, 1, len(testSet.Tests))
	assert.Equal(t, 1, testSet.Total, "Total")
	assert.Equal(t, 0, testSet.Failures, "Failures")
	assert.Equal(t, 1, testSet.Errors, "Errors")
	assert.Equal(t, 0, testSet.Skipped, "Skipped") //if property not exist we return 0
	assert.Equal(t, Error, testSet.Tests[0].Status, "Check Test's status")
	assert.NotEqual(t, "", testSet.Tests[0].Error, "Check that Test's Error contains something (the exception)")
	assert.Equal(t, "", testSet.Tests[0].Failure, "Check Test's Failure is empty - Error should be with data")
}

func TestToTestSet_InvalidXML(t *testing.T) {
	log.Info("TestToTestSet_InvalidXML")
	_, err := ToTestSet([]byte("abc"))
	assert.Error(t, err)
}

func getFailureTests(tests []Test) []Test {
	var ret []Test
	for _, currTest := range tests {
		if currTest.Failure != "" {
			ret = append(ret, currTest)
		}
	}

	return ret
}

func getSkippedTests(tests []Test) []Test {
	var ret []Test
	for _, currTest := range tests {
		if currTest.Status == Skipped {
			ret = append(ret, currTest)
		}
	}

	return ret
}
