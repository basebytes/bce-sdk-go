// media.go - the media APIs definition supported by the VCA service

// Package api defines all APIs supported by the VCA service of BCE.
package api

import (
	"encoding/json"
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/http"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/constant"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
	"io/ioutil"
)

func PutMedia(cli bce.Client, args *model.PutMediaArgs) (*model.MediaResultCommon, error) {
	if body, err := encode(args); err != nil {
		return nil, err
	} else if resp, err := sendRequest(cli, http.PUT, constant.MediaUriPrefix, body, nil); err != nil {
		return nil, err
	} else {
		jsonBody := &model.MediaResultCommon{}
		return jsonBody, resp.ParseJsonBody(jsonBody)
	}

}

func GetMedia(cli bce.Client, params map[string]string) (*model.GetMediaResult, error) {
	if resp, err := sendRequest(cli, http.GET, constant.MediaUriPrefix, nil, params); err != nil {
		//fmt.Println(resp.StatusCode())
		return nil, err
	} else {
		jsonBody := &model.GetMediaResult{}
		defer resp.Body().Close()
		b,_:=ioutil.ReadAll(resp.Body())

		return jsonBody,json.Unmarshal(b,jsonBody)
		//return jsonBody, resp.ParseJsonBody(jsonBody)
	}
}

func GetSubTask(cli bce.Client, task model.SubTask, params map[string]string) (*model.SubTaskResult, error) {
	if resp, err := sendRequest(cli, http.GET, constant.MediaUriPrefix+"/"+task.Value(), nil, params); err != nil {
		return nil, err
	} else {
		result := NewSubTaskResult(task)
		if err = resp.ParseJsonBody(result); err == nil {
			err = parseSubTask(result)
		}
		return result, err
	}
}

func parseSubTask(sub *model.SubTaskResult) error {
	var err error
	switch sub.Status {
	case constant.FINISHED:
		if sub.Type == constant.TaskImageClassify {
			if len(sub.AggResult) > 0 {
				err = json.Unmarshal([]byte(sub.AggResult), &sub.SubTaskItem)
			}
			if err == nil && len(sub.Result) > 0 {
				err = json.Unmarshal([]byte(sub.Result), &sub.SubTaskItem.(*model.ImageClassifyResult).Detail)
			}
		} else {
			if len(sub.Result) > 0 {
				err = json.Unmarshal([]byte(sub.Result), sub.SubTaskItem)
			}
		}

		//} else {
		//	err = errors.New(fmt.Sprintf("UnSupport sub task type %s,%s", sub.Type.Value(), sub.Result))
		//}
	case constant.ERROR:
	case constant.CANCELLED:
	}
	return err
}

func NewSubTaskResult(taskType model.SubTask) *model.SubTaskResult {
	result := &model.SubTaskResult{}
	if subResult, OK := interResult[taskType]; OK {
		result.SubTaskItem = subResult.New()
	}
	return result
}

var interResult = map[model.SubTask]model.SubTaskItem{
	//none
	constant.TaskVideo:        &model.VideoResult{},
	constant.TaskCover:        &model.CoverResult{},
	constant.TaskVideoCover:   &model.Other{},
	constant.TaskHighLight:    &model.Other{},
	constant.TaskThumbnail:    &model.ThumbnailResult{},
	constant.TaskCharacter:    &model.CharacterResult{},
	constant.TaskOcrStructure: &model.OcrStructureResult{},
	constant.TaskAudio:        &model.AudioResult{},
	constant.TaskSpeech:       &model.SpeechResult{},
	constant.TaskTitle:        &model.TitleResult{},
	constant.TaskFaceTracking: &model.FaceTrackingResult{},
	constant.TaskFaceDetect:   &model.Other{},

	//figure
	constant.TaskFaceRecognitionTracking:  &model.FaceRecTrackingResult{},
	constant.TaskFaceRecognitionThumbnail: &model.FaceRecThumbnailResult{},
	constant.TaskPrivateFaceTracking:      &model.PrivateFaceTrackingResult{},
	constant.TaskPrivateFaceImage:         &model.PrivateFaceImageResult{},

	//logo
	constant.TaskLogo:        &model.LogoResult{},
	constant.TaskPrivateLogo: &model.LogoResult{},

	//scenario
	constant.TaskScenarioClassifyV2: &model.ScenarioClassifyResultV2{},
	constant.TaskShortVideoClassify: &model.Other{},

	//entity
	constant.TaskLandmark:      &model.Other{},
	constant.TaskImageClassify: &model.ImageClassifyResult{},
	constant.TaskObjectDetect:  &model.ObjectDetectResult{},

	//keyword
	constant.TaskTextrankCharacter: &model.TextrankResult{},
	constant.TaskKeywordMerge:      &model.Other{},
	constant.TaskTextrankSpeech:    &model.TextrankResult{},

	//knowledge_graph
	constant.TaskKnowledgeGraph:     &model.KnowledgeGraphResult{},
	constant.TaskKnowledgeGraphPoem: &model.Other{},
}

var ScenarioTaskMap = map[model.Scenario][]model.SubTask{
	constant.FIGURE:         {constant.TaskFaceRecognitionTracking, constant.TaskFaceRecognitionThumbnail, constant.TaskPrivateFaceTracking, constant.TaskPrivateFaceImage},
	constant.LOGO:           {constant.TaskLogo, constant.TaskPrivateLogo},
	constant.ENTITY:         {constant.TaskLandmark, constant.TaskImageClassify, constant.TaskObjectDetect},
	constant.KEYWORD:        {constant.TaskTextrankCharacter, constant.TaskKeywordMerge, constant.TaskTextrankSpeech},
	constant.SCENARIO:       {constant.TaskScenarioClassifyV2, constant.TaskShortVideoClassify},
	constant.KnowledgeGraph: {constant.TaskKnowledgeGraph, constant.TaskKnowledgeGraphPoem},
}

var ScenarioSourceMap = map[model.Scenario][]model.Source{
	constant.FIGURE:         {constant.SourceFaceRecognition, constant.SourcePrivateFace},
	constant.LOGO:           {constant.SourceLogo, constant.SourcePrivateLogo},
	constant.ENTITY:         {constant.SourceImageClassify, constant.SourceObjectDetect, constant.SourceLandmark},
	constant.KEYWORD:        {constant.SourceCharacter, constant.SourceSpeech},
	constant.SCENARIO:       {constant.SourceScenarioClassify},
	constant.KnowledgeGraph: {constant.SourceKnowledgeGraph, constant.SourceKnowledgeGraphPoem},
}

var SubTaskSourceMap = map[model.SubTask]model.Source{
	constant.TaskCharacter:                constant.SourceCharacter,
	constant.TaskOcrStructure:             constant.SourceCharacter,
	constant.TaskSpeech:                   constant.SourceSpeech,
	constant.TaskFaceTracking:             constant.SourceFaceRecognition,
	constant.TaskFaceDetect:               constant.SourceFaceRecognition,
	constant.TaskFaceRecognitionTracking:  constant.SourceFaceRecognition,
	constant.TaskFaceRecognitionThumbnail: constant.SourceFaceRecognition,
	constant.TaskPrivateFaceTracking:      constant.SourcePrivateFace,
	constant.TaskPrivateFaceImage:         constant.SourcePrivateFace,
	constant.TaskLogo:                     constant.SourceLogo,
	constant.TaskPrivateLogo:              constant.SourcePrivateLogo,
	constant.TaskTextrankCharacter:        constant.SourceCharacter,
	constant.TaskKeywordMerge:             constant.SourceCharacter,
	constant.TaskTextrankSpeech:           constant.SourceSpeech,
	constant.TaskScenarioClassifyV2:       constant.SourceScenarioClassify,
	constant.TaskShortVideoClassify:       constant.SourceScenarioClassify,
	constant.TaskLandmark:                 constant.SourceLandmark,
	constant.TaskImageClassify:            constant.SourceImageClassify,
	constant.TaskObjectDetect:             constant.SourceObjectDetect,
	constant.TaskKnowledgeGraph:           constant.SourceKnowledgeGraph,
	constant.TaskKnowledgeGraphPoem:       constant.SourceKnowledgeGraphPoem,
}
