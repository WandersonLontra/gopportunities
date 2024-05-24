package handler

import (
	"github.com/WandersonLontra/gopportunities/configs"
	"gorm.io/gorm"
)

var (
	logger 	*configs.Logger
	db 		*gorm.DB
)

func InitializeHandler(){
	logger = configs.GetLogger("handler")
	db = configs.GetDatabase()
}