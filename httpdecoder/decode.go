package httpdecoder

import (
	"encoding/json"
	"errors"
	"github.com/tinylib/msgp/msgp"
	"net/http"
	"strings"
)

const (
	contentType = "Content-Type"
)

func DecodeRequest(r *http.Request, v interface{}) {
	defer r.Body.Close()
	split := strings.SplitN(r.Header.Get(contentType), ";", 1)
	if len(split) < 1 {
		panic(errors.New("unsupport content-type"))
	}

	switch split[0] {
	case "application/octet-stream":
		if x, ok := v.(msgp.Decodable); ok {
			panic(msgp.Decode(r.Body, x))
		}

		panic(errors.New("unable to decode msgpack"))
	case "application/json":
		json.NewDecoder(r.Body).Decode(v)
	}

	panic(errors.New("unsupport content-type"))
}

//ParseJSON decode json to interface{}
func DecodeJSON(r *http.Request, v interface{}) {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		panic(err)
	}
}

//ParseMSGPack decode msgpack to interface{}
func DecodeMSGPack(r *http.Request, v msgp.Decodable) {
	defer r.Body.Close()
	if err := msgp.Decode(r.Body, v); err != nil {
		panic(err)
	}
}
