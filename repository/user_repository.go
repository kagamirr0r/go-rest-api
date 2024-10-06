package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// UserRepositoryのインターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// UserRepositoryの構造体
type UserRepository struct {
	db *gorm.DB
}

// UserRepositoryのコンストラクタ(ファクトリ)関数
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
