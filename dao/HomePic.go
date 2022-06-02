package dao

import "oldshop/model"

//获取轮播图列表
func FindHomePicList(page int, pageSize int) ([]model.HomePic, error) {
	var homePicList []model.HomePic
	err := db.Limit(pageSize).Offset((page - 1) * pageSize).Order("`index` asc").Find(&homePicList).Error
	return homePicList, err
}
