package dao

import "oldshop/model"

//查询分类列表
func FindCategoryList(page, limit int) ([]*model.Category, error) {
	var categoryList []*model.Category
	err := db.Model(&model.Category{}).Offset((page - 1) * limit).Limit(limit).Find(&categoryList).Error
	return categoryList, err
}
