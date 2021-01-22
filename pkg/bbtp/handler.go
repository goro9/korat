package bbtp

import (
	"github.com/fxamacker/cbor/v2"
)

type Req struct {
	Method byte   `cbor:"method"`
	Auth   []byte `cbor:"auth,omitempty"`
	Path   string `cbor:"path,omitempty"`
	Body   []byte `cbor:"body"`
}

type Res struct {
	Method byte   `cbor:"method"`
	Body   []byte `cbor:"body"`
}

func Handle(reqBin []byte) ([]byte, error) {
	// unmarshal request
	var req Req
	err := cbor.Unmarshal(reqBin, &req)
	if err != nil {
		return nil, err
	}

	// handle request
	res := Res{
		Method: 0,
	}
	res.Body = hm[req.Path].h(req.Body)

	// marshal response
	resBin, err := cbor.Marshal(&res)

	return resBin, nil
}

type handlerMap map[string]handler

var hm handlerMap

type handler struct {
	path string
	h    func([]byte) []byte
}

func HandleFunc(path string, h func([]byte) []byte) {
	hBuf := handler{path: path, h: h}
	hm[path] = hBuf
}

func init() {
	hm = handlerMap{}
}
