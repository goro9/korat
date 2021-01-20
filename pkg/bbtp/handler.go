package bbtp

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

type Req struct {
	Method byte   `cbor:"method"`
	Auth   []byte `cbor:"auth,omitempty"`
	Path   []byte `cbor:"path,omitempty"`
	Body   []byte `cbor:"body"`
}

type Res struct {
	Method byte   `cbor:"method"`
	Body   []byte `cbor:"body"`
}

type reqSampleBody struct {
	TestStr string `cbor:"str"`
	TestNum int    `cbor:"num"`
}

type resSampleBody struct {
	TestStr string `cbor:"str"`
	TestNum int    `cbor:"num"`
}

func Handle(reqBin []byte) ([]byte, error) {
	// unmarshal request
	var req Req
	err := cbor.Unmarshal(reqBin, &req)
	if err != nil {
		return nil, err
	}

	// unmarshal request body
	var reqBody reqSampleBody
	err = cbor.Unmarshal(req.Body, &reqBody)
	fmt.Printf("request: str=%s, num=%d\n", reqBody.TestStr, reqBody.TestNum)

	// marshal response body
	resBody := resSampleBody{
		TestStr: reqBody.TestStr,
		TestNum: reqBody.TestNum,
	}
	res := Res{
		Method: req.Method,
	}
	fmt.Printf("response: str=%s, num=%d\n", resBody.TestStr, resBody.TestNum)
	res.Body, err = cbor.Marshal(&resBody)

	// marshal response
	resBin, err := cbor.Marshal(&res)

	return resBin, nil
}
