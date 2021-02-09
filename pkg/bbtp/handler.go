package bbtp

type Method int

const (
	Get Method = iota
	Post
	Put
	Delete
)

type Request struct {
	Method   Method `cbor:"method"`
	UUID     string
	Path     string `cbor:"path"`
	Header   Header `cbor:"header"`
	Body     Body   `cbor:"body"`
	Response *Response
}

type Header struct {
	Auth []byte `cbor:"auth,omitempty"`
}

type Body []byte

type Response struct {
	StatusCode int    `cbor:"code"`
	Header     Header `cbor:"header"`
	Body       Body   `cbor:"body"`
	Request    *Request
}

func Handle(uuid string, reqRaw []byte) ([]byte, error) {
	// unmarshal request
	var req Request
	err := decodeReq(uuid, reqRaw, &req)
	if err != nil {
		return nil, err
	}

	// handle request
	var res Response
	err = hm[req.Path].h(&res, &req)
	if err != nil {
		return nil, err
	}

	// marshal response
	resRaw, err := encodeRes(&res)

	return resRaw, nil
}

type handlerMap map[string]handler

var hm handlerMap

type handler struct {
	path string
	h    func(*Response, *Request) error
}

func HandleFunc(path string, h func(*Response, *Request) error) {
	hBuf := handler{path: path, h: h}
	hm[path] = hBuf
}

func init() {
	hm = handlerMap{}
}
