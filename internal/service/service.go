/******
** @创建时间 : 2020/12/16 18:02
** @作者 : SongZhiBin
******/
package service

import (
	app1userDemov1 "Songzhibin/EngineeringDemo/api"
	"Songzhibin/EngineeringDemo/internal/biz"
	"context"
)

// Service: 实现api定义
type Service struct {
	userCase *biz.UserCase
}

func (s *Service) Login(ctx context.Context, request *app1userDemov1.LoginRequest) (*app1userDemov1.LoginResponse, error) {
	// 不需要依赖
	return &app1userDemov1.LoginResponse{
		IsLogin: s.userCase.QLogin(request.Username, request.Password),
	}, nil
}

func (s *Service) Registered(ctx context.Context, request *app1userDemov1.RegisteredRequest) (*app1userDemov1.RegisteredResponse, error) {
	// 转化为领域对象
	do := new(biz.User)
	do.Username = request.Username
	do.Password = request.Password
	do.Age = uint(request.Age)
	return &app1userDemov1.RegisteredResponse{
		IsRegistered: s.userCase.QReg(do),
	}, nil
}

func NewService(u *biz.UserCase) *Service {
	return &Service{userCase: u}
}
