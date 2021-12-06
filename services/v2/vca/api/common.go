package api

import (
	"encoding/json"
	"fmt"
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
)

func sendRequest(cli bce.Client, method, uri string, body *bce.Body, params map[string]string) (*bce.BceResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(uri)
	req.SetMethod(method)
	if body != nil {
		req.SetBody(body)
	}
	if params != nil {
		for k, v := range params {
			req.SetParam(k, v)
		}
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		fmt.Println(resp.StatusCode())
		return nil, err
	}
	if resp.IsFail() {
		return nil, vcaError.AnalyzeError{BceServiceError: resp.ServiceError()}
	}
	return resp, nil
}

func encode(value interface{}) (*bce.Body, error) {
	if jsonBytes, err := json.Marshal(value); err != nil {
		return nil, err
	} else {
		return bce.NewBodyFromBytes(jsonBytes)
	}
}
