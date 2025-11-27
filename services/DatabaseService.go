package services

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DataBaseService struct{
	Conn *pgx.Conn
	Ctx context.Context
}


