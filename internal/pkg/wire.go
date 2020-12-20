// +build wireinject

package pkg

import (
	"github.com/google/wire"
)

func NewApp() (*APP, func(), error) {
	panic(wire.Build(ViperInit, LoggerInit, DataInit, ReturnGrpcInit, wire.Struct(new(APP), "*")))
}
