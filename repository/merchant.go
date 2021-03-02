package repository

import "merchant/model"

func (r *repo) ListMerchants() (model.Merchants, error) {
	ms := make([]*model.Merchant, 0)
	err := r.DB.Find(&ms).Error
	return ms, err
}

func (r *repo) CreateMerchant(u *model.Merchant) error {
	return r.DB.Create(&u).Error
}

func (r *repo) ReadMerchantById(id string) (*model.Merchant, error) {
	m := &model.Merchant{}
	if err := r.DB.Where(`id = ?`, id).First(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *repo) ReadMerchantByEmail(email string) (*model.Merchant, error) {
	m := &model.Merchant{}
	if err := r.DB.Where(`email = ?`, email).First(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *repo) UpdateMerchantDescriptionById(id, des string) error {
	return r.DB.Model(&model.Merchant{}).Where(`id = ?`, id).Update("description", des).Error
}

func (r *repo) DeleteMerchant(id string) error {
	return r.DB.Where(`id = ?`, id).Delete(&model.Merchant{}).Error
}
