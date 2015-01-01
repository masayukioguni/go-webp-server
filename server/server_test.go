package server

import (
	"github.com/drone/routes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type MockServer struct {
	Mux    *routes.RouteMux
	Server *Server
}

func NewMockServer() *MockServer {
	m := &MockServer{}
	m.Mux = routes.New()

	m.Server = NewServer("http://localhost/")
	m.Server.Route(m.Mux)

	return m
}

func TestServer_IsValidImagePath(t *testing.T) {
	mockServer := NewMockServer()

	res := mockServer.Server.IsValidImagePath("")
	wantRes := false

	if !reflect.DeepEqual(res, wantRes) {
		t.Errorf("TestServer_IsValidExt returned %+v, want %+v", res, wantRes)
	}

	res = mockServer.Server.IsValidImagePath("a.gif")
	wantRes = false

	if !reflect.DeepEqual(res, wantRes) {
		t.Errorf("TestServer_IsValidExt returned %+v, want %+v", res, wantRes)
	}

	res = mockServer.Server.IsValidImagePath("a.jpg")
	wantRes = true

	if !reflect.DeepEqual(res, wantRes) {
		t.Errorf("TestServer_IsValidExt returned %+v, want %+v", res, wantRes)
	}

}

func TestWebpHandler_NotFoundImagePath(t *testing.T) {
	mockServer := NewMockServer()

	ts := httptest.NewServer(mockServer.Mux)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/webp/")

	wantStatusCode := http.StatusBadRequest

	if !reflect.DeepEqual(res.StatusCode, wantStatusCode) {
		t.Errorf("TestWebpHandler_NotFoundImagePath Response Code returned %+v, want %+v", res.StatusCode, wantStatusCode)
	}

}

func TestWebpHandler_InvalidImage(t *testing.T) {
	mockServer := NewMockServer()

	ts := httptest.NewServer(mockServer.Mux)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/webp/a.jpg")

	wantStatusCode := http.StatusInternalServerError

	if !reflect.DeepEqual(res.StatusCode, wantStatusCode) {
		t.Errorf("TestWebpHandler_InvalidImage Response Code returned %+v, want %+v", res.StatusCode, wantStatusCode)
	}

}

func TestWebpHandler_InvalidWidth(t *testing.T) {
	mockServer := NewMockServer()

	ts := httptest.NewServer(mockServer.Mux)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/700/0/webp/a.jpg")

	wantStatusCode := http.StatusBadRequest

	if !reflect.DeepEqual(res.StatusCode, wantStatusCode) {
		t.Errorf("TestWebpHandler_InvalidWidth Response Code returned %+v, want %+v", res.StatusCode, wantStatusCode)
	}

}

func TestWebpHandler_InvalidHeight(t *testing.T) {
	mockServer := NewMockServer()

	ts := httptest.NewServer(mockServer.Mux)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/0/500/webp/a.jpg")

	wantStatusCode := http.StatusBadRequest

	if !reflect.DeepEqual(res.StatusCode, wantStatusCode) {
		t.Errorf("TestWebpHandler_InvalidWidth Response Code returned %+v, want %+v", res.StatusCode, wantStatusCode)
	}

}
