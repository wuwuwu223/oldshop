package model

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	UserId     uint       `json:"user_id"  gorm:"comment:'用户ID'"`
	Name       string     `json:"name" gorm:"type:varchar(100);comment:'商品名称'"`
	Status     uint       `json:"status" gorm:"default:0;comment:'状态;0.草稿;1.上架中;2.已下架;3.交易完成'"`
	Price      uint64     `json:"price" gorm:"comment:'价格'"`
	Freight    uint64     `json:"freight" gorm:"comment:'运费'"`
	Detail     string     `json:"detail" gorm:"comment:'商品详情'"`
	Order      Order      `json:"order"`
	Pictures   []Picture  `json:"pictures"`
	Categories []Category `json:"categories" gorm:"many2many:good_categories"`
	Tags       []Tag      `json:"tags" gorm:"many2many:good_tags"`
}

type Picture struct {
	gorm.Model
	Url    string `json:"url" gorm:"type:longtext;comment:'图片地址'"`
	GoodID uint   `json:"good_id " gorm:"comment:'商品ID'"`
}

type Tag struct {
	gorm.Model
	Key   string `json:"key" gorm:"type:varchar(20);comment:'关键字'"`
	Goods []Good `json:"goods" gorm:"many2many:good_tags"`
}

type Category struct {
	gorm.Model
	Index    int        `json:"index" gorm:"comment:'菜单排序'"`
	CIndex   uint8      `json:"c_index" gorm:"comment:'菜单级数'"`
	Name     string     `json:"name" gorm:"comment:'分类名称'"`
	ParentID *uint      `json:"parent_id" gorm:"comment:'上级菜单ID'"`
	Child    []Category `gorm:"foreignkey:ParentID"`
	Goods    []Good     `json:"goods" gorm:"many2many:good_categories"`
}
