package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"ginApi/example"
	"github.com/golang/protobuf/proto"
	"bytes"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestProto(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	body := &example.Test{Label: proto.String("test"), Type: proto.Int32(22)}
	inData, _ := proto.Marshal(body)

	req, _ := http.NewRequest("POST", "/proto", bytes.NewReader(inData))
	req.Header.Set("Content-Type","application/x-protobuf")
	router.ServeHTTP(w, req)

	var test example.Test
	_ = proto.Unmarshal(w.Body.Bytes(), &test)

	assert.Equal(t, "test123", test.GetLabel())
}
