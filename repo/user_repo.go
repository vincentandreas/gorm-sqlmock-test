package repo

import (
	"context"
	"gorm-sqlmock-test/models"
)

func (repo *Implementation) SaveUser(user models.User, ctx context.Context) error {
	err := repo.db.WithContext(ctx).Create(&user).Error
	return err
}

func (repo *Implementation) FindUserById(id uint, ctx context.Context) (*models.User, error) {
	var user *models.User
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

func (repo *Implementation) UpdateUserPasswordById(id uint, newPassword string, ctx context.Context) (*models.User, error) {
	var user *models.User
	err := repo.db.WithContext(ctx).Model(user).Where("id = ?", id).Update("password", newPassword).Error
	return user, err
}

func (repo *Implementation) DeleteUserById(id uint, ctx context.Context) error {
	var user models.User
	err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(user).Error
	return err
}
