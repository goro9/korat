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
	expectRes := Res{
		Method: 0,
		Body:   expectResBodyBin,
	}

	// test data
	reqBody := reqSampleBody{
		TestStr: expectResBody.TestStr,
		TestNum: expectResBody.TestNum,
	}
	reqBodyBin, _ := cbor.Marshal(&reqBody)
	req := Req{
		Method: expectRes.Method,
		Auth:   nil,
		Path:   "test",
		Body:   reqBodyBin,
	}
	reqBin, _ := cbor.Marshal(&req)

	// execute
	HandleFunc("test", testHandler)
	resBin, err := Handle(reqBin)
	if err != nil {
		t.Fatal(err)
	}

	// judge pre process
	var res Res
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

func testHandler(bodyBin []byte) []byte {
	// unmarshal request body
	var reqBody reqSampleBody
	if err := cbor.Unmarshal(bodyBin, &reqBody); err != nil {
		return nil
	}
	fmt.Printf("request: str=%s, num=%d\n", reqBody.TestStr, reqBody.TestNum)

	// marshal response body
	resBody := resSampleBody{
		TestStr: reqBody.TestStr,
		TestNum: reqBody.TestNum,
	}
	fmt.Printf("response: str=%s, num=%d\n", resBody.TestStr, resBody.TestNum)
	resBodyBin, err := cbor.Marshal(&resBody)
	if err != nil {
		return nil
	}
	return resBodyBin
}
