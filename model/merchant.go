package model

type Merchant struct {
	Model
	Email        string
	Password     string
	BusinessName string
	Description  string
	Status       string
}

type Merchants []*Merchant

type MerchantDto struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	BusinessName string `json:"businessName"`
	Description  string `json:"description"`
	Status       string `json:"status"`
}

func (m Merchant) ToDto() *MerchantDto {
	return &MerchantDto{
		ID:           m.ID,
		Email:        m.Email,
		BusinessName: m.BusinessName,
		Description:  m.Description,
		Status:       m.Status,
	}
}

type MerchantDtos []*MerchantDto

func (ms Merchants) ToDto() MerchantDtos {
	result := make([]*MerchantDto, len(ms))
	for k, v := range ms {
		result[k] = v.ToDto()
	}

	return result
}
