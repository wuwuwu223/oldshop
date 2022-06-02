package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oldshop/dao"
	"oldshop/model"
	"oldshop/utils"
)

func Login(c *gin.Context) {
	//用PostForm从post表单中获取用户名和密码
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.BindJSON(&data)
	username := data.Username
	password := data.Password
	password = utils.Md5(password)
	fmt.Println(username, password)

	//查询数据库
	var user model.User
	user.Username = username
	err := dao.FindUserByUsername(&user)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData("用户不存在", false, nil))
		return
	}
	//密码加密
	if user.Password != password {
		c.JSON(200, model.NewOkMsgWithData("密码错误", false, nil))
		return
	}
	//登录成功生成token
	token, err := utils.GenJwtToken(user.ID, user.Username)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData("生成token失败", false, nil))
		return
	}
	//返回token
	c.JSON(200, model.NewOkMsgWithData("登录成功", true, gin.H{"token": token}))
}

//注册
func Register(c *gin.Context) {
	var data struct {
		Avatar   string `json:"avatar"`
		Phone    string `json:"phone"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.BindJSON(&data)
	var user model.User
	user.Username = data.Username
	user.Password = data.Password
	user.Phone = data.Phone
	user.Avatar = data.Avatar
	//密码加密
	user.Password = utils.Md5(user.Password)
	err := dao.CreateUser(&user)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("注册成功", true, nil))
}

//获取用户信息
func GetUserInfo(c *gin.Context) {
	id, _ := c.Get("id")
	var user model.User
	user.ID = (id).(uint)
	err := dao.FindUserById(&user)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData("用户不存在", false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("获取用户信息成功", true, gin.H{
		"id":        user.ID,
		"name":      user.Username,
		"account":   user.Phone,
		"headerImg": user.Avatar,
	}))
}

//修改用户信息
func UpdateUserInfo(c *gin.Context) {
	id, _ := c.Get("id")
	var user model.User
	user.ID = (id).(uint)
	err := dao.FindUserById(&user)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData("用户不存在", false, nil))
		return
	}
	var data struct {
		Name string `json:"name"`
	}
	c.BindJSON(&data)
	user.Username = data.Name
	err = dao.UpdateUser(&user)
	if err != nil {
		c.JSON(200, model.NewOkMsgWithData(err.Error(), false, nil))
		return
	}
	c.JSON(200, model.NewOkMsgWithData("修改用户信息成功", true, nil))
}
