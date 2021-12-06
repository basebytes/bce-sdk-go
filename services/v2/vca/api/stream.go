package api

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/http"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/constant"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
)

func PutStream(cli bce.Client, args *model.StreamArgs) (*model.StreamResultCommon, error) {
	if body, err := encode(args); err != nil {
		return nil, err
	} else if resp, err := sendRequest(cli, http.PUT, constant.StreamUriPrefix, body, nil); err != nil {
		return nil, err
	} else {
		jsonBody := &model.StreamResultCommon{}
		return jsonBody, resp.ParseJsonBody(jsonBody)
	}
}

func GetStream(cli bce.Client, params map[string]string) (*model.GetStreamResult, error) {
	if resp, err := sendRequest(cli, http.GET, constant.StreamUriPrefix, nil, params); err != nil {
		return nil, err
	} else {
		jsonBody := &model.GetStreamResult{}
		//resp.ParseResponse()
		//defer resp.Body().Close()
		//b,_:=ioutil.ReadAll(resp.Body())
		//fmt.Println(string(b))
		//return jsonBody,json.Unmarshal(b,jsonBody)
		return jsonBody, resp.ParseJsonBody(jsonBody)
	}
}

var stop = map[string]string{"stop": ""}

func StopStream(cli bce.Client, args *model.StreamArgs) error {
	var (
		resp      *bce.BceResponse
		body, err = encode(args)
	)
	if err == nil {
		if resp, err = sendRequest(cli, http.PUT, constant.StreamUriPrefix, body, stop); err == nil {
			resp.ParseResponse()
			if resp.StatusCode() != 200 {
				err = vcaError.ErrorStopStreamFailed
			}
		}
	}
	return err
}
