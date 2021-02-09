package bbtp

import (
	"github.com/fxamacker/cbor/v2"
)

func decodeReq(uuid string, data []byte, req *Request) error {
	req.UUID = uuid
	err := cbor.Unmarshal(data, &req)
	if err != nil {
		return err
	}
	return nil
}

func encodeRes(res *Response) ([]byte, error) {
	data, err := cbor.Marshal(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
