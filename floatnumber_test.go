package parse

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	log.Info("TestMarshalJSON")
	json, err := FloatNumber(0).MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "0.000000", string(json))
}
