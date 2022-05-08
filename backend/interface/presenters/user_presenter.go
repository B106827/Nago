package presenters

import (
    "NagoBackend/domain/model"
)

type userPresenter struct {
}

func NewUserPresenter() UserPresenter {
    return &userPresenter{}
}

type UserPresenter interface {
    ResponseUsers(users []*model.User) []*model.User
}

func (userPresenter *userPresenter) ResponseUsers(users []*model.User) []*model.User {
    for _, u := range users {
        u.Name = u.Name + "hoge"
    }
    return users
}
