package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"

	"merchant/model"
)

var merchant = &model.Merchant{
	Model: model.Model{
		CreatedAt: &now,
		UpdatedAt: &now,
	},
}

var merchants = model.Merchants{
	merchant,
}

func (s *Suite) Test_repository_List_Merchant() {
	query := "SELECT * FROM `merchants`"
	rows := sqlmock.NewRows([]string{"id", "business_name", "status", "created_at", "updated_at"}).
		AddRow(merchant.ID, merchant.BusinessName, merchant.Status, now, now)

	s.mock.ExpectQuery(query).WillReturnRows(rows)

	res, err := s.repository.ListMerchants()

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(merchants, res))
}

func (s *Suite) Test_repository_Read_Merchant() {
	query := "SELECT * FROM `merchants` WHERE email = ? ORDER BY `merchants`.`id` LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "business_name", "status", "created_at", "updated_at"}).
		AddRow(merchant.ID, merchant.BusinessName, merchant.Status, now, now)

	s.mock.ExpectQuery(query).WithArgs(merchant.Email).WillReturnRows(rows)

	res, err := s.repository.ReadMerchantByEmail(merchant.Email)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(merchant, res))
}
