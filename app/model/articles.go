// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"wizz-home-page/app/model/internal"
)

// Articles is the golang structure for table articles.
type Articles internal.Articles

// Fill with you ideas below.
type ArticleApiCreateReq struct {
	Title      string `orm:"title"       json:"title"`      //
	ArticleUrl string `orm:"article_url" json:"articleUrl"` //
}

type ArticlesApiRep struct {
	Id         int    `orm:"id,primary"  json:"id"`         //
	Title      string `orm:"title"       json:"title"`      //
	ArticleUrl string `orm:"article_url" json:"articleUrl"` //
}
