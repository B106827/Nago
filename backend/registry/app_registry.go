package registry

import "github.com/jinzhu/gorm"

type interactor struct {
    memberInteractor
}

type Interactor interface {
    MemberInteractor
}

func NewInteractor (conn *gorm.DB) Interactor {
    return &Interactor{
        memberInteractor{conn},
    }
}
