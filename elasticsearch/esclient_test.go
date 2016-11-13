package elasticsearch

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex_CreateIndexAndDocument(t *testing.T) {
	const index = "jenkins"
	const typ = "build"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{
		   "_index": "%s",
		   "_type": "%s",
		   "_id": "AVhdyGBQTm1o5HUvuwJa",
		   "_version": 1,
		   "created": true
		}`, index, typ)))
	}))
	defer server.Close()

	client, err := NewESClient(server.URL)
	assert.NoError(t, err)
	err = client.Initialize(index, typ)
	assert.NoError(t, err)
	err = client.Index(`{"id":"1", "status": "running"}`)
	assert.NoError(t, err)
}

func TestIndexInitialize_IllegalIndexName(t *testing.T) {
	client, err := NewESClient("stam.com")
	assert.NoError(t, err)
	err = client.Initialize("my jenkins", "build")
	assert.Error(t, err)
}

func TestIndexInitialize_IllegalType(t *testing.T) {
	client, err := NewESClient("stam.com")
	assert.NoError(t, err)
	err = client.Initialize("jenkins", "my[build")
	assert.Error(t, err)
}