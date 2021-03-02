package repository

//
//import (
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/go-test/deep"
//	"github.com/stretchr/testify/require"
//
//	"merchant/model"
//)
//
//var merchant = &model.Merchant{
//	Model: model.Model{
//		ID:        "adfasdf",
//		CreatedAt: &now,
//		UpdatedAt: &now,
//	},
//	BusinessName: "PaceNow",
//	Status:       "Active",
//}
//
//var merchants = []*model.Merchant{
//	merchant,
//}
//
//func (s *Suite) Test_repository_List_Merchant() {
//	s.T().Skip("TODO")
//	query := "SELECT (.*) FROM merchants"
//	rows := sqlmock.NewRows([]string{"id", "business_name", "status", "created_at", "updated_at"}).
//		AddRow(merchant.ID, merchant.BusinessName, merchant.Status, now, now)
//
//	s.mock.MatchExpectationsInOrder(false)
//	s.mock.ExpectQuery(query).WillReturnRows(rows)
//
//	res, err := s.repository.ListMerchants()
//
//	require.NoError(s.T(), err)
//	require.Nil(s.T(), deep.Equal(&merchants, res))
//}
//
//func (s *Suite) Test_repository_Read_Merchant() {
//	s.T().Skip("TODO")
//	var (
//		id   uint = 1
//		name      = "test-name"
//	)
//
//	s.mock.ExpectQuery(
//		`SELECT * FROM merchants WHERE (id = :id)`).
//		WithArgs(id).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//			AddRow(id, name))
//
//	res, err := s.repository.ReadMerchantByEmail(merchant.Email)
//
//	require.NoError(s.T(), err)
//	require.Nil(s.T(), deep.Equal(&merchant, res))
//}
