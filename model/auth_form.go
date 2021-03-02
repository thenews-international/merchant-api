package model

import "github.com/google/uuid"

type RegistrationForm struct {
	Email           string `json:"email" form:"required,email"`
	Password        string `json:"password" form:"required"`
	ConfirmPassword string `json:"confirmPassword" form:"required,eqfield=Password"`
	BusinessName    string `json:"businessName" form:"required"`
}

type LoginForm struct {
	Email    string `json:"email" form:"required,email"`
	Password string `json:"password" form:"required"`
}

type ChangePasswordForm struct {
	OldPassword     string `json:"oldPassword" form:"required,min=8"`
	Password        string `json:"password" form:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" form:"required,eqfield=Password"`
}

func (f *RegistrationForm) ToMerchantModel() *Merchant {
	id := uuid.New().String()
	return &Merchant{
		Model: Model{
			ID: id,
		},
		Email:        f.Email,
		Password:     f.Password,
		BusinessName: f.BusinessName,
		Status:       "Active",
	}
}
