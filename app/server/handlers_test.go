package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/openmind13/link-shortener/app/config"
	"github.com/stretchr/testify/assert"
)

var (
	testConfig = &config.Config{
		BindAddr:          "localhost:8000",
		ShortURLLength:    7,
		MongodbConnection: "mongodb://localhost:27017",
		DBName:            "linkshortener_test",
		CollectionName:    "links_test",
	}
	configPath = "../../config/server.toml"
)

func Test_handleCreateRandomURL(t *testing.T) {
	testServer, err := New(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"longurl": "https://youtube.com",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "bad request",
			payload: map[string]string{
				"bad data": "https://youtube.com",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "incorrect url",
			payload: map[string]string{
				"longurl": "/$alkjdf$$$",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "incorrect data",
			payload: map[string]string{
				"longurl": "$--/\alkjdf$$$",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "additional payload",
			payload: map[string]string{
				"longurl":  "https://youtube.com",
				"shorturl": "yt",
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/create", b)
			testServer.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func Test_handleCreateCustomURL(t *testing.T) {
	testServer, err := New(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"longurl":  "https://youtube.com",
				"shorturl": "yt",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "doesn't have shorturl parameter",
			payload: map[string]string{
				"longurl": "https://youtube.com",
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/createcustom", b)
			testServer.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func Test_handleShortURL(t *testing.T) {
	testConfig := &config.Config{}
	if _, err := toml.DecodeFile(configPath, testConfig); err != nil {
		t.Fatal(err)
	}
	testServer, err := New(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(map[string]string{
		"longurl":  "https://youtube.com",
		"shorturl": "yt",
	})
	req, _ := http.NewRequest(http.MethodPost, "/createcustom", b)
	testServer.ServeHTTP(rec, req)
	// store value in db
	// test handleShortURL
}
