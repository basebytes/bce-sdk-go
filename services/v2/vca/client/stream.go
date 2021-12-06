package client

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/api"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
)

type StreamClient struct {
	*bce.BceClient
}

func NewStreamClient(client *bce.BceClient) *StreamClient {
	return &StreamClient{client}
}

func (c *StreamClient) PutStream(args *model.StreamArgs) (*model.StreamResultCommon, error) {
	if err := args.CheckParams(); err != nil {
		return nil, err
	}
	return api.PutStream(c, args)
}

func (c *StreamClient) GetStream(source string) (*model.GetStreamResult, error) {
	if source == "" {
		return nil, model.NewInvalidParamValue("source", source)
	}
	return api.GetStream(c, map[string]string{"source": source})
}

func (c *StreamClient) StopStream(source string) error {
	if source == "" {
		return model.NewInvalidParamValue("source", source)
	}
	return api.StopStream(c, &model.StreamArgs{Source: source})
}
