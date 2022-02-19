package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	res := rr.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Want %d; got %d", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "OK" {
		t.Errorf("Want %s; got %s", "OK", body)
	}
}

func TestPingE2E (t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _ , body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body equal to %q", "OK")
	}
}