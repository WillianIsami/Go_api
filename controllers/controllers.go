package controllers

import (
	"github.com/WillianIsami/go_api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeControllers() {
	logger = config.GetLogger("controller")
	db = config.GetDB()
}
