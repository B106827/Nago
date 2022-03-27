package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "NagoBackend/domain/model"
)

type memberRepository struct {
    db *gorm.DB
}

type MemberRepository interface {
    FindAll(members []*model.Member) ([]*model.Member, error)
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
    return &memberRepository{db}
}

func (memberRepository *memberRepository) FindAll(members []*model.Member) ([]*model.Member, error) {
    err := memberRepository.db.Find(&members).Error

    if err != nil {
        return nil, fmt.Errorf("sql error", err)
    }
    return members, nil
}
