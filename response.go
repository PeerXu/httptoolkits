package httptoolkits

import (
	"encoding/json"
	"net/http"

	"github.com/stretchr/objx"
)

type JSONResponseWriter interface {
	http.ResponseWriter
	WriteJSON(interface{}) error
}

type jsonResponseWriter struct {
	http.ResponseWriter

	contentType string
}

func (w *jsonResponseWriter) WriteJSON(i interface{}) error {
	w.Header().Set("Content-Type", w.contentType)
	return json.NewEncoder(w).Encode(i)
}

func NewWrapJSONResponseWriterOption() objx.Map {
	return objx.New(map[string]interface{}{
		"content-type": "application/json",
	})

}

type WrapJSONResponseWriterOption func(objx.Map)

func SetContentType(ct string) WrapJSONResponseWriterOption {
	return func(o objx.Map) {
		o["content-type"] = ct
	}
}

func WrapJSONResponseWriter(w http.ResponseWriter, opts ...WrapJSONResponseWriterOption) JSONResponseWriter {
	o := NewWrapJSONResponseWriterOption()
	for _, opt := range opts {
		opt(o)
	}

	ct := o.Get("content-type").String()

	return &jsonResponseWriter{
		ResponseWriter: w,
		contentType:    ct,
	}
}
