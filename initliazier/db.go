package initliazier

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"oldshop/global"
	"oldshop/model"
)

func InitDB() {
	//Db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true", global.Config.DB.User, global.Config.DB.Password, global.Config.DB.Host, global.Config.DB.Port, global.Config.DB.Name)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, //开启缓存
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	err = Db.AutoMigrate(&model.HomePic{}, &model.User{}, &model.RealProfile{}, &model.Address{}, &model.Good{}, &model.Category{}, &model.Tag{}, &model.Picture{}, &model.Order{}, &model.Logistics{}, &model.Conversation{}, &model.Message{})
	if err != nil {
		log.Fatal(err)
	}
	global.DB = Db
}
