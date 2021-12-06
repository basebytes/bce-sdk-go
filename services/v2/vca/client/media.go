package client

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/api"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
	"github.com/basebytes/bce-sdk-go/util"
)

type MediaClient struct {
	*bce.BceClient
}

func NewMediaClient(client *bce.BceClient) *MediaClient {
	return &MediaClient{client}
}

//media
func (c *MediaClient) PutMedia(args *model.PutMediaArgs) (*model.MediaResultCommon, error) {
	if err := args.CheckParams(); err != nil {
		return nil, err
	}
	return api.PutMedia(c, args)
}

func (c *MediaClient) GetMedia(source string) (*model.GetMediaResult, error) {
	if util.CheckStringLength(source, 1, 1024) {
		return api.GetMedia(c, map[string]string{"source": source})
	}
	return nil, model.NewInvalidParamValue("source", source)
}

func (c *MediaClient) GetSubTask(source string, task model.SubTask) (*model.SubTaskResult, error) {
	if util.CheckStringLength(source, 1, 1024) {
		return api.GetSubTask(c, task, map[string]string{"source": source})
	}
	return nil, model.NewInvalidParamValue("source", source)
}
