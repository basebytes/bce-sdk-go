package api

import (
	"errors"
	"fmt"
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/http"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
	"log"
	"net/url"
)

//common for face and logo

//image lib

func CreateImageLib(cli bce.Client, uri string, args *model.PostImageLibArgs) error {
	if body, err := encode(args); err != nil {
		return err
	} else {
		if resp, err := sendRequest(cli, http.POST, uri, body, nil); err == nil {
			resp.ParseResponse()
			if resp.StatusCode() != 200 {
				err = vcaError.ErrorCreateFaceLibFailed
			}
		}
		return err
	}
}

func GetImageLib(cli bce.Client, uri string) (*model.ImageLibList, error) {
	if resp, err := sendRequest(cli, http.GET, uri, nil, nil); err == nil {
		faceLibList := &model.ImageLibList{}
		return faceLibList, resp.ParseJsonBody(faceLibList)
	} else {
		return nil, err
	}
}

//image brief

func ListImageBrief(cli bce.Client, uri string) (*model.BriefList, error) {
	resp, err := sendRequest(cli, http.GET, uri, nil, nil)
	if err == nil {
		briefList := &model.BriefList{}
		return briefList, resp.ParseJsonBody(briefList)
	}
	return nil, err
}

//目前不能直接删除图集
func DeleteImageBrief(cli bce.Client, uri, brief string) error {
	params := map[string]string{"brief": brief}
	resp, err := sendRequest(cli, http.GET, uri, nil, params)
	if err == nil {
		resp.ParseResponse()
		if resp.StatusCode() != 200 {
			err = errors.New(fmt.Sprintf("delete brief[%s] failed", brief))
		}
	}
	return err
}

//image

//若brief不存在会自动创建
func AddImage(cli bce.Client, uri string, args *model.ImageArgs) error {
	var (
		err  error
		resp *bce.BceResponse
		body *bce.Body
	)
	if body, err = encode(args); err == nil {
		if resp, err = sendRequest(cli, http.POST, uri, body, nil); err == nil {
			resp.ParseResponse()
			if resp.StatusCode() != 200 {
				err = errors.New(fmt.Sprintf("add image[%s] failed", args.Image))
			}
		} else if vcaError.IsImageAdded(err) {
			log.Printf("image[%s] alreay added before!", args.Image)
			err = nil
		}
	}
	return err
}

func AddImages(cli bce.Client, uri string, images *model.ImageList) *model.ImageList {
	var (
		err  error
		args = &model.ImageArgs{Brief: images.Brief}
		resp *bce.BceResponse
		body *bce.Body
	)
	var failed []string
	for _, image := range images.Images {
		args.Image = image
		body, err = encode(args)
		if err == nil {
			if resp, err = sendRequest(cli, http.POST, uri, body, nil); err == nil {
				resp.ParseResponse()
				if resp.StatusCode() != 200 {
					failed = append(failed, image)
				}
			}
		}
		if err != nil {
			if vcaError.IsImageAdded(err) {
				log.Printf("image[%s] alreay added before!", image)
				continue
			}
			failed = append(failed, image)
		}
	}
	if len(failed) > 0 {
		return &model.ImageList{LibName: images.LibName, Brief: images.Brief, Images: failed}
	}
	return nil
}

func ListImage(cli bce.Client, uri string, images *model.ImageList) error {
	params := map[string]string{"brief": images.Brief}
	resp, err := sendRequest(cli, http.GET, uri, nil, params)
	if err == nil {
		err = resp.ParseJsonBody(images)
	}
	return err
}

//若删除的图片为brief中最后一张图片，则brief同时会被删除
func DeleteImage(cli bce.Client, uri string, image *model.ImageArgs) error {
	params := map[string]string{"image": image.Image}
	if image.Brief != "" {
		params["brief"] = image.Brief
	}
	resp, err := sendRequest(cli, http.DELETE, uri, nil, params)
	if err == nil {
		resp.ParseResponse()
		if resp.StatusCode() != 200 {
			err = errors.New(fmt.Sprintf("delete image[%s] failed", params["image"]))
		}
	} else if vcaError.IsNoSuchImage(err) {
		log.Printf("No such image %s ignored. ", image.Image)
		err = nil
	}
	return err
}

func DeleteImages(cli bce.Client, uri string, images *model.ImageList) *model.ImageList {
	if images == nil || len(images.Images) == 0 {
		return nil
	}
	params := map[string]string{}
	if images.Brief != "" {
		params["brief"] = images.Brief
	}
	var failed []string
	for _, image := range images.Images {
		params["image"] = image
		resp, err := sendRequest(cli, http.DELETE, uri, nil, params)
		if err == nil {
			resp.ParseResponse()
			if resp.StatusCode() != 200 {
				rawImg := image
				if images.Brief == "" {
					rawImg, _ = url.QueryUnescape(image)
				}
				failed = append(failed, rawImg)
			}
		} else {
			rawImg := image
			if images.Brief == "" {
				rawImg, _ = url.QueryUnescape(image)
			}
			if vcaError.IsNoSuchImage(err) {
				log.Printf("No such image %s ignored. ", rawImg)
			} else {
				failed = append(failed, rawImg)
			}
		}
	}
	if len(failed) > 0 {
		return &model.ImageList{LibName: images.LibName, Brief: images.Brief, Images: failed}
	}
	return nil
}
