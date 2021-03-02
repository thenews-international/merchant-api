package model

import "github.com/google/uuid"

type TeamMemberCreateForm struct {
	IsOwner    bool   `json:"isOwner" form:"required"`
	GivenName  string `json:"givenName" form:"required"`
	FamilyName string `json:"familyName" form:"required"`
	Email      string `json:"email" form:"required"`
}

type TeamMemberUpdateForm struct {
	IsOwner    bool   `json:"isOwner" form:"required"`
	GivenName  string `json:"givenName" form:"required"`
	FamilyName string `json:"familyName" form:"required"`
}

func (f *TeamMemberCreateForm) ToModel(merchantId string) *TeamMember {
	id := uuid.New().String()
	return &TeamMember{
		Model: Model{
			ID: id,
		},
		IsOwner:    f.IsOwner,
		GivenName:  f.GivenName,
		FamilyName: f.FamilyName,
		Email:      f.Email,
		Status:     "Active",
		MerchantID: merchantId,
	}
}

func (f *TeamMemberUpdateForm) ToModel(id string) *TeamMember {
	return &TeamMember{
		Model: Model{
			ID: id,
		},
		IsOwner:    f.IsOwner,
		GivenName:  f.GivenName,
		FamilyName: f.FamilyName,
	}
}
