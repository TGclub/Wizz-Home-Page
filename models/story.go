package models
//历史事件
type Story struct {
	// ID is story's id
	ID            int   `json:"id" example:"1"`
	TimeStamp     int64   `example:"1580397149"`
	Name          string `example:"为之诞生"`
	StoryDescribe string `example:"11月,前身 TgClub 诞生"`
}
