package bbtp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fxamacker/cbor/v2"
)

type reqSampleBody struct {
	TestStr string `cbor:"str"`
	TestNum int    `cbor:"num"`
}

type resSampleBody struct {
	TestStr string `cbor:"str"`
	TestNum int    `cbor:"num"`
}

func TestHandle(t *testing.T) {
	// expect data
	expectResBody := resSampleBody{
		TestStr: "hello world",
		TestNum: 255,
	}
	expectResBodyBin, _ := cbor.Marshal(&expectResBody)
	expectRes := Response{
		StatusCode: 200,
		Body:       expectResBodyBin,
	}

	// test data
	reqBody := reqSampleBody{
		TestStr: expectResBody.TestStr,
		TestNum: expectResBody.TestNum,
	}
	reqBodyBin, _ := cbor.Marshal(&reqBody)
	req := Request{
		Method: Post,
		Path:   "test",
		Header: Header{
			Auth: nil,
		},
		Body: reqBodyBin,
	}
	reqBin, _ := cbor.Marshal(&req)
	uuid := "00000000-0000-0000-0000-000000000000"

	// execute
	HandleFunc("test", testHandler)
	resBin, err := Handle(uuid, reqBin)
	if err != nil {
		t.Fatal(err)
	}

	// judge pre process
	var res Response
	err = cbor.Unmarshal(resBin, &res)
	if err != nil {
		t.Fatal(err)
	}
	var resBody resSampleBody
	err = cbor.Unmarshal(res.Body, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	// judge
	if !reflect.DeepEqual(res, expectRes) {
		t.Fatalf("res=%v, expectRes=%v\n", res, expectRes)
	}
	if !reflect.DeepEqual(resBody, expectResBody) {
		t.Fatalf("resBody=%v, expectResBody=%v\n", resBody, expectResBody)
	}
}

func testHandler(res *Response, req *Request) error {
	// unmarshal request body
	var reqBody reqSampleBody
	err := cbor.Unmarshal([]byte(req.Body), &reqBody)
	if err != nil {
		return err
	}
	fmt.Printf("request: str=%s, num=%d\n", reqBody.TestStr, reqBody.TestNum)

	// marshal response body
	resBody := resSampleBody{
		TestStr: reqBody.TestStr,
		TestNum: reqBody.TestNum,
	}
	fmt.Printf("response: str=%s, num=%d\n", resBody.TestStr, resBody.TestNum)
	res.Body, err = cbor.Marshal(&resBody)
	if err != nil {
		return err
	}

	res.Header = Header{}
	res.StatusCode = 200
	return nil
}
