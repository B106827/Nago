package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "NagoBackend/domain/model"
)

type userRepository struct {
    db *gorm.DB
}

type UserRepository interface {
    FindAll(users []*model.User) ([]*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (userRepository *userRepository) FindAll(users []*model.User) ([]*model.User, error) {
    err := userRepository.db.Find(&users).Error

    if err != nil {
        return nil, fmt.Errorf("sql error", err)
    }
    return users, nil
}
