package dao

import (
	"oldshop/model"
	"oldshop/utils"
	"sort"
)

//创建商品
func CreateGood(good *model.Good) error {
	return db.Create(good).Error
}

//获取商品列表并预加载图片
func FindGoodList(page int, pageSize int) ([]model.Good, error) {
	var goodList []model.Good
	err := db.Limit(pageSize).Offset((page - 1) * pageSize).Order("`id` desc").Preload("Pictures").Find(&goodList).Error
	//只返回第一张图片
	for i := 0; i < len(goodList); i++ {
		goodList[i].Pictures = goodList[i].Pictures[:1]
	}
	return goodList, err
}

//获取商品详情并预加载图片
func FindGoodDetail(id uint) (*model.Good, error) {
	var good model.Good
	err := db.Preload("Pictures").First(&good, id).Error
	return &good, err
}

//根据分类id获取商品列表
func FindGoodListByCategoryId(categoryId uint, page int, pageSize int) ([]model.Good, error) {
	var goodList []model.Good
	//查找所有匹配的关联记录
	var category model.Category
	err := db.Model(&model.Category{}).Where("id=?", categoryId).Preload("Goods").Find(&category).Error
	if err != nil {
		return nil, err
	}
	goodList = category.Goods
	//只返回第一张图片
	//for i := 0; i < len(goodList); i++ {
	//	goodList[i].Pictures = goodList[i].Pictures[:1]
	//}

	//获取good的图片

	for i := 0; i < len(goodList); i++ {
		goodList[i].Pictures = make([]model.Picture, 1)
		goodList[i].Pictures[0], err = FindPictureByGoodId(goodList[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return goodList, err
}

//根据商品id获取商品图片列表
func FindPictureByGoodId(goodId uint) (model.Picture, error) {
	var picture model.Picture
	err := db.Model(&model.Picture{}).Where("good_id=?", goodId).First(&picture).Error
	return picture, err
}

//关键字搜索商品
func FindGoodListByKey(key string, page int, pageSize int) ([]model.Good, error) {
	//关键字分词
	arr := utils.Segment(key)
	var tags []model.Tag
	//查找所有匹配的关联记录
	err := db.Model(&model.Tag{}).Where("`key` in (?)", arr).Preload("Goods").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	var goodList []model.Good
	//遍历所有tag去除重复的good 分析词频
	mapping := make(map[uint]int)
	map2 := make(map[uint]model.Good)
	for _, tag := range tags {
		for _, good := range tag.Goods {
			if _, ok := mapping[good.ID]; !ok {
				mapping[good.ID] = 1
				map2[good.ID] = good
			} else {
				mapping[good.ID]++
			}
		}
	}
	var sdatas sDatas
	for k, v := range mapping {
		var sdata sData
		sdata.Id = k
		sdata.Count = v
		sdatas = append(sdatas, sdata)
	}
	sort.Sort(&sdatas) //词频排序
	for i := range sdatas {
		goodList = append(goodList, map2[sdatas[i].Id])
	}
	//获取good的图片
	for i := 0; i < len(goodList); i++ {
		goodList[i].Pictures = make([]model.Picture, 1)
		goodList[i].Pictures[0], err = FindPictureByGoodId(goodList[i].ID)
		if err != nil {
			return nil, err
		}
	}
	return goodList, err
}

//Contains
func Contains(arr []model.Good, good model.Good) bool {
	for _, v := range arr {
		if v.ID == good.ID {
			return true
		}
	}
	return false
}

type sData struct {
	Id    uint
	Count int
}

type sDatas []sData

func (s sDatas) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func (s sDatas) Swap(i, j int) {
	//TODO implement me
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func (s sDatas) Len() int { return len(s) }
