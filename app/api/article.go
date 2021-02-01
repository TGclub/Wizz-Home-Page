package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Article = new(articlesApi)

type articlesApi struct{}

// @Summary 获取所有文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Articles
// @Router /api/articles [get]
func (*articlesApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var articles []model.Articles
	if err := dao.Articles.Structs(&articles); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(articles) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		response.JsonOld(r, 200, articles)
	}
}

// @Summary 获取一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Success 200 {object} model.Articles
// @Failure 404 {string} string "{"message":"Article not found"}"
// @Router /api/articles/{id} [get]
func (*articlesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var articles model.Articles
	if err := dao.Articles.Where("id = ", id).Struct(&articles); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, articles)
}

// @Summary 添加一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   articles      body model.Articles true  "文章"
// @Success 200 {object} model.Articles
// @Router /api/articles [POST]
// @Security JWT
func (*articlesApi) Create(r *ghttp.Request) {
	var (
		apiReq   *model.ArticleApiCreateReq
		articles *model.Articles
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a articles")
	}
	if err := gconv.Struct(apiReq, &articles); err != nil {
		response.JsonOld(r, 400, "not a articles")
	}
	if result, err := dao.Articles.Insert(articles); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		articles.Id = gconv.Int(id)
		response.JsonOld(r, 200, articles)
	}
}

// @Summary 删除一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Article not found"}"
// @Router /api/articles/{id} [DELETE]
// @Security JWT
func (*articlesApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Articles.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Param   articles      body model.Articles true  "文章"
// @Success 200 {object} model.Articles
// @Failure 404 {string} string "{"message": "Article not found"}"
// @Router /api/articles/{id} [PUT]
// @Security JWT
func (*articlesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var articles model.Articles

	var (
		apiReq *model.ArticleApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a articles")
	}
	if err := gconv.Struct(apiReq, &articles); err != nil {
		response.JsonOld(r, 400, "not a articles")
	}
	if _, err := dao.Articles.Data(articles).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	response.JsonOld(r, 200, articles)
}
