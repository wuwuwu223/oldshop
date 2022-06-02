package controller

import (
	"github.com/gin-gonic/gin"
	"oldshop/dao"
	"oldshop/model"
	"oldshop/utils"
	"strconv"
)

//创建商品api
func CreateGood(c *gin.Context) {
	id, _ := c.Get("id")
	var data struct {
		Name     string `json:"name"`
		Detail   string `json:"detail"`
		Price    string `json:"price"`
		Freight  string `json:"freight"`
		GoodType []uint `json:"good_type"`
		FileList []struct {
			Url     string `json:"url"`
			Content string `json:"content"`
		} `json:"file_list"`
	}
	var good model.Good
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	good.Name = data.Name
	good.Detail = data.Detail

	//Freight转float64
	ff, err := strconv.ParseFloat(data.Freight, 64)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	ff *= 100
	good.Freight = uint64(ff)
	//price转float64
	fprice, err := strconv.ParseFloat(data.Price, 64)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	fprice = fprice * 100

	//转换price
	good.Price = uint64(fprice)

	good.UserId = id.(uint)

	//把图片路径存入数据库
	for _, file := range data.FileList {
		var Pic model.Picture
		if file.Url != "" {
			Pic.Url = file.Url
		} else {
			Pic.Url = file.Content
		}
		good.Pictures = append(good.Pictures, Pic)
	}

	//把商品类型存入数据库
	for _, goodType := range data.GoodType {
		var gt model.Category
		gt.ID = goodType
		good.Categories = append(good.Categories, gt)
	}

	//分词name和detail
	NameSeg := utils.Segment(good.Name)
	DetailSeg := utils.Segment(good.Detail)

	for _, seg := range NameSeg {
		var tag model.Tag
		tag.Key = seg
		good.Tags = append(good.Tags, tag)
	}
	for _, seg := range DetailSeg {
		var tag model.Tag
		tag.Key = seg
		good.Tags = append(good.Tags, tag)
	}
	err = dao.CreateGood(&good)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("创建商品成功", true, nil))
}

//获取商品列表api
func GetGoodList(c *gin.Context) {
	goodList, err := dao.FindGoodList(1, 10)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("获取商品列表成功", true, goodList))
}

//获取商品详情api
func GetGoodDetail(c *gin.Context) {
	//从query获取商品id
	id, _ := c.GetQuery("id")
	//转换id
	goodId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	//获取商品详情
	good, err := dao.FindGoodDetail(uint(goodId))
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("获取商品详情成功", true, good))

}

//根据分类id获取商品列表api
func GetGoodListByCategoryId(c *gin.Context) {
	//从query获取分类id
	id, _ := c.GetQuery("id")
	//转换id
	categoryId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	//获取商品列表
	goodList, err := dao.FindGoodListByCategoryId(uint(categoryId), 1, 10)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("获取商品列表成功", true, goodList))
}

//根据关键字获取商品列表api
func GetGoodListByKey(c *gin.Context) {
	//从query获取关键字
	key, _ := c.GetQuery("key")
	//获取商品列表
	if key == "" || key == "undefined" {
		goodList, err := dao.FindGoodList(1, 10)
		if err != nil {
			c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
			return
		}
		c.JSON(200, model.NewOkMsgWithData("获取商品列表成功", true, goodList))
		return
	}
	goodList, err := dao.FindGoodListByKey(key, 1, 10)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("获取商品列表成功", true, goodList))
}
