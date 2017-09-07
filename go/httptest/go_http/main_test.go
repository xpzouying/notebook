package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/index", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	indexHandle(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("result of http request is not ok: %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	if string(bytes.TrimSpace(b)) != "hello world" {
		t.Errorf("body of request should be `hello world`, but get %s", b)
	}
}

func TestSvrListen(t *testing.T) {
	r := http.NewServeMux()
	r.HandleFunc("/hello", indexHandle)

	svr := httptest.NewServer(r)
	defer svr.Close()

	resp, err := http.Get(fmt.Sprintf("%s/hello", svr.URL))
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body failed: %v", err)
	}
	if string(bytes.TrimSpace(b)) != "hello world" {
		t.Errorf("result is not equal as expected: %s", b)
	}
}

func TestMuxSets(t *testing.T) {
	svr := httptest.NewServer(newHandle())
	defer svr.Close()

	for _, tc := range handleSets {
		resp, err := http.Get(svr.URL + tc.path)
		if err != nil {
			t.Errorf("http get error: %v", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("http get status should be ok, but is: %d", resp.StatusCode)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("read from response body error: %v", err)
		}

		if string(b) != tc.resp {
			t.Errorf("expected: %s, but response is: %s", tc.resp, b)
		}
	}
}
