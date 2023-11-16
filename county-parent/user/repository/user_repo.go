package repository

import (
	"context"
	"errors"
	"github.com/jenson/county/core/models"
	"github.com/jenson/county/core/util"
	"github.com/jenson/county/user/global"
	"log"
)

type UserRepo struct{}

func (ur UserRepo) Login(ctx context.Context, req *models.LoginReq) (*models.LoginResp, error) {
	db := global.DB
	var resp models.UserEntity
	err := db.Model(&models.UserEntity{}).Where("username = ?", req.Username).Find(&resp).Error
	if err != nil {
		log.Printf("execution database error %v", err)
		return nil, err
	}
	inputPwd := util.EncryptHex(req.Password)
	if inputPwd != resp.Password {
		log.Printf("password of user error")
		return nil, errors.New("登录密码错误")
	}
	//绑定为响应体
	result := models.LoginResp{ID: resp.ID, Username: resp.Username, Nickname: resp.Nickname}
	return &result, nil
}
