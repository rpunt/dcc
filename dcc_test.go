package dcc

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	// client := NewClient()
}

func TestSendRequest(t *testing.T) {

}

func TestGetRequest(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open("testdata/icanhazdadjoke.json")
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			io.Copy(w, f)
		},
	))
	defer ts.Close()
	c := NewClient()
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	response, err := c.Get("/")
	if err != nil {
		log.Panicln("error:", err)
	}

	got := response
	want := HttpResponse{
		Body: "",
		Code: 200,
	}

	if !cmp.Equal(want.Code, got.Code) {
		t.Error(cmp.Diff(want.Code, got.Code))
	}
}

func TestPostRequest(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open("testdata/icanhazdadjoke.json")
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			io.Copy(w, f)
		},
	))
	defer ts.Close()
	c := NewClient()
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	response, err := c.Post("/")
	if err != nil {
		log.Panicln("error:", err)
	}

	got := response
	want := HttpResponse{
		Body: "",
		Code: 200,
	}

	if !cmp.Equal(want.Code, got.Code) {
		t.Error(cmp.Diff(want.Code, got.Code))
	}
}
