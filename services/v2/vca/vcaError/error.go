package vcaError

import (
	"errors"
	"fmt"
	"github.com/basebytes/bce-sdk-go/bce"
	"strings"
)

type AnalyzeError struct {
	*bce.BceServiceError
	//Code    string `json:"code,omitempty"`
	//Message string `json:"message,omitempty"`
}

func (e AnalyzeError) Error() string {
	return fmt.Sprintf("Error code[%s] messgae[%s]", e.Code, e.Message)
}

const (
	ErrorCodeInternalError    = "InternalError"
	ErrorCodeTimeout          = "TimeOut"
	ErrorCOdeInvalidParameter = "InvalidParameter"
	ErrorCodeInvalidMedia     = "InvalidMedia"

	ErrorMsgNoSuchSubtask      = "invalid media: media has no such subtask"
	ErrorMsgSubtaskNotFinished = "invalid media: subtask is not FINISHED"

	errorMsgAddedSuffix     = "already added before!"
	errorMsgNoSuchImgPrefix = "No such image="
)

func IsNoSuchSubtask(err error) bool {
	if err, OK := err.(*bce.BceServiceError); OK {
		return err.Message == ErrorMsgNoSuchSubtask
	}
	return false
}

func IsTimeOut(err error) bool {
	if err, OK := err.(*bce.BceServiceError); OK {
		return strings.HasSuffix(err.Code, ErrorCodeTimeout)
	}
	return false
}

func IsSubtaskNotFinished(err error) bool {
	if err, OK := err.(*bce.BceServiceError); OK {
		return err.Message == ErrorMsgSubtaskNotFinished
	}
	return false
}

func IsImageAdded(err error) bool {
	if err, OK := err.(*bce.BceServiceError); OK {
		return strings.HasSuffix(err.Message, errorMsgAddedSuffix)
	}
	return false
}

func IsNoSuchImage(err error) bool {
	if err, OK := err.(*bce.BceServiceError); OK {
		return strings.HasPrefix(err.Message, errorMsgNoSuchImgPrefix)
	}
	return false
}

var (
	ErrorLibNameMissed       = errors.New("No lib name specified. ")
	ErrorBriefNameMissed     = errors.New("No brief specified. ")
	ErrorCreateFaceLibFailed = errors.New("Create face lib failed. ")
	ErrorImageMissed         = errors.New("Image not found. ")
	ErrorDeleteImageFailed   = errors.New("Delete images failed. ")
	ErrorStopStreamFailed    = errors.New("Stop stream failed. ")
)
