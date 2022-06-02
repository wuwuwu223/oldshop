package controller

import (
	"github.com/gin-gonic/gin"
	"oldshop/dao"
	"oldshop/model"
)

//轮播图列表api
func HomePicList(c *gin.Context) {
	//var data struct {
	//	Page  int `json:"page"`
	//	Limit int `json:"limit"`
	//}
	//c.BindJSON(&data)
	homePicList, err := dao.FindHomePicList(1, 4)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("获取轮播图列表成功", true, homePicList))
}
