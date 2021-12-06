package client

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/api"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/client/internal"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/constant"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
	"net/url"
)

type LogoClient struct {
	*bce.BceClient
}

func NewLogoClient(client *bce.BceClient) *LogoClient {
	return &LogoClient{client}
}

//logo lib

func (c *LogoClient) CreateLogoLib(args *model.PostImageLibArgs) error {
	if err := args.CheckParams(); err != nil {
		return err
	}
	return api.CreateImageLib(c, constant.LogoUriPrefix, args)
}

func (c *LogoClient) GetLogoLib() (*model.ImageLibList, error) {
	return api.GetImageLib(c, constant.LogoUriPrefix)
}

//logo brief

func (c *LogoClient) ListLogoBrief(libName string) (*model.BriefList, error) {
	return api.ListImageBrief(c, constant.LogoUriPrefix+"/"+libName)
}

func (c *LogoClient) DeleteLogoBrief(libName, brief string) error {
	return api.DeleteImageBrief(c, constant.LogoUriPrefix+"/"+libName, brief)
}

//logo image

func (c *LogoClient) AddLogoImage(libName string, image *model.ImageArgs) error {
	return internal.AddImage(c, libName, constant.LogoUriPrefix, image)
}

func (c *LogoClient) AddLogoImages(images []*model.ImageList) []*model.ImageList {
	return internal.AddImages(c, constant.LogoUriPrefix, images)
}

func (c *LogoClient) ListLogoImage(images *model.ImageList) error {
	return internal.ListImage(c, constant.LogoUriPrefix, images)
}

func (c *LogoClient) DeleteLogoImage(libName string, rawImage *model.ImageArgs) error {
	image := &model.ImageArgs{
		Image: url.QueryEscape(rawImage.Image),
	}
	return internal.DeleteImage(c, libName, constant.LogoUriPrefix, image, false)
}

func (c *LogoClient) DeleteLogoImages(rawImages []*model.ImageList) []*model.ImageList {
	var images []*model.ImageList
	for _, rawImage := range rawImages {
		if rawImage != nil {
			image := &model.ImageList{LibName: rawImage.LibName}
			var imgs []string
			for _, img := range rawImage.Images {
				imgs = append(imgs, url.QueryEscape(img))
			}
			image.Images = imgs
			images = append(images, image)
		}
	}
	return internal.DeleteImages(c, constant.LogoUriPrefix, images)
}
