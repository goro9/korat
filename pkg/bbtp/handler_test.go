package bbtp

import (
	"reflect"
	"testing"

	"github.com/fxamacker/cbor/v2"
)

func TestHandle(t *testing.T) {
	expectResBody := resSampleBody{
		TestStr: "hello world",
		TestNum: 255,
	}
	expectResBodyBin, _ := cbor.Marshal(&expectResBody)
	expectRes := Res{
		Method: 0,
		Body:   expectResBodyBin,
	}

	reqBody := reqSampleBody{
		TestStr: expectResBody.TestStr,
		TestNum: expectResBody.TestNum,
	}
	reqBodyBin, _ := cbor.Marshal(&reqBody)
	req := Req{
		Method: expectRes.Method,
		Auth:   nil,
		Path:   nil,
		Body:   reqBodyBin,
	}
	reqBin, _ := cbor.Marshal(&req)
	resBin, err := Handle(reqBin)
	if err != nil {
		t.Fatal(err)
	}

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

	if !reflect.DeepEqual(res, expectRes) {
		t.Fatalf("res=%v, expectRes=%v\n", res, expectRes)
	}
	if !reflect.DeepEqual(resBody, expectResBody) {
		t.Fatalf("resBody=%v, expectResBody=%v\n", resBody, expectResBody)
	}
}
