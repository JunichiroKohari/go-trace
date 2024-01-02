package go_trace

type Tracer interface {
	Trace(...interface{})
}
