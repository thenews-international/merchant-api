package model

type TeamMember struct {
	Model
	IsOwner    bool
	GivenName  string
	FamilyName string
	Email      string
	Status     string
	MerchantID string
}

type TeamMembers []*TeamMember

type TeamMemberDto struct {
	ID         string `json:"id"`
	IsOwner    bool   `json:"isOwner"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	MerchantID string `json:"merchantID"`
}

func (t TeamMember) ToDto() *TeamMemberDto {
	return &TeamMemberDto{
		ID:         t.ID,
		IsOwner:    t.IsOwner,
		GivenName:  t.GivenName,
		FamilyName: t.FamilyName,
		Email:      t.Email,
		Status:     t.Status,
		MerchantID: t.MerchantID,
	}
}

type TeamMemberDtos []*TeamMemberDto

func (ts TeamMembers) ToDto() TeamMemberDtos {
	result := make([]*TeamMemberDto, len(ts))
	for k, v := range ts {
		result[k] = v.ToDto()
	}

	return result
}
