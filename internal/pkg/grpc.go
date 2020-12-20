package pkg

import (
	app1_userDemo_v1 "Songzhibin/EngineeringDemo/api"
	"Songzhibin/EngineeringDemo/internal/biz"
	"Songzhibin/EngineeringDemo/internal/data"
	"Songzhibin/EngineeringDemo/internal/service"
	"net"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

var Service = grpc.NewServer()

func ReturnGrpcInit(db *gorm.DB) (net.Listener, func(), error) {
	server := service.NewService(biz.NewUserCase(data.NewLogicData(db)))
	app1_userDemo_v1.RegisterUserDemoServer(Service, server)
	lis, err := net.Listen("tcp", ":"+viper.GetString("APP.Port"))

	return lis, func() {
		lis.Close()
	}, err
}
