package repository

import (
	"gorm.io/gorm"

	"merchant/model"
)

type repo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repo{
		db,
	}
}

type Repository interface {
	ListMerchants() (model.Merchants, error)
	CreateMerchant(u *model.Merchant) error
	ReadMerchantById(id string) (*model.Merchant, error)
	ReadMerchantByEmail(email string) (*model.Merchant, error)
	UpdateMerchantDescriptionById(id string, d string) error
	DeleteMerchant(id string) error

	ListTeamMembersByMerchantId(merchantId string) (model.TeamMembers, error)
	CreateTeamMember(t *model.TeamMember) error
	ReadTeamMemberById(id string) (*model.TeamMember, error)
	ReadTeamMemberByEmail(email string) (*model.TeamMember, error)
	UpdateTeamMemberById(id string, t *model.TeamMember) error
	DeleteTeamMember(id string) error
}
