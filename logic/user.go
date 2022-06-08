package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		//数据库查询出错
		return err
	}

	// 生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	err = mysql.InsterUser(user)
	//数据库插入出错
	if err != nil {

	}
	return err
}

func Login(p *models.ParamLogin)(token string,err error){
	user := &models.User{
		Username:p.Username,
		Password: p.Password,
	}
	//传递的user是指针，就能拿到user.UserID
	if err := mysql.Login(user);err != nil{
		return "",err
	}
	//生成 JWT 的Token
	return jwt.GenToken(user.UserID,user.Username)
}
