package model

import (
	"github.com/basebytes/bce-sdk-go/services/v2/vca/vcaError"
	"github.com/basebytes/bce-sdk-go/util"
	"time"
)

type PutMediaArgs struct {
	Source       string `json:"source"`
	Auth         string `json:"auth,omitempty"`
	Title        string `json:"title,omitempty"`
	SubTitle     string `json:"subTitle,omitempty"`
	Category     string `json:"category,omitempty"`
	Description  string `json:"description,omitempty"`
	Preset       string `json:"preset,omitempty"`
	Notification string `json:"notification,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

func (args *PutMediaArgs) CheckParams() error {
	var err error
	switch {
	case !util.CheckStringLength(args.Source, 1, 1024):
		err = NewInvalidParamValue("source", args.Source)
	case !util.CheckStringLength(args.Title, 0, 256):
		err = NewInvalidParamValue("title", args.Title)
	case !util.CheckStringLength(args.SubTitle, 0, 100):
		err = NewInvalidParamValue("subTitle", args.SubTitle)
	case !util.CheckStringLength(args.Preset, 0, 40):
		err = NewInvalidParamValue("preset", args.Preset)
	case !util.CheckStringLength(args.Notification, 0, 40):
		err = NewInvalidParamValue("notification", args.Notification)
	case !util.CheckStringLength(args.Category, 0, 100):
		err = NewInvalidParamValue("category", args.Category)
	case !util.CheckStringLength(args.Description, 0, 256):
		err = NewInvalidParamValue("description", args.Description)
	case !util.CheckIntRange(args.Priority, 0, 100):
		err = NewInvalidParamValue("priority", args.Priority)
	}
	return err
}

type MediaResultCommon struct {
	TaskId       string `json:"taskId,omitempty"`
	Source       string `json:"source,omitempty"`
	MediaId      string `json:"mediaId,omitempty"`
	Title        string `json:"title,omitempty"`
	SubTitle     string `json:"subTitle,omitempty"`
	Category     string `json:"category,omitempty"`
	Description  string `json:"description,omitempty"`
	Preset       string `json:"preset,omitempty"`
	Notification string `json:"notification,omitempty"`
	Status       Status `json:"status,omitempty"`
	Percent      int    `json:"percent,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
}

type GetMediaResult struct {
	MediaResultCommon
	StartTime        *time.Time             `json:"startTime,omitempty"`
	PublishTime      *time.Time             `json:"publishTime,omitempty"`
	DurationInSecond int                    `json:"durationInSecond,omitempty"`
	Results          []*AnalyzeResult       `json:"results,omitempty"`
	Error            *vcaError.AnalyzeError `json:"error,omitempty"`
}

type AnalyzeResult struct {
	Type       Scenario            `json:"type,omitempty"`
	Attributes []*AnalyzeAttribute `json:"result,omitempty"`
}

type AnalyzeAttribute struct {
	Attribute     string        `json:"attribute"`
	Score         float32       `json:"confidence"`
	Source        Source        `json:"source"`
	Version       string        `json:"version"`
	AttributeTime []*TimePeriod `json:"time"`
}


type SubTaskResult struct {
	TaskId    string  `json:"taskId"`
	SubTaskId string  `json:"subTaskId"`
	Source    string  `json:"source"`
	Type      SubTask `json:"type"`
	Status    Status  `json:"status"`
	Result    string  `json:"result"`
	AggResult string  `json:"aggregationResult,omitempty"`
	SubTaskItem
}

type SubTaskItem interface {
	New() SubTaskItem
}

//video
type VideoResult struct {
	VideoUrl string
}

func (v VideoResult)New() SubTaskItem{
	return &VideoResult{}
}

//cover
type CoverResult struct {
	CoverUrl string
}

func (c CoverResult)New() SubTaskItem{
	return &CoverResult{}
}

//thumbnail
type ThumbnailResult struct {
	Prefix string   `json:"thumbnailPrefix"`
	Urls   []string `json:"thumbnailUrls"` //可能为空
}

func (t ThumbnailResult)New() SubTaskItem {
	return &ThumbnailResult{}
}

type Other map[string]interface{}

func (o Other)New() SubTaskItem{
	return &Other{}
}

type CharacterResult []*CharacterResultItem

//character
type CharacterResultItem struct {
	Image     string
	Timestamp int
	Words     []*Word
}

func (c CharacterResult)New() SubTaskItem{
	return &CharacterResult{}
}

//ocr_structure
type OcrStructureResult struct {
	Status  string
	Results map[string][]*OcrStructureResultItem
}

func (c OcrStructureResult)New() SubTaskItem{
	return &OcrStructureResult{}
}

//audio
type AudioResult struct {
	AudioUrl string `json:"audioUrl"`
}

func (c AudioResult)New() SubTaskItem{
	return &AudioResult{}
}

type SpeechResult []*SpeechResultItem
//speech
type SpeechResultItem struct {
	*HitsTime
	Text string `json:"statement"`
}

func (c SpeechResult)New() SubTaskItem{
	return &SpeechResult{}
}

//title
type TitleResult struct {
	Title string `json:"title"`
}

func (c TitleResult)New() SubTaskItem{
	return &TitleResult{}
}

//face_tracking
type FaceTrackingResult struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Faces   []*FaceBaseItem `json:"faces"`
}

func (c FaceTrackingResult)New() SubTaskItem{
	return &FaceTrackingResult{}
}


//face_recognition_tracking
type FaceRecTrackingResult struct {
	Faces []*FaceRecTrackingData `json:"faces"`
}

func (f FaceRecTrackingResult)New() SubTaskItem{
	return &FaceRecTrackingResult{}
}

//face_recognition_thumbnail
type FaceRecThumbnailResult struct {
	Faces []*FaceRecThumbnailData `json:"faces"`
}

func (f FaceRecThumbnailResult)New() SubTaskItem{
	return &FaceRecThumbnailResult{}
}

//private_face_tracking
type PrivateFaceTrackingResult struct {
	Status    string                     `json:"status"`
	Threshold float32                    `json:"threshold"`
	Faces     []*PrivateFaceTrackingData `json:"faces"`
}

func (p PrivateFaceTrackingResult)New() SubTaskItem{
	return &PrivateFaceTrackingResult{}
}

//private_face_image
type PrivateFaceImageResult struct {
	Status  string                  `json:"status"`
	Results []*PrivateFaceImageData `json:"results"`
}

func (p PrivateFaceImageResult)New() SubTaskItem{
	return &PrivateFaceImageResult{}
}

//logo private_logo
type LogoResultItem struct {
	Image     string      `json:"image"`
	Timestamp int         `json:"timestamp"`
	Result    []*LogoItem `json:"result"`
}
type LogoResult []*LogoResultItem
func (l LogoResult)New() SubTaskItem{
	return &LogoResult{}
}

//scenario_classify_v2
type ScenarioClassifyResultV2 struct {
	Results []*ScenarioClassifyItemV2 `json:"results"`
}

func (s ScenarioClassifyResultV2)New() SubTaskItem{
	return &ScenarioClassifyResultV2{}
}

//image_classify
type ImageClassifyResult struct {
	Detail []*ImageClassifyItem      `json:"detail"`
	Agg    []*ImageClassifyAggResult `json:"result"`
}

func (i ImageClassifyResult)New() SubTaskItem{
	return &ImageClassifyResult{}
}

type ObjectDetectResult []*ObjectDetectResultItem
//object_detect
type ObjectDetectResultItem struct {
	Image          string              `json:"image"`
	Timestamp      int                 `json:"timestamp"`
	ClassifyResult []*ObjectDetectItem `json:"classifyResult"`
}

func (o ObjectDetectResult) New() SubTaskItem{
	return &ObjectDetectResult{}
}

//textrank_character textrank_speech
type TextrankResult []*TextrankResultItem
type TextrankResultItem struct {
	Keyword  string  `json:"keyword"`
	Weight   float64 `json:"weight"`
	TimeList [][]int `json:"time_list"`
}

func (t TextrankResult) New()SubTaskItem{
	return &TextrankResult{}
}

//knowledge_graph
type KnowledgeGraphResult struct {
	Genre               []string      `json:"genre,omitempty"`               //类型
	Name                []string      `json:"name,omitempty"`                //名称
	Country             []string      `json:"country,omitempty"`             //国家
	Region              []string      `json:"region,omitempty"`              //地区
	InLanguage          []string      `json:"inLanguage,omitempty"`          //内容语言
	DatePublished       []string      `json:"datePublished,omitempty"`       //年代
	RegionalReleaseDate []string      `json:"regionalReleaseDate,omitempty"` //上映时间
	ProductionCompany   []string      `json:"productionCompany,omitempty"`   //出品方
	PartOfSeries        []interface{} `json:"partOfSeries,omitempty"`
	Cast                []interface{} `json:"cast,omitempty"`
	Starring            []interface{} `json:"starring,omitempty"`
}

func (k KnowledgeGraphResult)New()SubTaskItem{
	return &KnowledgeGraphResult{}
}



type FaceBaseItem struct {
	FaceUrl    string `json:"faceUrl,omitempty"`
	Image      string `json:"image"`
	FrameNum   int    `json:"frame_num"`
	StartFrame int    `json:"start_frame"`
}

type FaceRecTrackingData struct {
	*FaceBaseItem
	Data *FaceRecTrackingDataItem `json:"data"`
}
type FaceRecThumbnailData struct {
	*FaceBaseItem
	Data     *FaceRecThumbnailDataItem `json:"data"`
	Location *FaceWithDegree           `json:"location"`
}

type PrivateFaceTrackingData struct {
	*FaceBaseItem
	Data     *PrivateFaceTrackingDataItem   `json:"data"`
	DataList []*PrivateFaceTrackingDataItem `json:"dataList"`
}

type PrivateFaceImageData struct {
	Image  string                      `json:"image"`
	Result []*PrivateFaceImageDataItem `json:"result"`
}

type FaceWithDegree struct {
	Left   float32 `json:"left"`
	Top    float32 `json:"top"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Degree float32 `json:"degree"`
}

type BaikeInfo map[string]interface{}

type FaceRecTrackingDataItem struct {
	Name      string     `json:"name,omitempty"`
	Score     float64    `json:"radio,omitempty"`
	BaikeInfo *BaikeInfo `json:"baikeInfo,omitempty"`
}

type FaceRecThumbnailDataItem struct {
	Name      string     `json:"name,omitempty"`
	Score     float64    `json:"radio,omitempty"`
	BaikeInfo *BaikeInfo `json:"baikeInfo,omitempty"`
}

type PrivateFaceTrackingDataItem struct {
	Name  string  `json:"name,omitempty"`
	Score float64 `json:"score,omitempty"`
}

type PrivateFaceImageDataItem struct {
	Brief string  `json:"brief,omitempty"`
	Score float64 `json:"score,omitempty"`
}

type LogoItem struct {
	Name     string    `json:"name"`
	Score    float64   `json:"probability"`
	Location *Location `json:"location"`
}

type ScenarioClassifyItemV2 struct {
	StartFrame int                   `json:"startframe"`
	EndFrame   int                   `json:"endframe"`
	Labels     []*ScenarioClassifyV2 `json:"labels"`
}

type ScenarioClassifyV2 struct {
	Name  string  `json:"name"`
	Score float64 `json:"confidence"`
}

type ImageClassify struct {
	Name  string  `json:"class_name"`
	Score float64 `json:"probability"`
}

type VideoClassify struct {
	Name  string  `json:"value"`
	Score float64 `json:"probability"`
}

type ImageClassifyAggResult struct {
	ClassName string    `json:"className"`
	ScoreList []float64 `json:"confidenceList"`
	TimeList  []*TimePeriod `json:"timeListForHigherThanThreshold"`
	Score     float64 `json:"score"`
	MeanScore float64 `json:"meanConfidence"`
}


type ImageClassifyItem struct {
	Image          string           `json:"image"`
	Timestamp      int              `json:"timestamp"`
	BlurScore      float64          `json:"blurScore"`
	ClassifyResult []*ImageClassify `json:"classifyResult"`
}

type ObjectDetectItem struct {
	Keyword string  `json:"keyword"`
	Score   float64 `json:"score"`
	Root    string  `json:"root"`
}

type VideoClassifyKgItem struct {
	Title    *VideoClassify   `json:"title"`
	Subtitle []*VideoClassify `json:"subtitle"`
}

type OcrStructureResultItem struct {
	StartTimeInSeconds int       `json:"startTimeInSeconds"`
	EndTimeInSeconds   int       `json:"endTimeInSeconds"`
	Word               string    `json:"word"`
	Score              float64   `json:"confidence"`
	Rect               *Location `json:"rect"`
	Duration           int       `json:"duration"`
	TimeList           [][]int   `json:"time_list"`
}

type SubTask string

func (s SubTask) Value() string {
	return string(s)
}

type Scenario string

func (s Scenario) Value() string {
	return string(s)
}

type Source string

func (s Source) Value() string {
	return string(s)
}
