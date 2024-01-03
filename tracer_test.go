package go_trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Errorf("returned value is nil")
	}
	tracer.Trace("hello")
	if buf.String() != "hello" {
		t.Errorf("output the wrong string: '%s'", buf.String())
	}
}
