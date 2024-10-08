package service

import (
	"bluebook/biz/model/api"
	"bluebook/dal/db"
	utils "bluebook/pkg/util"
	"context"
	"fmt"
)

type UserService struct {
	ctx context.Context
}

// NewUserService new UserService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) Register(req *api.RegisterRequest) error {
	dao := db.NewDao(s.ctx)
	// 加密🔐密码
	passwd, err := utils.SetPassword(req.Password)
	if err != nil {
		return fmt.Errorf("service.Register failed, err: %v", err)
	}
	// 保证不出现空指针错误
	if req.Avator == nil {
		req.Avator = new(string)
		*req.Avator = ""
	}
	userModel := &db.User{
		Username: req.Username,
		Password: passwd,
		Avatar:   *req.Avator,
		Major:    req.Major,
		Email:    req.Email,
		Account:  req.Account,
		Role:     req.Role,
	}
	err = dao.CreateUser(userModel)
	if err != nil {
		return fmt.Errorf("service.Register failed, err: %v", err)
	}
	return nil
}

func (s *UserService) Login(req *api.LoginRequest) (*db.User, error) {
	dao := db.NewDao(s.ctx)
	user, err := dao.GetUserByAccount(req.Account)
	if err != nil {
		return nil, fmt.Errorf("service.Login failed, err: %v", err)
	}
	// 校验密码
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, fmt.Errorf("service.Login failed, err: %v", "invalid password")
	}
	return user, nil
}

func (s *UserService) GetUserInfo(username string) (*db.User, error) {
	dao := db.NewDao(s.ctx)
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("service.GetUserInfo failed, err: %v", err)
	}
	return user, nil
}
