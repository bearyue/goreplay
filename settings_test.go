package goreplay

import (
	"encoding/json"
	"testing"
)

func TestAppSettings(t *testing.T) {
	a := AppSettings{}
	_, err := json.Marshal(&a)
	if err != nil {
		t.Error(err)
	}
}

func TestMultiOptionOutputHTTP(t *testing.T) {
	var urls []string
	opt := &MultiOption{a: &urls}

	if err := opt.Set("http://172.16.71.232:10801"); err != nil {
		t.Fatal(err)
	}
	if err := opt.Set("http://120.79.175.70:10801"); err != nil {
		t.Fatal(err)
	}

	if len(urls) != 2 {
		t.Fatalf("expected 2 output-http values, got %d", len(urls))
	}
	if urls[0] != "http://172.16.71.232:10801" {
		t.Errorf("unexpected first value: %s", urls[0])
	}
	if urls[1] != "http://120.79.175.70:10801" {
		t.Errorf("unexpected second value: %s", urls[1])
	}
}

func TestMultiOptionOutputHTTPSingle(t *testing.T) {
	var urls []string
	opt := &MultiOption{a: &urls}

	if err := opt.Set("http://staging.example.com:8080"); err != nil {
		t.Fatal(err)
	}

	if len(urls) != 1 {
		t.Fatalf("expected 1 output-http value, got %d", len(urls))
	}
	if urls[0] != "http://staging.example.com:8080" {
		t.Errorf("unexpected value: %s", urls[0])
	}
}
