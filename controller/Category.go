package controller

import (
	"github.com/gin-gonic/gin"
	"oldshop/dao"
	"oldshop/model"
)

//查询分类列表 api
func FindCategoryList(c *gin.Context) {
	categoryList, err := dao.FindCategoryList(1, 100)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("查询分类列表成功", true, categoryList))
}
