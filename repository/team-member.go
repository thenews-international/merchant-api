package repository

import "merchant/model"

func (r *repo) ListTeamMembersByMerchantId(merchantId string) (model.TeamMembers, error) {
	ts := make([]*model.TeamMember, 0)
	err := r.DB.Where("merchant_id", merchantId).Find(&ts).Error
	return ts, err
}

func (r *repo) CreateTeamMember(t *model.TeamMember) error {
	return r.DB.Create(&t).Error
}

func (r *repo) ReadTeamMemberById(id string) (*model.TeamMember, error) {
	t := &model.TeamMember{}
	if err := r.DB.Where(`id = ?`, id).First(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (r *repo) ReadTeamMemberByEmail(email string) (*model.TeamMember, error) {
	m := &model.TeamMember{}
	if err := r.DB.Where(`email = ?`, email).First(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *repo) UpdateTeamMemberById(id string, t *model.TeamMember) error {
	return r.DB.Model(&model.TeamMember{}).Where(`id = ?`, id).Updates(t).Error
}

func (r *repo) DeleteTeamMember(id string) error {
	return r.DB.Where(`id = ?`, id).Delete(&model.TeamMember{}).Error
}
