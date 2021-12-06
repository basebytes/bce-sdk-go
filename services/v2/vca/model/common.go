package model

import (
	"fmt"
	"sort"
)

type InvalidParamValue struct {
	name  string
	value interface{}
}

func (e InvalidParamValue) Error() string {
	return fmt.Sprintf("Invalid Param %s : [%v]", e.name, e.value)
}

func NewInvalidParamValue(name string, value interface{}) InvalidParamValue {
	return InvalidParamValue{name, value}
}

var (
	fullNil  = fullLocationName{}
	shortNil = shortLocationName{}
)

type Location struct {
	fullLocationName
	shortLocationName
}

//left,top,width,height
func (loc *Location) GetLocation() []int {
	l, t, w, h := 0, 0, 0, 0
	if loc.fullLocationName != fullNil {
		l, t, w, h = loc.LeftOffsetInPixel, loc.TopOffsetInPixel, loc.WidthInPixel, loc.HeightInPixel
	} else if loc.shortLocationName != shortNil {
		l, t, w, h = loc.Left, loc.Top, loc.Width, loc.Height
	}
	return []int{l, t, w, h}
}

type fullLocationName struct {
	LeftOffsetInPixel int `json:"leftOffsetInPixel,omitempty"`
	TopOffsetInPixel  int `json:"topOffsetInPixel,omitempty"`
	WidthInPixel      int `json:"widthInPixel,omitempty"`
	HeightInPixel     int `json:"heightInPixel,omitempty"`
}

type shortLocationName struct {
	Left   int `json:"left,omitempty"`
	Top    int `json:"top,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Word struct {
	Text  string    `json:"word"`
	Rect  *Location `json:"rect,omitempty"`
	Score float64   `json:"confidence"`
}

type HitsTime struct {
	StartInSecond int `json:"startTimestampInSecond"`
	EndInSecond   int `json:"endTimestampInSecond"`
	StartInMS     int `json:"startTimestampInMS"`
	EndInMS       int `json:"endTimestampInMS"`
}



type TimePeriodSlice []*TimePeriod

func (t TimePeriodSlice)Len() int{
	return len(t)
}

func (t TimePeriodSlice)Less(i,j int)bool{
	return t[i].Start<t[j].Start
}

func (t TimePeriodSlice)Swap(i,j int){
	t[i],t[j]=t[j],t[i]
}

func (t *TimePeriodSlice) Remove(period *TimePeriod){
	i,l,tmp:=0,t.Len(),*t
	if l==0||period.End<tmp[0].Start||period.Start>tmp[l-1].End{
		return
	}
	var res []*TimePeriod
	for ;i<l;i++{
		if tmp[i].End<period.Start{
			res=append(res, tmp[i])
		}else if tmp[i].End==period.Start{
			if tmp[i].Start==tmp[i].End{
				continue
			}
			tmp[i].End=period.Start-1
			res=append(res, tmp[i])
		}else if tmp[i].End<=period.End{
			if tmp[i].Start<period.Start{
				tmp[i].End=period.Start-1
				res=append(res, tmp[i])
			}
		}else{
			if tmp[i].Start<period.Start{
				res=append(res,
					NewTimePeriod(tmp[i].Start,period.Start-1),
					NewTimePeriod(period.End+1,tmp[i].End),
				)
			}else if tmp[i].Start<=period.End{
				tmp[i].Start=period.End+1
				res=append(res, tmp[i])
			}else {
				break
			}
		}

	}
	for ;i<l;i++{
		res=append(res,tmp[i])
	}
	*t=res
}

func (t *TimePeriodSlice) Append(period *TimePeriod){
	i,l,tmp:=0,t.Len(),*t
	var res []*TimePeriod
	for ;i<l;i++{
		if tmp[i].End+1<period.Start{
			res=append(res, tmp[i])
		}else if tmp[i].Start<=period.End+1{
			if period.Start>tmp[i].Start{
				period.Start=tmp[i].Start
			}
			if period.End<tmp[i].End{
				period.End=tmp[i].End
			}
		}else{
			break
		}
	}
	res=append(res, period)
	for ;i<l;i++{
		res=append(res,tmp[i])
	}
	*t=res
}

func (t *TimePeriodSlice)Init(){
	prev,l,tmp:=0,t.Len(),*t
	if l>1&&!sort.IsSorted(t){
		sort.Sort(t)
	}
	for i:=1;i<l;i++{
		if tmp[prev].End+1<tmp[i].Start{
			prev++
			tmp.Swap(prev,i)
		}else if tmp[prev].End<tmp[i].End{
			tmp[prev].End=tmp[i].End
		}
	}
	*t=tmp[:prev+1]
}

func NewTimePeriod(start,end int64) *TimePeriod{
	return &TimePeriod{Start: start,End: end}
}

type TimePeriod struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

func (t *TimePeriod) Near(p int64) bool {
	return t.End+1 == p || t.Start-1 == p
}

func (t *TimePeriod) Before(start int64) bool {
	return t.End < start
}

func (t *TimePeriod) After(end int64) bool {
	return t.Start > end
}

func (t *TimePeriod) IsPoint() bool {
	return t.Start == t.End
}

//return -1 if offset<start,1 if offset> end ,else 0
func (t *TimePeriod) InPeriod(offset int64) int {
	switch {
	case offset < t.Start:
		return -1
	case offset > t.End:
		return 1
	default:
		return 0
	}
}

func (t *TimePeriod)Equal(o *TimePeriod) bool{
	if t != o {
		return t.Start == o.Start && t.End == o.End
	}
	return true
}

type Status string

const UnLimitInt = 1<<31 - 1
