package pkg

import (
	"net"

	"gorm.io/gorm"
)

type APP struct {
	I1       int
	I2       float64
	DB       *gorm.DB
	Listener net.Listener
}

func GetAPP(i1 int, i2 float64, db *gorm.DB, listener net.Listener) *APP {
	return &APP{I1: i1, I2: i2, DB: db, Listener: listener}
}
