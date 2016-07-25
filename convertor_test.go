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
	assert.NotEqual(t, "", getFailure(testSet.Tests))
}

func getFailure(tests []Test) string {
	ret := ""
	for _, currTest := range tests {
		if currTest.Failure != "" {
			ret = currTest.Failure
			break
		}
	}

	return ret
}
