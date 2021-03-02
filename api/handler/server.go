package handler

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"merchant/repository"
)

var (
	srvErrDataCreationFailure    = "data creation failure"
	srvErrDataAccessFailure      = "data access failure"
	srvErrDataDuplicateInsertion = "data duplicate insertion"
	srvErrDataUpdateFailure      = "data update failure"
	srvErrDataDeleteFailure      = "data delete failure"

	srvErrFormDecodingFailure   = "form decoding failure"
	srvErrHashGenerationFailure = "hash generation failure"
	srvErrJsonCreationFailure   = "json creation failure"

	srvErrAuthenticationFailure = "authentication failure"
)

type Server struct {
	DB        repository.Repository
	Validator *validator.Validate
	Logger    *zap.Logger
}

func New(
	db *gorm.DB,
	validator *validator.Validate,
	logger *zap.Logger,
) *Server {
	return &Server{
		DB:        repository.New(db),
		Validator: validator,
		Logger:    logger,
	}
}
