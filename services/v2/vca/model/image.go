package model

import (
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
	"github.com/basebytes/bce-sdk-go/util"
	"time"
)

type PostImageLibArgs struct {
	Lib         string `json:"lib"`                   //库名称，长度不超过20，支持小写字母、数字和_，以字母开头，且必须全局唯一
	Description string `json:"description,omitempty"` //图片描述，长度不超过256
}

func (args *PostImageLibArgs) CheckParams() error {
	var err error
	if !util.CheckStringLength(args.Lib, 1, 20) {
		err = NewInvalidParamValue("lib", args.Lib)
	}
	return err
}

type ImageLibList struct {
	Libs []*ImageLib `json:"libs"`
}

type ImageLib struct {
	UserId      string    `json:"-"`
	Lib         string    `json:"lib"`
	CreateTime  time.Time `json:"createTime"`
	Description string    `json:",omitempty"`
}

type ImageList struct {
	LibName string   `json:"libName,omitempty"`
	Brief   string   `json:"brief,omitempty"`
	Images  []string `json:"images"`
}

func (i *ImageList) CheckParams(checkImages, checkBrief bool) error {
	var err error
	if i.LibName == "" {
		err = vcaError.ErrorLibNameMissed
	} else if checkBrief && i.Brief == "" {
		err = vcaError.ErrorBriefNameMissed
	} else if checkImages && (i.Images == nil || len(i.Images) == 0) {
		err = vcaError.ErrorImageMissed
	}
	return err
}

type ImageArgs struct {
	Image string `json:"image,omitempty"` //图片url
	Brief string `json:"brief,omitempty"` //图片集名称
}

func (args *ImageArgs) CheckParams(checkBrief bool) error {
	var err error
	if !util.CheckStringLength(args.Image, 1, UnLimitInt) {
		err = NewInvalidParamValue("image", args.Image)
	} else if checkBrief && !util.CheckStringLength(args.Brief, 1, UnLimitInt) {
		err = NewInvalidParamValue("brief", args.Brief)
	}
	return err
}

type BriefList struct {
	Briefs []string `json:"briefs,omitempty"`
}
