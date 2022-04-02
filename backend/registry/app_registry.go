package registry

import "github.com/jinzhu/gorm"

type interactor struct {
    userInteractor
}

type Interactor interface {
    MemberInteractor
}

func NewInteractor (conn *gorm.DB) Interactor {
    return &Interactor{
        userInteractor{conn},
    }
}
