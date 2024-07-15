package controllers

import (
	"go-blog/databases"

	"go.uber.org/zap"
)

type basicAttributes struct {
	db     databases.Database
	logger *zap.Logger
}
