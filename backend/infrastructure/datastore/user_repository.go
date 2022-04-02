package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "NagoBackend/domain/model"
)

type userRepository struct {
    db *gorm.DB
}

type MemberRepository interface {
    FindAll(users []*model.Member) ([]*model.Member, error)
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
    return &userRepository{db}
}

func (userRepository *userRepository) FindAll(users []*model.Member) ([]*model.Member, error) {
    err := userRepository.db.Find(&users).Error

    if err != nil {
        return nil, fmt.Errorf("sql error", err)
    }
    return users, nil
}
