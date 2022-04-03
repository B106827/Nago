package registry

import "github.com/jinzhu/gorm"

type interactor struct {
    userInteractor
}

type Interactor interface {
    UserInteractor
}

func NewInteractor (conn *gorm.DB) Interactor {
    return &Interactor{
        userInteractor{conn},
    }
}
