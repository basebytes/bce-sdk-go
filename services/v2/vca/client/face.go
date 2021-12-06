package client

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/api"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/client/internal"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/constant"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
)

type FaceClient struct {
	*bce.BceClient
}

func NewFaceClient(client *bce.BceClient) *FaceClient {
	return &FaceClient{client}
}

//face lib

func (c *FaceClient) CreateFaceLib(args *model.PostImageLibArgs) error {
	if err := args.CheckParams(); err != nil {
		return err
	}
	return api.CreateImageLib(c, constant.FaceUriPrefix, args)
}

func (c *FaceClient) GetFaceLib() (*model.ImageLibList, error) {
	return api.GetImageLib(c, constant.FaceUriPrefix)
}

//face brief

func (c *FaceClient) ListFaceBrief(libName string) (*model.BriefList, error) {
	return api.ListImageBrief(c, constant.FaceUriPrefix+"/"+libName)
}

func (c *FaceClient) DeleteFaceBrief(libName, brief string) error {
	return api.DeleteImageBrief(c, constant.FaceUriPrefix+"/"+libName, brief)
}

//face image

func (c *FaceClient) AddFaceImage(libName string, image *model.ImageArgs) error {
	return internal.AddImage(c, libName, constant.FaceUriPrefix, image)
}

func (c *FaceClient) AddFaceImages(images []*model.ImageList) []*model.ImageList {
	return internal.AddImages(c, constant.FaceUriPrefix, images)
}

func (c *FaceClient) ListFaceImage(images *model.ImageList) error {
	return internal.ListImage(c, constant.FaceUriPrefix, images)
}

func (c *FaceClient) DeleteFaceImage(libName string, image *model.ImageArgs) error {
	return internal.DeleteImage(c, libName, constant.FaceUriPrefix, image, true)
}

func (c *FaceClient) DeleteFaceImages(images []*model.ImageList) []*model.ImageList {
	return internal.DeleteImages(c, constant.FaceUriPrefix, images)
}
