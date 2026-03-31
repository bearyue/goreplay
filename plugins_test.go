package goreplay

import (
	"testing"
)

func TestPluginsRegistration(t *testing.T) {
	Settings.InputDummy = []string{"[]"}
	Settings.OutputDummy = []string{"[]"}
	Settings.OutputHTTP = []string{"www.example.com|10"}
	Settings.InputFile = []string{"/dev/null"}

	plugins := NewPlugins()

	if len(plugins.Inputs) != 3 {
		t.Errorf("Should be 3 inputs got %d", len(plugins.Inputs))
	}

	if _, ok := plugins.Inputs[0].(*DummyInput); !ok {
		t.Errorf("First input should be DummyInput")
	}

	if _, ok := plugins.Inputs[1].(*FileInput); !ok {
		t.Errorf("Second input should be FileInput")
	}

	if len(plugins.Outputs) != 2 {
		t.Errorf("Should be 2 output %d", len(plugins.Outputs))
	}

	if _, ok := plugins.Outputs[0].(*DummyOutput); !ok {
		t.Errorf("First output should be DummyOutput")
	}

	if l, ok := plugins.Outputs[1].(*Limiter); ok {
		if _, ok := l.plugin.(*HTTPOutput); !ok {
			t.Errorf("HTTPOutput should be wrapped in limiter")
		}
	} else {
		t.Errorf("Second output should be Limiter")
	}

}

// TestMultipleHTTPOutputsRegistration verifies that repeating --output-http
// results in one HTTPOutput instance per configured address.
func TestMultipleHTTPOutputsRegistration(t *testing.T) {
	Settings.InputDummy = nil
	Settings.OutputDummy = nil
	Settings.InputFile = nil
	Settings.OutputHTTP = []string{
		"http://192.168.1.1:8080",
		"http://192.168.1.2:8080",
	}

	plugins := NewPlugins()

	if len(plugins.Outputs) != 2 {
		t.Fatalf("expected 2 HTTP outputs, got %d", len(plugins.Outputs))
	}

	for i, out := range plugins.Outputs {
		if _, ok := out.(*HTTPOutput); !ok {
			t.Errorf("output[%d] should be *HTTPOutput, got %T", i, out)
		}
	}

	Settings.OutputHTTP = nil
}
