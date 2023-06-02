package repo

import (
	"context"
	"gorm-sqlmock-test/models"
)

func (repo *Implementation) SaveCompany(company models.Company, ctx context.Context) error {
	err := repo.db.WithContext(ctx).Create(&company).Error
	return err
}

func (repo *Implementation) FindCompanyById(id uint, ctx context.Context) (*models.Company, error) {
	var company *models.Company
	err := repo.db.WithContext(ctx).Joins("JOIN company_scales ON companies.company_scale_id = company_scales.id").Where("id = ?", id).First(&company).Error
	return company, err
}

func (repo *Implementation) DeleteCompanyById(id uint, ctx context.Context) error {
	var company models.Company
	err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&company).Error
	return err
}
