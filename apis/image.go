package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"log"
	"strconv"
)

var Ak string
var Sk string
var Bucket string

// @Summary 获取一个上传文件的upToken
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   fileName      query string true  "要上传的文件名,如 abc.png"
// @Success 200 {string} string "upToken"
// @Router /image/UpToken [get]
// @Security ApiKeyAuth
func GetUpToken(c *gin.Context) {
	keyToOverwrite := c.Query("fileName")
	//log.Println(keyToOverwrite)
	//log.Println(Ak)
	//log.Println(Sk)
	//log.Println(Bucket)
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", Bucket, keyToOverwrite),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket  )","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(Ak, Sk)
	upToken := putPolicy.UploadToken(mac)
	c.JSON(200, upToken)
}

// @Summary 获取所有图片的url
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {string} string "["http://q52qkptnh.bkt.clouddn.com/1.png","http://q52qkptnh.bkt.clouddn.com/2.png","http://q52qkptnh.bkt.clouddn.com/3.png"]"
// @Router /image [get]
func GetBackGroundImageUrls(c *gin.Context) {
	var BackGroundImageUrls []models.Image
	Global.Database.Find(&BackGroundImageUrls)
	c.JSON(200, BackGroundImageUrls)
}

// @Summary 更改一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   id      path int true  "图片id" default(1)
// @Param   image      body httpModels.NoIdImage true  "图片"
// @Success 200 {object} models.Image
// @Failure 404 {string} string "{"message": "Image not found"}"
// @Router /image/{id} [PUT]
// @Security ApiKeyAuth
func UpdateImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your id is not a number"})
		return
	}
	var image models.Image
	Global.Database.First(&image, id)
	if image.ID == 0 {
		c.JSON(404, gin.H{"message": "image not found"})
		return
	}
	err = c.ShouldBindJSON(&image)
	if err != nil {
		log.Println(err)
	}
	if image.ID != id {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	if image.ImageType != 0 && image.ImageType != 1 && image.ImageType != 2 {
		c.JSON(400, gin.H{"message": "imagetype is not correct"})
		return
	}
	Global.Database.Save(&image)
	c.JSON(200, image)

}

// @Summary 获取一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   id      path int true  "图片id" default(1)
// @Success 200 {object} models.Image
// @Failure 404 {string} string "{"message":"Image not found"}"
// @Router /image/read/{id} [get]
func ReadImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var image models.Image
	Global.Database.First(&image, id)
	if image.ID == 0 {
		c.JSON(404, gin.H{"message": "Image not found"})
		return
	}
	c.JSON(200, image)
}

// @Summary 添加一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   image      body httpModels.NoIdImage true  "图片"
// @Success 200 {object} models.Image
// @Router /image [POST]
// @Security ApiKeyAuth
func CreateImage(c *gin.Context) {
	var image models.Image

	if err := c.BindJSON(&image); err != nil {
		log.Println(err)
		c.JSON(400, "Not a Image")
		return
	}

	if image.ID != 0 {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	if image.ImageType != 0 && image.ImageType != 1 && image.ImageType != 2 {
		c.JSON(400, gin.H{"message": "imagetype is not correct"})
		return
	}
	Global.Database.Create(&image)
	c.JSON(200, image)
}

var Place string
var Domain string

// @Summary 获取七牛云空间的地区和域名
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {string} string "{"domain":"q52qkptnh.bkt.clouddn.com","place": "华东}"
// @Router /image/PlaceAndDomain [get]
func GetPlaceAndDomain(c *gin.Context) {
	c.JSON(200, gin.H{"place": Place, "domain": Domain})
}
