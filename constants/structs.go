package constants

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Task struct {
	Id int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
	Tags []string `json:"tags"`
	CreatedAt timestamppb.Timestamp `json:"createdAt"`
	PriorityCode int16 `json:"priorityCode"`
	ExpectedTime int16 `json:"expectedTime"`
	DueTime int16 `json:"dueTime"`
}

type View struct {
	Tail int32
	Sort string
	TagFilter string
	Regex string
}