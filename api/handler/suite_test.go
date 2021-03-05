package handler_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"merchant/api/handler"
	"merchant/mock/mock_repository"
)

type Suite struct {
	suite.Suite
	server *handler.Server
	db     *mock_repository.MockRepository
}

func (s *Suite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.db = mock_repository.NewMockRepository(ctrl)

	logger := zap.NewExample()
	_ = logger.Sync()

	s.server = handler.NewMockServer(s.db, validator.New(), logger)

	defer ctrl.Finish()
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
