package internal

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/api"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
	"log"
)

func AddImage(cli bce.Client, libName, uriPrefix string, image *model.ImageArgs) error {
	if libName == "" {
		return vcaError.ErrorLibNameMissed
	}
	err := image.CheckParams(true)
	if err == nil {
		err = api.AddImage(cli, uriPrefix+"/"+libName, image)
	}
	return err
}

func AddImages(cli bce.Client, uriPrefix string, images []*model.ImageList) []*model.ImageList {
	var failed []*model.ImageList
	for _, list := range images {
		if err := list.CheckParams(true, true); err != nil {
			log.Printf("add images faild: %s,%s", err, list)
			failed = append(failed, list)
			continue
		}
		if res := api.AddImages(cli, uriPrefix+"/"+list.LibName, list); res != nil {
			failed = append(failed, res)
		}
	}
	if len(failed) > 0 {
		return failed
	}
	return nil
}

func ListImage(cli bce.Client, uriPrefix string, images *model.ImageList) error {
	if err := images.CheckParams(false, true); err != nil {
		return err
	}
	return api.ListImage(cli, uriPrefix+"/"+images.LibName, images)
}

func DeleteImage(cli bce.Client, libName, uriPrefix string, image *model.ImageArgs, checkBrief bool) error {
	var err error
	err = image.CheckParams(checkBrief)
	if err == nil {
		if libName == "" {
			err = vcaError.ErrorLibNameMissed
		} else {
			err = api.DeleteImage(cli, uriPrefix+"/"+libName, image)
		}
	}
	return err
}

func DeleteImages(cli bce.Client, uriPrefix string, images []*model.ImageList) []*model.ImageList {
	var failed []*model.ImageList
	for _, list := range images {
		if err := list.CheckParams(true, false); err != nil {
			log.Printf("delete image faild %s,%s", err, list)
			failed = append(failed, list)
			continue
		}
		if res := api.DeleteImages(cli, uriPrefix+"/"+list.LibName, list); res != nil {
			failed = append(failed, res)
		}
	}
	if len(failed) > 0 {
		return failed
	}
	return nil
}
