package model

import (
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
	"github.com/basebytes/bce-sdk-go/util"
)

type StreamArgs struct {
	Source           string `json:"source"`
	Preset           string `json:"preset,omitempty"`
	Notification     string `json:"notification"`
	IntervalInSecond int    `json:"intervalInSecond,omitempty"`
	Description      string `json:"description,omitempty"`
}

func (args *StreamArgs) CheckParams() error {
	var err error
	switch {
	case !util.CheckStringLength(args.Source, 1, UnLimitInt):
		err = NewInvalidParamValue("source", args.Source)
	case !util.CheckStringLength(args.Preset, 0, 40):
		err = NewInvalidParamValue("preset", args.Preset)
	case !util.CheckStringLength(args.Notification, 1, 40):
		err = NewInvalidParamValue("notification", args.Notification)
	case !util.CheckStringLength(args.Description, 0, 100):
		err = NewInvalidParamValue("description", args.Description)
	case !util.CheckIntRange(args.IntervalInSecond, 1, UnLimitInt):
		args.IntervalInSecond = 10
	}
	return err
}

type StreamResultCommon struct {
	StreamId         string `json:"streamId,omitempty"`
	Source           string `json:"source,omitempty"`
	Preset           string `json:"preset,omitempty"`
	Notification     string `json:"notification,omitempty"`
	Description      string `json:"description,omitempty"`
	IntervalInSecond int    `json:"intervalInSecond,omitempty"`
	CreateTime       string `json:"createTime,omitempty"`
	Status           Status `json:"status,omitempty"`
}

type GetStreamResult struct {
	StreamResultCommon
	StartTime        string                 `json:"start_time,omitempty"`
	DurationInSecond int                    `json:"duration_in_second,omitempty"`
	AnalyzerError    *vcaError.AnalyzeError `json:"error,omitempty"`
}
