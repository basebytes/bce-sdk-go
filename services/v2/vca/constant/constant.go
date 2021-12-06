// constant.go - definitions of the request arguments and results data structure model for VCA

package constant

import (
	"github.com/basebytes/bce-sdk-go/bce"
	"github.com/basebytes/bce-sdk-go/services/v2/vca/model"
)

const (
	UriPrefix = bce.UriPrefix + "v2"

	MediaUri  = "/media"
	StreamUri = "/stream"
	FaceUri   = "/face/lib"
	LogoUri   = "/logo/lib"

	MediaUriPrefix  = UriPrefix + MediaUri  //视频
	StreamUriPrefix = UriPrefix + StreamUri //直播
	FaceUriPrefix   = UriPrefix + FaceUri   //自定义人脸库
	LogoUriPrefix   = UriPrefix + LogoUri   //自定义logo库
)

const (
	//undefined
	TaskVideo           model.SubTask = "video"             //视频转码结果
	TaskCover           model.SubTask = "cover"             //视频封面选取
	TaskVideoCover model.SubTask = "video_cover" //视频精彩图 TODO
	TaskHighLight  model.SubTask = "highlight"   //视频精彩片段 TODO
	TaskThumbnail       model.SubTask = "thumbnail"         //以一定策略从视频截取的缩略图集合
	TaskCharacter       model.SubTask = "character"         //视频缩略图通过OCR技术获取的文字识别结果
	TaskOcrStructure    model.SubTask = "ocr_structure"     //文字结构化
	TaskAudio           model.SubTask = "audio"             //视频音频
	TaskSpeech          model.SubTask = "speech"            //视频音频通过ASR技术获取的语音识别结果
	TaskTitle           model.SubTask = "title"             //视频标题
	TaskFaceTracking    model.SubTask = "face_tracking"     //人脸追踪特征
	TaskFaceDetect model.SubTask = "face_detect" //人脸属性检测 TODO
	//figure
	TaskFaceRecognitionTracking  model.SubTask = "face_recognition_tracking"  //tracking版本的公有人脸识别模型
	TaskFaceRecognitionThumbnail model.SubTask = "face_recognition_thumbnail" //非tracking版本的公有人脸识别模型
	TaskPrivateFaceTracking      model.SubTask = "private_face_tracking"      //tracking版本的私有人脸识别模型
	TaskPrivateFaceImage model.SubTask = "private_face_image" //非tracking版本的私有人脸识别模型 TODO
	//logo
	TaskLogo model.SubTask = "logo" //公有logo识别模型
	TaskPrivateLogo model.SubTask = "private_logo" //私有logo识别模型 //TODO
	//keyword
	TaskTextrankCharacter model.SubTask = "textrank_character" //对文字识别结果textrank（通用版本）
	TaskKeywordMerge model.SubTask = "keyword_merge" //对文字识别结果textrank（小视频专用版本） TODO
	TaskTextrankSpeech    model.SubTask = "textrank_speech"    //对语音识别结果textrank
	//scenario
	TaskScenarioClassifyV2 model.SubTask = "scenario_classify_v2" //视频场景分类模型（通用版本）
	TaskShortVideoClassify model.SubTask = "short_video_classify" //场景分类模型（小视频专用版本） TODO
	//entity
	TaskLandmark model.SubTask = "landmark" //地标识别 TODO
	TaskImageClassify model.SubTask = "image_classify" //图像分类模型
	TaskObjectDetect  model.SubTask = "object_detect"  //物体检测模型
	//knowledge_graph
	TaskKnowledgeGraph model.SubTask = "knowledge_graph" //知识图谱
	TaskKnowledgeGraphPoem model.SubTask = "knowledge_graph_poem" //诗词识别 TODO
)

const (
	PROVISIONING model.Status = "PROVISIONING" //排队中
	PROCESSING   model.Status = "PROCESSING"   //进行中
	FINISHED     model.Status = "FINISHED"     //分析结束
	ERROR        model.Status = "ERROR"        //分析失败
	CANCELLED    model.Status = "CANCELLED"    //分析取消
)

const (
	FIGURE         model.Scenario = "figure"          //人脸
	LOGO           model.Scenario = "logo"            //logo
	KEYWORD        model.Scenario = "keyword"         //关键字
	SCENARIO       model.Scenario = "scenario"        //场景
	ENTITY         model.Scenario = "entity"          //实体
	KnowledgeGraph model.Scenario = "knowledge_graph" //知识图谱
)

const (
	//figure
	SourceFaceRecognition model.Source = "face_recognition" //公众人脸识别
	SourcePrivateFace     model.Source = "private_face"     //自定义人脸识别
	//logo
	SourceLogo        model.Source = "logo"         //公众logo识别
	SourcePrivateLogo model.Source = "private_logo" //自定义logo识别
	//keyword
	SourceCharacter model.Source = "character" //文字识别
	SourceSpeech    model.Source = "speech"    //语音识别
	//scenario
	SourceScenarioClassify model.Source = "scenario_classify" //视频场景分类
	//entity
	SourceLandmark      model.Source = "landmark"       //地标识别
	SourceImageClassify model.Source = "image_classify" //图像分类
	SourceObjectDetect  model.Source = "object_detect"  //物体识别
	//knowledge_graph
	SourceKnowledgeGraph     model.Source = "knowledge_graph"      //知识图谱
	SourceKnowledgeGraphPoem model.Source = "knowledge_graph_poem" //诗词识别
)
