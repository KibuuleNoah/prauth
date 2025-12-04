package services

import (
	"context"

	"gorm.io/gorm"
)

type DataBaseService struct{
	DB *gorm.DB
	Ctx context.Context
}


