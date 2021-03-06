// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"wizz-home-page/app/model/internal"
)

// Stories is the golang structure for table stories.
type Stories internal.Stories

// Fill with you ideas below.

type StoryApiCreateReq struct {
	TimeStamp     int64  `json:"timeStamp"`
	Name          string `json:"name"`
	StoryDescribe string `json:"storyDescribe"`
}

type StoriesApiRep struct {
	Id            int    `orm:"id,primary"     json:"id"`            //
	TimeStamp     int64  `orm:"time_stamp"     json:"timeStamp"`     //
	Name          string `orm:"name"           json:"name"`          //
	StoryDescribe string `orm:"story_describe" json:"storyDescribe"` //
}
